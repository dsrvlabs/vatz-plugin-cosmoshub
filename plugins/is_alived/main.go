package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	defaultAddr  = "127.0.0.1"
	defaultPort  = 9098
	localRPCAddr = "http://localhost:26657"
)

var (
	addr string
	port int
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
	flag.IntVar(&port, "port", defaultPort, "Port number, default 9098")

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

	req, err := http.NewRequest("GET", localRPCAddr+"/health", nil)
	if err != nil {
		contentMSG := "request error"
		severity := pluginpb.SEVERITY_ERROR
		return sdk.CallResponse{
			FuncName:   "gaiadUP",
			Message:    contentMSG,
			Severity:   severity,
			State:      state,
			AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
		}, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		contentMSG := "request progress error"
		severity := pluginpb.SEVERITY_ERROR
		return sdk.CallResponse{
			FuncName:   "gaiadUP",
			Message:    contentMSG,
			Severity:   severity,
			State:      state,
			AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
		}, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		contentMSG := "request result error"
		severity := pluginpb.SEVERITY_ERROR
		return sdk.CallResponse{
			FuncName:   "gaiadUP",
			Message:    contentMSG,
			Severity:   severity,
			State:      state,
			AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
		}, err
	}
	var health Health
	json.Unmarshal([]byte(string(bytes)), &health)

	contentMSG := ""
	if health.Result == (HealthResult{}) {
		log.Info().Str("process", "up").Msg(fmt.Sprintf("gaiad Process alive"))
		contentMSG = "gaiad Process is UP"
		state = pluginpb.STATE_SUCCESS
	} else {
		log.Info().Str("process", "up").Msg(fmt.Sprintf("gaiad Process died"))
		contentMSG = "gaiad Process is DOWN"
		severity = pluginpb.SEVERITY_CRITICAL
	}

	ret := sdk.CallResponse{
		FuncName:   "gaiadUP",
		Message:    contentMSG,
		Severity:   severity,
		State:      state,
		AlertTypes: []pluginpb.ALERT_TYPE{pluginpb.ALERT_TYPE_DISCORD},
	}
	return ret, nil
}
