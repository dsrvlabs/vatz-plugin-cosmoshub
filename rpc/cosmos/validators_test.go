package rpc

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetBondStatus(t *testing.T) {
	const valoperAddr = "valoperAddress"
	httpmock.Activate()

	httpmock.RegisterResponder(
		http.MethodGet,
		"http://localhost:1317/cosmos/staking/v1beta1/validators/"+valoperAddr,
		httpmock.NewStringResponder(http.StatusOK, fixtureBondStatus),
	)

	status, err := GetBondStatus("http://localhost:1317", valoperAddr)

	assert.Nil(t, err)
	assert.Equal(t, true, status)
}
