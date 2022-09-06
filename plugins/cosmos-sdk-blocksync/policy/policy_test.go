package policy

import (
	"fmt"
	"testing"
	"time"

	rpcCosmos "github.com/dsrvlabs/vatz-plugin-cosmoshub/rpc/cosmos"
	"github.com/stretchr/testify/assert"
)

func TestSyncPolicy(t *testing.T) {
	// Tests
	tests := []struct {
		Desc          string
		MockBlockTime time.Duration
		TestSamples   int
		ExpectStatus  AlertStatus
	}{
		{
			Desc:          "No Samples.",
			MockBlockTime: time.Second * 5,
			TestSamples:   0,
			ExpectStatus:  AlertStatusUnknown,
		},
		{
			Desc:          "Normal block time.",
			MockBlockTime: time.Second * 5,
			TestSamples:   10,
			ExpectStatus:  AlertStatusNormal,
		},
		{
			Desc:          "Block time is too slow",
			MockBlockTime: time.Second * 15,
			TestSamples:   10,
			ExpectStatus:  AlertStatusAlert,
		},
	}

	for _, test := range tests {
		// Mock status.
		statusHistory := make([]*rpcCosmos.Status, test.TestSamples)

		timeStart := time.Now().UTC().Add(-(test.MockBlockTime * time.Duration(test.TestSamples)))
		for i := 0; i < test.TestSamples; i++ {
			dummyStatus := createDummyStatus(i, timeStart.Add(test.MockBlockTime*time.Duration(i)))
			statusHistory[i] = dummyStatus
		}

		// Test
		e := blockSyncEstimator{}
		status, err := e.Estimate(statusHistory, time.Second*30, time.Second*10)

		// Asserts
		assert.Nil(t, err)
		assert.Equal(t, test.ExpectStatus, status)
	}
}

func createDummyStatus(blockNo int, blockTime time.Time) *rpcCosmos.Status {
	newStatus := rpcCosmos.Status{}
	newStatus.Result.SyncInfo.LatestBlockHeight = fmt.Sprintf("%d", blockNo)
	newStatus.Result.SyncInfo.LatestBlockTime = blockTime.Format(time.RFC3339Nano)
	return &newStatus
}
