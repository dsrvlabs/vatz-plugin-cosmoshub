package rpc

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

// GetHealth is cosmos health chcke function
func GetHealth(rpcAddr string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, rpcAddr+"/health", nil)
	if err != nil {
		log.Error().Str("Request", "Error").Msg(fmt.Sprintf("%v", err))
		return -1, err
	}

	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Str("Response", "Error").Msg(fmt.Sprintf("%v", err))
		return -1, err
	}
	return resp.StatusCode, nil
}
