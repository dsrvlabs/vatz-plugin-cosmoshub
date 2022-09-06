package main

import (
	"flag"
	"fmt"
	"time"

	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/cosmos-sdk-blocksync/policy"
	"github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/cosmos-sdk-blocksync/status"
)

const (
	// Default values.
	defaultAddr = "127.0.0.1"
	defaultPort = 9091

	pluginName = "cosmos-sdk-block-height"

	// Parameters for alert
	maxHistoryCount = 100
	fetchInterval   = time.Second * 2

	estimateTimeWindow     = time.Second * 30
	blockIntervalThreshold = time.Millisecond * 7500
)

var (
	addr string
	port int

	collector status.Collector
	estimator policy.Estimator
)

func init() {
	collector = status.NewStatusCollector()
	estimator = policy.NewEstimator()
}

func main() {
	flag.StringVar(&addr, "addr", defaultAddr, "IP Address(e.g. 0.0.0.0, 127.0.0.1)")
	flag.IntVar(&port, "port", defaultPort, "Port number, default 9091")

	flag.Parse()

	collector.Start(maxHistoryCount, fetchInterval)

	p := sdk.NewPlugin(pluginName)
	p.Register(pluginFeature)

	ctx := context.Background()
	if err := p.Start(ctx, addr, port); err != nil {
		// TODO: Stop collector
		fmt.Println("exit")
	}
}

func pluginFeature(info, option map[string]*structpb.Value) (sdk.CallResponse, error) {
	log.Info().Str("main", "main").Msg(fmt.Sprintf("pluginFeature: %s", info["execute_method"]))

	// TODO: filter out by execute_method.
	histories := collector.GetHistories()
	status, err := estimator.Estimate(histories, estimateTimeWindow, blockIntervalThreshold)
	if err != nil {
		log.Error().Str("main", "main").Msg(err.Error())
		return sdk.CallResponse{}, err
	}

	ret := sdk.CallResponse{
		FuncName:   info["execute_method"].GetStringValue(),
		Message:    "OK",
		Severity:   pluginpb.SEVERITY_UNKNOWN,
		State:      pluginpb.STATE_NONE,
		AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
	}

	if status == policy.AlertStatusAlert {
		ret.Severity = pluginpb.SEVERITY_CRITICAL
		ret.State = pluginpb.STATE_FAILURE
		ret.Message = "Block Height stucks"
	}

	return ret, nil
}
