package status

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	rpcCosmos "github.com/dsrvlabs/vatz-plugin-common/rpc/cosmos"
	"github.com/dsrvlabs/vatz-plugin-common/rpc/cosmos/mocks"
)

func init() {
}

func TestCollector(t *testing.T) {
	// Mock
	mockStatus := rpcCosmos.Status{}
	mockStatus.Result.SyncInfo.LatestBlockHeight = "123456"
	mockStatus.Result.SyncInfo.CachingUp = false

	mockClient := mocks.Client{}
	mockClient.On("GetStatus").Return(&mockStatus, nil)

	// Test
	c := statusCollector{rpcClient: &mockClient}
	c.Start(10, time.Millisecond*100)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond * 120)
		c.Stop()
		wg.Done()
	}()

	wg.Wait()

	// Assert
	assert.Equal(t, 1, len(c.history))
	assert.Equal(t, "123456", c.history[len(c.history)-1].Result.SyncInfo.LatestBlockHeight)
}

func TestIgnoreDuplicate(t *testing.T) {
	// Mock
	mockStatus := rpcCosmos.Status{}
	mockStatus.Result.SyncInfo.LatestBlockHeight = "123456"
	mockStatus.Result.SyncInfo.CachingUp = false

	mockClient := mocks.Client{}
	mockClient.On("GetStatus").Return(&mockStatus, nil)

	c := statusCollector{rpcClient: &mockClient}

	// Test
	c.Start(10, time.Second*2)

	// Prepare: Intentionally add to simulate duplication
	c.history = append(c.history, &mockStatus)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 3)
		c.Stop()
		wg.Done()
	}()

	wg.Wait()

	// Assert
	assert.Equal(t, 1, len(c.history))
	assert.Equal(t, "123456", c.history[len(c.history)-1].Result.SyncInfo.LatestBlockHeight)
}
