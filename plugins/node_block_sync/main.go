package main

import (
	"flag"
	"fmt"
	"strconv"

	rpcCosmos "github.com/dsrvlabs/vatz-plugin-cosmoshub/rpc/cosmos"
	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	// Default values.
	defaultRPCAddr = "http://localhost:26657"
	defaultAddr    = "127.0.0.1"
	defaultPort    = 9091
	pluginName     = "cosmoshub-block-sync"
	defaultCriticalCount = 3
)

var (
	rpcAddr		string
	addr		string
	port		int
	prevHeight	int
	latestHeight	int
	warningCount	int
	criticalCount   int
)

func init() {
	flag.StringVar(&rpcAddr, "rpcURI", defaultRPCAddr, "Tendermint RPC URI Address")
	flag.StringVar(&addr, "addr", defaultAddr, "Listening address")
	flag.IntVar(&port, "port", defaultPort, "Listening port")
	flag.IntVar(&criticalCount, "critical", defaultCriticalCount, "block height stucked count to raise critical level of alert")

	flag.Parse()
}

func main() {
	p := sdk.NewPlugin(pluginName)
	p.Register(pluginFeature)

	ctx := context.Background()
	if err := p.Start(ctx, addr, port); err != nil {
		fmt.Println("exit")
	}
}

func pluginFeature(info, option map[string]*structpb.Value) (sdk.CallResponse, error) {
	severity := pluginpb.SEVERITY_INFO
	state := pluginpb.STATE_NONE

	var msg string

	status, err := rpcCosmos.GetStatus(rpcAddr)

	if err == nil {
		latestHeight, _ = strconv.Atoi(status.Result.SyncInfo.LatestBlockHeight)
		log.Info().Str("module", "plugin").Msgf("previous block height: %d, latest block height: %d", prevHeight, latestHeight)
		state = pluginpb.STATE_SUCCESS

		if latestHeight > prevHeight {
			severity = pluginpb.SEVERITY_INFO
			msg = fmt.Sprintf("block height increasing")
			warningCount = 0
		} else {
			severity = pluginpb.SEVERITY_WARNING
			warningCount++
			msg = fmt.Sprintf("block height stucked %d times", warningCount)
		}

		if warningCount > criticalCount {
			severity = pluginpb.SEVERITY_CRITICAL
			msg = fmt.Sprintf("block height stucked more than %d times", warningCount)
		}
		log.Debug().Str("module", "plugin").Msg(msg)
	} else {
		// Maybe node will be killed. So other alert comes to you.
		severity = pluginpb.SEVERITY_CRITICAL
		state = pluginpb.STATE_FAILURE
		msg = "Failed to get node status"
		log.Info().Str("moudle", "plugin").Msg(msg)
	}

	ret := sdk.CallResponse{
		FuncName:   info["execute_method"].GetStringValue(),
		Message:    msg,
		Severity:   severity,
		State:      state,
		AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
	}

	prevHeight = latestHeight
	return ret, nil
}
