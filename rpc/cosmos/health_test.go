package rpc

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	httpmock.Activate()

	httpmock.RegisterResponder(
		http.MethodGet,
		"http://localhost:26657/health",
		httpmock.NewStringResponder(http.StatusOK, fixtureHealth),
	)
	status, err := GetHealth("http://localhost:26657")

	assert.Nil(t, err)
	assert.Equal(t, 200, status)
	assert.False(t, status == 500, true)
}
