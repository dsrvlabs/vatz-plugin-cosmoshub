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
	defaultProposalId = 0
	iterCnt = 100

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
	flag.UintVar(&proposalId, "proposalId", defaultProposalId, "Need to know last proposal id")

	flag.Parse()

	findLastestProposalId()
}

func findLastestProposalId() {
	proposalId = 0
	cntNotFoundProp := 0
	fmt.Println("zz", proposalId)

	var msg string

	for i := 1; ; i++ {
		prop, err := rpcGovernance.GetProposal(apiPort, uint(i))
		if err == nil {
			num_prop, _ := strconv.Atoi(prop)
			msg += fmt.Sprintf("New proposal : %d\n", num_prop)
			proposalId = uint(i)
			cntNotFoundProp = 0
		} else {
			cntNotFoundProp++
			if cntNotFoundProp == iterCnt {
				msg += fmt.Sprintf("Lastest proposal is #%d", proposalId)
				break
			}
		}
	}
	log.Info().Str("module", "plugin").Msg(msg)
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
	tmp := proposalId
	fmt.Println("Debug: last proposal = ", proposalId)

	for i := 1; i <= iterCnt; i++ {
		prop, err := rpcGovernance.GetProposal(apiPort, tmp + uint(i))
		if err == nil {
			num_prop, _ := strconv.Atoi(prop)
			msg += fmt.Sprintf("New proposal : %d\n", num_prop)
			proposalId = tmp + uint(i)
		} else {
			if i == iterCnt {
				msg += fmt.Sprintf("Lastest proposal is #%d", proposalId)
			}
		}
	}
	if tmp == proposalId {
		fmt.Println("DEBUG : tmp == proposalId")
		severity = pluginpb.SEVERITY_INFO
		state = pluginpb.STATE_SUCCESS
	} else {
		fmt.Println("DEBUG : tmp != proposalId")
		severity = pluginpb.SEVERITY_CRITICAL
		state = pluginpb.STATE_SUCCESS
	}
	log.Info().Str("module", "plugin").Msg(msg)
	ret := sdk.CallResponse{
		FuncName:   info["execute_method"].GetStringValue(),
		Message:    msg,
		Severity:   severity,
		State:      state,
		AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
	}
	return ret, nil
}
