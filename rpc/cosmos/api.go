package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	localRPCAddr = "http://localhost:26657"
)

// Status is response entity frmo REST.
type Status struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int `json:"id"`
	Result  struct {
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
		SyncInfo struct {
			LatestBlockHash   string `json:"latest_block_hash"`
			LatestAppHash     string `json:"latest_app_hash"`
			LatestBlockHeight string `json:"latest_block_height"`
			LatestBlockTime   string `json:"latest_block_time"`
			EarliestBlockHash	string	`json:"earliest_block_hash"`
			EarliestAppHash		string	`json:"earliest_app_hash"`
			EarliestBlockHeight	string	`json:"earliest_block_height"`
			EarliestBlockTime	string	`json:"earliest_block_time"`
			CatchingUp         bool   `json:"catching_up"`
		} `json:"sync_info"`
		ValidatorInfo struct {
			Address string `json:"address"`
			PubKey  struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"pub_key"`
			VotingPower string `json:"voting_power"`
		} `json:"validator_info"`
	} `json:"result"`
}

// Client provices rpc interfaces.
type Client interface {
	GetStatus() (*Status, error)
}

type heimdallClient struct {
}

func (c *heimdallClient) GetStatus() (*Status, error) {
	req, err := http.NewRequest(http.MethodGet, localRPCAddr+"/status", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed %s", string(rawBody))
	}

	status, err := c.parseRawStatus(rawBody)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (c *heimdallClient) parseRawStatus(content []byte) (*Status, error) {
	d := Status{}
	err := json.Unmarshal(content, &d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// NewClient creates client for RPC.
func NewClient() Client {
	return &heimdallClient{}
}
