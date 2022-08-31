package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"golang.org/x/net/context"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
        addr = "0.0.0.0"
        port = 9094
)

func main() {
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

	c := exec.Command("bash", "-c", "ps -e | grep seid | grep -v grep | wc -l")
	var out bytes.Buffer
	c.Stdout = &out
	c.Run()

	result, err := strconv.Atoi(strings.ReplaceAll(out.String(), "\n", ""))
	if err != nil {
		fmt.Printf("%v \n", err)
	}

	contentMSG := ""
	if result > 0 {
		log.Info().Str("process", "up").Msg(fmt.Sprintf("sei Process alive"))
		contentMSG = "sei Process is UP"
		state = pluginpb.STATE_SUCCESS
	} else {
		log.Info().Str("process", "up").Msg(fmt.Sprintf("sei Process died"))
		contentMSG = "sei Process is DOWN"
		severity = pluginpb.SEVERITY_CRITICAL
	}

	ret := sdk.CallResponse{
		FuncName:   "seiUP",
		Message:    contentMSG,
		Severity:   severity,
		State:      state,
		AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
	}
	return ret, nil
}
