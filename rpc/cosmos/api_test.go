package common

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
}

func TestParse(t *testing.T) {
	c := heimdallClient{}
	status, err := c.parseRawStatus([]byte(fixtureStatus))

	assert.Nil(t, err)
	assert.Equal(t, false, status.Result.SyncInfo.CachingUp)
	assert.Equal(t, "9869741", status.Result.SyncInfo.LatestBlockHeight)
}

func TestGetStatus(t *testing.T) {
	httpmock.Activate()

	httpmock.RegisterResponder(
		http.MethodGet,
		"http://localhost:26657/status",
		httpmock.NewStringResponder(http.StatusOK, fixtureStatus),
	)

	c := NewClient()

	status, err := c.GetStatus()

	assert.Nil(t, err)
	assert.Equal(t, "9869741", status.Result.SyncInfo.LatestBlockHeight)
	assert.False(t, status.Result.SyncInfo.CachingUp, true)
}
