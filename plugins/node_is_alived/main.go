package main

import (
	"flag"
	"fmt"

	health "github.com/dsrvlabs/vatz-plugin-cosmoshub/rpc/cosmos"
	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	defaultAddr    = "127.0.0.1"
	defaultPort    = 10002
	defaultRPCAddr = "http://localhost:26657"
)

var (
	addr    string
	port    int
	rpcAddr string
)

// Health is response entity from REST.
type Health struct {
	ID      int          `json:"id"`
	Jsonrpc string       `json:"jsonrpc"`
	Result  HealthResult `json:"result"`
}

// HealthResult is Health.Result sturct
type HealthResult struct {
}

func main() {
	flag.StringVar(&addr, "addr", defaultAddr, "IP Address(e.g. 0.0.0.0, 127.0.0.1)")
	flag.IntVar(&port, "port", defaultPort, "Port number")
	flag.StringVar(&rpcAddr, "rpcAddr", defaultRPCAddr, "RPC addrest:port (e.g. http://127.0.0.1:26667)")

	flag.Parse()

	p := sdk.NewPlugin("is_alived")
	p.Register(pluginFeature)

	ctx := context.Background()
	if err := p.Start(ctx, addr, port); err != nil {
		fmt.Println("exit")
	}
}

func pluginFeature(info, option map[string]*structpb.Value) (sdk.CallResponse, error) {
	state := pluginpb.STATE_NONE
	severity := pluginpb.SEVERITY_INFO
	healthStatus, err := health.GetHealth(rpcAddr)
	if err != nil {
		log.Error().Str("GetHealth", "Error").Msg(fmt.Sprintf("%v", err))
		contentMSG := "UNHEALTHY"
		return sdk.CallResponse{
			FuncName:   "GetHealth",
			Message:    contentMSG,
			Severity:   pluginpb.SEVERITY_CRITICAL,
			State:      pluginpb.STATE_FAILURE,
			AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
		}, nil
	}

	contentMSG := ""

	if healthStatus == 200 {
		log.Info().Str("process", "up").Msg(fmt.Sprintf("HEALTHY"))
		contentMSG = "HEALTHY"
		state = pluginpb.STATE_SUCCESS
	} else {
		log.Info().Str("process", "up").Msg(fmt.Sprintf("UNHEALTHY"))
		contentMSG = "UNHEALTHY"
		state = pluginpb.STATE_SUCCESS
		severity = pluginpb.SEVERITY_CRITICAL
	}

	ret := sdk.CallResponse{
		FuncName:   "GetHealth",
		Message:    contentMSG,
		Severity:   severity,
		State:      state,
		AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
	}
	return ret, nil
}
