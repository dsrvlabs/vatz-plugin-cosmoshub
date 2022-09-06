package main

import (
        "flag"
        "fmt"

        pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
        "github.com/dsrvlabs/vatz/sdk"
        "golang.org/x/net/context"
        "google.golang.org/protobuf/types/known/structpb"
)

const (
        // Default values.
        defaultAddr = "127.0.0.1"
        defaultPort = 9091

        pluginName = "cosmoshub-peer-count"
)

var (
        addr string
        port int
)

func init() {
        flag.StringVar(&addr, "addr", defaultAddr, "IP Address(e.g. 0.0.0.0, 127.0.0.1)")
        flag.IntVar(&port, "port", defaultPort, "Port number, default 9091")

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
        // TODO: Fill here.
        ret := sdk.CallResponse{
                FuncName:   info["execute_method"].GetStringValue(),
                Message:    "Peer Count is too low",
                Severity:   pluginpb.SEVERITY_UNKNOWN,
                State:      pluginpb.STATE_NONE,
                AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
        }

        return ret, nil
}
