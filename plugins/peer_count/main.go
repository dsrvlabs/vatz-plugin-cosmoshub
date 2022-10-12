package main

import (
	"flag"
	"fmt"
	"strconv"
	"os"

	rpcNetInfo "github.com/dsrvlabs/vatz-plugin-cosmoshub/rpc/cosmos"
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

	pluginName = "cosmoshub-peer-count"
)

var (
	addr string
	port int
	minPeer int
)

func init() {
	flag.StringVar(&addr, "addr", defaultAddr, "IP Address(e.g. 0.0.0.0, 127.0.0.1)")
	flag.IntVar(&port, "port", defaultPort, "Port number, default 9091")
	flag.IntVar(&minPeer, "minPeer", defaultPeer, "minimum peer count, default 5")

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
	var hostname string

	peers, err := rpcNetInfo.GetNpeers()
	if err == nil {
		npeer, _ := strconv.Atoi(peers)
		if npeer < minPeer {
			severity = pluginpb.SEVERITY_CRITICAL
			hostname, _ = os.Hostname()
			msg = fmt.Sprintf("[%s]Bad: peer_count is %d", hostname, npeer)
			log.Info().Str("moudle", "plugin").Msg(msg)
		} else {
			state = pluginpb.STATE_SUCCESS
			hostname, _ = os.Hostname()
			msg = fmt.Sprintf("[%s]Good: peer_count is %d", hostname, npeer)
			log.Info().Str("moudle", "plugin").Msg(msg)
		}
	} else {
		severity = pluginpb.SEVERITY_ERROR
		hostname, _ = os.Hostname()
		msg = fmt.Sprintf("[%s]Error to get #N peers", hostname)
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
