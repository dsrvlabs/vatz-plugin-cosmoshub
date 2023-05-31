package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	rpcGovernance "github.com/dsrvlabs/vatz-plugin-cosmoshub/rpc/cosmos"
	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	// Default values.
	defaultAddr       = "127.0.0.1"
	defaultPort       = 10005
	defaultPeer       = 5
	defaultApiPort    = 1317
	defaultProposalId = 0
	defaultVoterVote  = "address"
	iterCnt           = 100

	pluginName = "node-governance-alarm"
)

var (
	addr       string
	port       int
	apiPort    uint
	proposalId uint
	voterAddr  string
	firstRun   int
)

func init() {
	flag.StringVar(&addr, "addr", defaultAddr, "IP Address(e.g. 0.0.0.0, 127.0.0.1)")
	flag.IntVar(&port, "port", defaultPort, "Port number, default 9091")
	flag.UintVar(&apiPort, "apiPort", defaultApiPort, "Need to know proposal id")
	flag.UintVar(&proposalId, "proposalId", defaultProposalId, "Need to know last proposal id")
	flag.StringVar(&voterAddr, "voterAddr", defaultVoterVote, "Need to voter address")

	flag.Parse()

	findLastestProposalId()
	firstRun = 0
}

func CheckVoterVote() string {
	tmpCnt := iterCnt
	var needToVote string
	for j := proposalId; ; j-- {
		_, t, _ := rpcGovernance.GetProposal(apiPort, uint(j))
		if j == 0 || tmpCnt == 0 || t.Before(time.Now()) {
			break
		}
		tmpCnt--
		vote, _ := rpcGovernance.GetVoterVote(apiPort, uint(j), voterAddr)
		var myVote string
		if vote == "VOTE_OPTION_YES" {
			myVote = "Yes"
		} else if vote == "VOTE_OPTION_NO" {
			myVote = "No"
		} else if vote == "VOTE_OPTION_ABSTAIN" {
			myVote = "Abstain"
		} else if vote == "VOTE_OPTION_NO_WITH_VETO" {
			myVote = "NoWithVeto"
		} else {
			myVote = "Dit not vote"
		}
		voteMsg := fmt.Sprintf("Proposal %d: %s", j, myVote)
		needToVote += voteMsg + "\n"
		log.Info().Str("module", "plugin").Msg(voteMsg)
	}
	return needToVote
}

func findLastestProposalId() {
	proposalId = 0
	cntNotFoundProp := 0

	var msg string

	for i := 1; ; i++ {
		prop, t, err := rpcGovernance.GetProposal(apiPort, uint(i))
		if err == nil && t.Year() > time.Now().Year()/2 {
			num_prop, _ := strconv.Atoi(prop)
			msg += fmt.Sprintf("New proposal : %d, %d-%02d-%02d\n", num_prop, t.Year(), t.Month(), t.Day())
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
	CheckVoterVote()
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

	for i := 1; i <= iterCnt; i++ {
		prop, t, err := rpcGovernance.GetProposal(apiPort, tmp+uint(i))
		if err == nil && t.After(time.Now()) {
			num_prop, _ := strconv.Atoi(prop)
			msg += fmt.Sprintf("New proposal : %d, %d-%02d-%02d\n", num_prop, t.Year(), t.Month(), t.Day())
			proposalId = tmp + uint(i)
		} else {
			if i == iterCnt {
				msg += fmt.Sprintf("Lastest proposal is #%d\n", proposalId)
			}
		}
	}
	if tmp == proposalId {
		log.Debug().Str("module", "plugin").Msg("DEBUG : tmp == proposalId")
		severity = pluginpb.SEVERITY_INFO
		state = pluginpb.STATE_SUCCESS
		list := CheckVoterVote()
		if len(list) != 0 && firstRun == 0 {
			severity = pluginpb.SEVERITY_CRITICAL
			state = pluginpb.STATE_SUCCESS
			msg += list
			firstRun++
		}
	} else {
		log.Debug().Str("module", "plugin").Msg("DEBUG : tmp != proposalId")
		severity = pluginpb.SEVERITY_CRITICAL
		state = pluginpb.STATE_SUCCESS
		list := CheckVoterVote()
		if len(list) != 0 {
			msg += list
		}
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
