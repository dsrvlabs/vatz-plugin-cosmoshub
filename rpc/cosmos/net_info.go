package rpc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type netInfo struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Listening bool     `json:"listening"`
		Listeners []string `json:"listeners"`
		NPeers    string   `json:"n_peers"`
		Peers     []struct {
			NodeInfo struct {
				ProtocolVersion struct {
					P2P   string `json:"p2p"`
					Block string `json:"block"`
					App   string `json:"app"`
				} `json:"protocol_version"`
				ID         string `json:"id"`
				ListenAddr string `json:"listen_addr"`
				Network    string `json:"network"`
				Version    string `json:"version"`
				Channels   string `json:"channels"`
				Moniker    string `json:"moniker"`
				Other      struct {
					TxIndex    string `json:"tx_index"`
					RPCAddress string `json:"rpc_address"`
				} `json:"other"`
			} `json:"node_info"`
			IsOutbound       bool `json:"is_outbound"`
			ConnectionStatus struct {
				Duration    string `json:"Duration"`
				SendMonitor struct {
					Start    time.Time `json:"Start"`
					Bytes    string    `json:"Bytes"`
					Samples  string    `json:"Samples"`
					InstRate string    `json:"InstRate"`
					CurRate  string    `json:"CurRate"`
					AvgRate  string    `json:"AvgRate"`
					PeakRate string    `json:"PeakRate"`
					BytesRem string    `json:"BytesRem"`
					Duration string    `json:"Duration"`
					Idle     string    `json:"Idle"`
					TimeRem  string    `json:"TimeRem"`
					Progress int       `json:"Progress"`
					Active   bool      `json:"Active"`
				} `json:"SendMonitor"`
				RecvMonitor struct {
					Start    time.Time `json:"Start"`
					Bytes    string    `json:"Bytes"`
					Samples  string    `json:"Samples"`
					InstRate string    `json:"InstRate"`
					CurRate  string    `json:"CurRate"`
					AvgRate  string    `json:"AvgRate"`
					PeakRate string    `json:"PeakRate"`
					BytesRem string    `json:"BytesRem"`
					Duration string    `json:"Duration"`
					Idle     string    `json:"Idle"`
					TimeRem  string    `json:"TimeRem"`
					Progress int       `json:"Progress"`
					Active   bool      `json:"Active"`
				} `json:"RecvMonitor"`
				Channels []struct {
					ID                int    `json:"ID"`
					SendQueueCapacity string `json:"SendQueueCapacity"`
					SendQueueSize     string `json:"SendQueueSize"`
					Priority          string `json:"Priority"`
					RecentlySent      string `json:"RecentlySent"`
				} `json:"Channels"`
			} `json:"connection_status"`
			RemoteIP string `json:"remote_ip"`
		} `json:"peers"`
	} `json:"result"`
}

func GetNpeers() (string, error) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:26657/net_info", nil)
	if err != nil {
		return err.Error(), err
	}

	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err.Error(), err
	}

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error(), err
	}

	if resp.StatusCode != http.StatusOK {
		return strconv.Itoa(resp.StatusCode), fmt.Errorf("request failed %s", string(rawBody))
	}

	info := netInfo{}
	err = json.Unmarshal(rawBody, &info)

	if err != nil {
		return err.Error(), err
	}

	return info.Result.NPeers, nil
}
