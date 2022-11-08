package rpc

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetNpeers(t *testing.T) {
	httpmock.Activate()

	httpmock.RegisterResponder(
		http.MethodGet,
		"http://localhost:26657/net_info",
		httpmock.NewStringResponder(http.StatusOK, fixturePeerCount),
	)

	ret, err := GetNpeers()
	assert.Nil(t, err)
	peerCount, _ := strconv.Atoi(ret)
	assert.GreaterOrEqual(t, peerCount, 0)
}
