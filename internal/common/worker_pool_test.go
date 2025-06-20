package common

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDTWorkerPool(t *testing.T) {
	t.Parallel()

	pool := NewDTWorkerPool(5)
	assert.NotNil(t, pool)
	assert.Equal(t, 5, cap(pool.semaphore))

	pool = NewDTWorkerPool(0)
	assert.NotNil(t, pool)
	assert.True(t, cap(pool.semaphore) > 0)
}

func TestDTWorkerPool_SubmitAndWait(t *testing.T) {
	t.Parallel()

	pool := NewDTWorkerPool(2) // 2 workers

	var counter int32

	// Submit 5 tasks
	for i := 0; i < 5; i++ {
		pool.Submit(func() {
			// Simulate work
			time.Sleep(10 * time.Millisecond)
			atomic.AddInt32(&counter, 1)
		})
	}

	pool.Wait()

	assert.Equal(t, int32(5), counter)
}

func TestDTWorkerPool_Concurrency(t *testing.T) {
	t.Parallel()

	maxWorkers := 2
	pool := NewDTWorkerPool(maxWorkers)

	var currentWorkers int32
	var maxConcurrentWorkers int32
	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		pool.Submit(func() {
			count := atomic.AddInt32(&currentWorkers, 1)

			mutex.Lock()
			if count > maxConcurrentWorkers {
				maxConcurrentWorkers = count
			}
			mutex.Unlock()

			time.Sleep(20 * time.Millisecond)

			atomic.AddInt32(&currentWorkers, -1)
		})
	}

	pool.Wait()

	assert.LessOrEqual(t, maxConcurrentWorkers, int32(maxWorkers))
}
