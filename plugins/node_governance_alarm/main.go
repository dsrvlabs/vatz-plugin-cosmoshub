package main

import (
	"flag"
	"fmt"
	"strconv"

	rpcGovernance "github.com/dsrvlabs/vatz-plugin-cosmoshub/rpc/cosmos"
	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	// Default values.
	defaultAddr = "127.0.0.1"
	defaultPort = 9091
	defaultPeer = 5
	defaultApiPort = 1317
	defaultProposalId = 100

	pluginName = "node-governance-alarm"
)

var (
	addr		string
	port		int
	apiPort		uint
	proposalId	uint
)

func init() {
	flag.StringVar(&addr, "addr", defaultAddr, "IP Address(e.g. 0.0.0.0, 127.0.0.1)")
	flag.IntVar(&port, "port", defaultPort, "Port number, default 9091")
	flag.UintVar(&apiPort, "apiPort", defaultApiPort, "Need to know proposal id")
	flag.UintVar(&proposalId, "last proposal id", defaultProposalId, "Need to know last proposal id")

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

	prop, err := rpcGovernance.GetProposal(apiPort, proposalId)
	if err == nil {
		num_prop, _ := strconv.Atoi(prop)
		severity = pluginpb.SEVERITY_INFO
		state = pluginpb.STATE_SUCCESS
		msg = fmt.Sprintf("New proposal : %d", num_prop)
		log.Info().Str("module", "plugin").Msg(msg)
		/*
		if npeer < minPeer {
			severity = pluginpb.SEVERITY_CRITICAL
			state = pluginpb.STATE_SUCCESS
			msg = fmt.Sprintf("Bad: peer_count is %d", npeer)
			log.Info().Str("moudle", "plugin").Msg(msg)
		} else {
			severity = pluginpb.SEVERITY_INFO
			state = pluginpb.STATE_SUCCESS
			msg = fmt.Sprintf("Good: peer_count is %d", npeer)
			log.Info().Str("moudle", "plugin").Msg(msg)
		}
		*/
	} else {
		// Maybe node wil be killed. So other alert comes to you.
		severity = pluginpb.SEVERITY_CRITICAL
		state = pluginpb.STATE_SUCCESS
		msg = "Error to get proposal id"
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