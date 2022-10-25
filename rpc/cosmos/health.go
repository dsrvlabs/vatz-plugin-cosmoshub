package common

import (
	"net/http"
)

// GetHealth is cosmos health chcke function
func GetHealth(rpcAddr string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, rpcAddr+"/health", nil)
	if err != nil {
		return -1, err
	}

	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	return resp.StatusCode, nil
}
