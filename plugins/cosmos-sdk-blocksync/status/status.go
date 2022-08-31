package status

import (
	"time"

	rpcCosmos "github.com/dsrvlabs/vatz-plugin-common/rpc/cosmos"
	"github.com/rs/zerolog/log"
)

// Collector provides start/stop features for collectors.
type Collector interface {
	Start(maxHistory int, interval time.Duration)
	Stop()

	GetHistories() []*rpcCosmos.Status
}

type statusCollector struct {
	chanTerminate chan bool

	maxHistory int
	history    []*rpcCosmos.Status

	rpcClient rpcCosmos.Client
}

func (c *statusCollector) Start(maxHistory int, interval time.Duration) {
	log.Info().Str("main", "statusCollector").Msg("Start")

	c.chanTerminate = make(chan bool, 1)
	c.maxHistory = maxHistory
	c.history = make([]*rpcCosmos.Status, 0)

	go func() {
		for {
			select {
			case <-time.Tick(interval):
				status, err := c.rpcClient.GetStatus()
				if err != nil {
					// TODO: How to notify this error?
					log.Error().Str("main", "statusCollector").Msg(err.Error())
					continue
				}

				if len(c.history) > 0 {
					lastBlockHeight := c.history[len(c.history)-1].Result.SyncInfo.LatestBlockHeight
					if lastBlockHeight == status.Result.SyncInfo.LatestBlockHeight {
						continue
					}
				}

				c.history = append(c.history, status)
				if len(c.history) > c.maxHistory {
					c.history = c.history[1:]
				}
			case <-c.chanTerminate:
				log.Warn().Str("main", "statusCollector").Msg("Terminate")
				return
			}
		}
	}()
}

func (c *statusCollector) Stop() {
	log.Warn().Str("main", "statusCollector").Msg("Stop")

	c.chanTerminate <- true
}

func (c *statusCollector) GetHistories() []*rpcCosmos.Status {
	return c.history
}

// NewStatusCollector creates new collector for status.
func NewStatusCollector() Collector {
	return &statusCollector{
		rpcClient: rpcCosmos.NewClient(),
	}
}
