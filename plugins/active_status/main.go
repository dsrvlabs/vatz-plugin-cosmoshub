package main

import (
	"flag"
	"fmt"

	rpcCosmos "github.com/dsrvlabs/vatz-plugin-cosmoshub/rpc/cosmos"
	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	// Default values.
	defaultRPCAddr = "http://localhost:1317"
	defaultAddr    = "127.0.0.1"
	defaultPort    = 9100
	pluginName     = "cosmoshub-active-status"
)

var (
	rpcAddr     string
	addr        string
	port        int
	valoperAddr string
)

func init() {
	flag.StringVar(&rpcAddr, "rpcURI", defaultRPCAddr, "CosmosHub RPC URI Address")
	flag.StringVar(&addr, "addr", defaultAddr, "Listening address")
	flag.IntVar(&port, "port", defaultPort, "Listening port")
	flag.StringVar(&valoperAddr, "valoperAddr", "", "CosmosHub validator operator address")

	flag.Parse()
}

func main() {
	if valoperAddr == "" {
		log.Fatal().Str("module", "plugin").Msg("Please specify -valoperAddr")
	}

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

	status, err := rpcCosmos.GetBondStatus(rpcAddr, valoperAddr)

	if err == nil {
		state = pluginpb.STATE_SUCCESS
		if status == true {
			severity = pluginpb.SEVERITY_INFO
			msg = fmt.Sprintf("Validator bonded. included active set")
		} else {
			severity = pluginpb.SEVERITY_CRITICAL
			msg = fmt.Sprintf("Validator unbonded. kick out from active set")
		}
		log.Debug().Str("module", "plugin").Msg(msg)
	} else {
		// Maybe node wil be killed. So other alert comes to you.
		severity = pluginpb.SEVERITY_CRITICAL
		state = pluginpb.STATE_FAILURE
		msg = "Failed to get validator status"
		log.Info().Str("moudle", "plugin").Msg(msg)
	}

	ret := sdk.CallResponse{
		FuncName:   info["execute_method"].GetStringValue(),
		Message:    msg,
		Severity:   severity,
		State:      state,
		AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
	}

	return ret, nil
}
