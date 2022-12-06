package rpc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Governance struct {
	Proposal struct {
		ProposalID string `json:"proposal_id"`
		Content    struct {
			Type                string      `json:"@type"`
			Title               string      `json:"title"`
			Description         string      `json:"description"`
			Ticker              string      `json:"ticker"`
			BaseDenom           string      `json:"base_denom"`
			QuoteDenom          string      `json:"quote_denom"`
			MinPriceTickSize    string      `json:"min_price_tick_size"`
			MinQuantityTickSize string      `json:"min_quantity_tick_size"`
			MakerFeeRate        interface{} `json:"maker_fee_rate"`
			TakerFeeRate        interface{} `json:"taker_fee_rate"`
		} `json:"content"`
		Status           string `json:"status"`
		FinalTallyResult struct {
			Yes        string `json:"yes"`
			Abstain    string `json:"abstain"`
			No         string `json:"no"`
			NoWithVeto string `json:"no_with_veto"`
		} `json:"final_tally_result"`
		SubmitTime     time.Time `json:"submit_time"`
		DepositEndTime time.Time `json:"deposit_end_time"`
		TotalDeposit   []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"total_deposit"`
		VotingStartTime time.Time `json:"voting_start_time"`
		VotingEndTime   time.Time `json:"voting_end_time"`
	} `json:"proposal"`
}

type VoterVote struct {
	Vote struct {
		ProposalID string `json:"proposal_id"`
		Voter      string `json:"voter"`
		Option     string `json:"option"`
		Options    []struct {
			Option string `json:"option"`
			Weight string `json:"weight"`
		} `json:"options"`
	} `json:"vote"`
}

func GetProposal(apiPort uint, prop uint) (string, time.Time, error) {
	url := fmt.Sprintf("http://localhost:%d/cosmos/gov/v1beta1/proposals/%d", apiPort, prop)
	//url := fmt.Sprintf("https://api.cosmos.network/cosmos/gov/v1beta1/proposals/%d", prop)
	//fmt.Println(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err.Error(), time.Now(), err
	}

	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err.Error(), time.Now(), err
	}

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error(), time.Now(), err
	}

	if resp.StatusCode != http.StatusOK {
		return strconv.Itoa(resp.StatusCode), time.Now(), fmt.Errorf("request failed %s", string(rawBody))
	}

	gov := Governance{}
	err = json.Unmarshal(rawBody, &gov)

	if err != nil {
		return err.Error(), time.Now(), err
	}

	return gov.Proposal.ProposalID, gov.Proposal.VotingEndTime, nil
}

func GetVoterVote(apiPort uint, prop uint, voterAddr string) (string, error) {
	url := fmt.Sprintf("http://localhost:%d/cosmos/gov/v1beta1/proposals/%d/votes/%s", apiPort, prop, voterAddr)
	//fmt.Println(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	vote := VoterVote{}
	err = json.Unmarshal(rawBody, &vote)

	if err != nil {
		return err.Error(), err
	}

	return vote.Vote.Option, nil
}
