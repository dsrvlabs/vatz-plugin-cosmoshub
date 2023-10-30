package rpc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type validatorStatus struct {
	Validator struct {
		OperatorAddress string `json:"operator_address"`
		ConsensusPubkey struct {
			Type string `json:"@type"`
			Key  string `json:"key"`
		} `json:"consensus_pubkey"`
		Jailed          bool   `json:"jailed"`
		Status          string `json:"status"`
		Tokens          string `json:"tokens"`
		DelegatorShares string `json:"delegator_shares"`
		Description     struct {
			Moniker         string `json:"moniker"`
			Identity        string `json:"identity"`
			Website         string `json:"website"`
			SecurityContact string `json:"security_contact"`
			Details         string `json:"details"`
		} `json:"description"`
		UnbondingHeight string    `json:"unbonding_height"`
		UnbondingTime   time.Time `json:"unbonding_time"`
		Commission      struct {
			CommissionRates struct {
				Rate          string `json:"rate"`
				MaxRate       string `json:"max_rate"`
				MaxChangeRate string `json:"max_change_rate"`
			} `json:"commission_rates"`
			UpdateTime time.Time `json:"update_time"`
		} `json:"commission"`
		MinSelfDelegation string `json:"min_self_delegation"`
	} `json:"validator"`
}

const (
	// UNSPECIFIED defines an invalid validator status.
	Unspecified = "BOND_STATUS_UNSPECIFIED"
	// UNBONDED defines a validator that is not bonded.
	Unbonded = "BOND_STATUS_UNBONDED"
	// UNBONDING defines a validator that is unbonding.
	Unbonding = "BOND_STATUS_UNBONDING"
	// BONDED defines a validator that is bonded.
	Bonded = "BOND_STATUS_BONDED"
)

func GetBondStatus(rpcAddr string, valoperAddr string) (bool, error) {
	req, err := http.NewRequest(http.MethodGet, rpcAddr+"/cosmos/staking/v1beta1/validators/"+valoperAddr, nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("request failed %d %s", resp.StatusCode, string(rawBody))
	}

	defer resp.Body.Close()

	status := validatorStatus{}
	err = json.Unmarshal(rawBody, &status)

	if err != nil {
		return false, err
	}

	return status.Validator.Status == Bonded, nil
}
