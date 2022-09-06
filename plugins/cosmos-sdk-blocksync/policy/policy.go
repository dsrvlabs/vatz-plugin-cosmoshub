package policy

import (
	"strconv"
	"time"

	rpcCosmos "github.com/dsrvlabs/vatz-plugin-cosmoshub/rpc/cosmos"
	"github.com/rs/zerolog/log"
)

// Alert status.
const (
	AlertStatusNormal AlertStatus = iota
	AlertStatusAlert
	AlertStatusUnknown
)

// AlertStatus describe alert status from PolicyEstimator.
type AlertStatus int

// Estimator provides interfaces for alert estimation.
type Estimator interface {
	Estimate(history []*rpcCosmos.Status, timeWindow, maxBlockInterval time.Duration) (AlertStatus, error)
}

type blockSyncEstimator struct {
}

func (p *blockSyncEstimator) Estimate(history []*rpcCosmos.Status, timeWindow, maxBlockInterval time.Duration) (AlertStatus, error) {
	log.Info().Str("policy", "blockSyncEstimator").Msg("estimate")

	// Filter status as time range.
	statusInRange := []struct {
		TimeCreated time.Time
		BlockHeight int64
	}{}

	timeBoundary := time.Now().Add(-timeWindow)
	for _, status := range history {
		blockTimeStr := status.Result.SyncInfo.LatestBlockTime
		blockTime, err := time.Parse(time.RFC3339, blockTimeStr)
		if err != nil {
			log.Info().Str("policy", "blockSyncEstimator").Msg(err.Error())
			return AlertStatusUnknown, err
		}

		if blockTime.Before(timeBoundary) {
			continue
		}

		blockHeight, err := strconv.ParseInt(status.Result.SyncInfo.LatestBlockHeight, 10, 64)
		if err != nil {
			log.Info().Str("policy", "blockSyncEstimator").Msg(err.Error())
			return AlertStatusUnknown, err
		}

		filteredStatus := struct {
			TimeCreated time.Time
			BlockHeight int64
		}{
			TimeCreated: blockTime,
			BlockHeight: blockHeight,
		}

		statusInRange = append(statusInRange, filteredStatus)
	}

	// TODO: Restrict minimum number of samples to estimate alerts.

	// calculate block time.
	if len(history) == 0 {
		log.Info().Str("policy", "blockSyncEstimator").Msg("not enough samples")
		return AlertStatusUnknown, nil
	}

	firstStatus := statusInRange[0]
	lastStatus := statusInRange[len(statusInRange)-1]

	timeRange := lastStatus.TimeCreated.Sub(firstStatus.TimeCreated)
	blockInc := lastStatus.BlockHeight - firstStatus.BlockHeight

	blockInterval := float64(timeRange.Nanoseconds()) / float64(blockInc)
	if blockInterval < float64(maxBlockInterval.Nanoseconds()) {
		return AlertStatusNormal, nil
	}

	return AlertStatusAlert, nil
}

// NewEstimator creates new estimator instance.
func NewEstimator() Estimator {
	return &blockSyncEstimator{}
}
