package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/dsrvlabs/vatz-plugin-common/plugins/cosmos-sdk-blocksync/mocks"
	"github.com/dsrvlabs/vatz-plugin-common/plugins/cosmos-sdk-blocksync/policy"
	rpcCosmos "github.com/dsrvlabs/vatz-plugin-common/rpc/cosmos"
	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
)

func TestPlugin(t *testing.T) {
	tests := []struct {
		Desc            string
		MockAlertStatus policy.AlertStatus
		ExpectSeverity  pluginpb.SEVERITY
		ExpectState     pluginpb.STATE
	}{
		{
			Desc:            "Unknown test",
			MockAlertStatus: policy.AlertStatusUnknown,
			ExpectSeverity:  pluginpb.SEVERITY_UNKNOWN,
			ExpectState:     pluginpb.STATE_NONE,
		},
		{
			Desc:            "Alert!",
			MockAlertStatus: policy.AlertStatusAlert,
			ExpectSeverity:  pluginpb.SEVERITY_CRITICAL,
			ExpectState:     pluginpb.STATE_FAILURE,
		},
	}

	for _, test := range tests {
		// Prepare mock.
		mockCollector := mocks.Collector{}
		mockEstimator := mocks.Estimator{}

		collector = &mockCollector
		estimator = &mockEstimator

		// Mock expects
		mockHistories := []*rpcCosmos.Status{}

		mockCollector.On("GetHistories").Return(mockHistories)
		mockEstimator.On("Estimate", mockHistories, estimateTimeWindow, blockIntervalThreshold).
			Return(test.MockAlertStatus, nil)

		// Test
		info := map[string]*structpb.Value{}
		resp, err := pluginFeature(info, nil)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, test.ExpectSeverity, resp.Severity)
		assert.Equal(t, test.ExpectState, resp.State)
	}
}
