package common

import (
	"context"
	"runtime"
	"sync"

	"golang.org/x/time/rate"
)

// DTWorkerPool provides a fixed-size pool of workers for parallel task processing
type DTWorkerPool struct {
	limiter   *rate.Limiter // used to limit task start rate
	semaphore chan struct{} // Used to limit concurrency
	wg        sync.WaitGroup
}

// NewDTWorkerPool creates a new worker pool with specified capacity
// If maxWorkers is <= 0, it defaults to number of CPUs
func NewDTWorkerPool(maxWorkers int) *DTWorkerPool {
	if maxWorkers <= 0 {
		maxWorkers = runtime.GOMAXPROCS(0)
	}

	return &DTWorkerPool{
		semaphore: make(chan struct{}, maxWorkers),
	}
}

// NewDTWorkerPoolWithLimiter creates a worker pool with a rate limiter controlling task start rate
func NewDTWorkerPoolWithLimiter(maxWorkers int, limiter *rate.Limiter) *DTWorkerPool {
	pool := NewDTWorkerPool(maxWorkers)
	pool.limiter = limiter
	return pool
}

// Submit submits a task to the pool and blocks if the pool is at capacity
func (p *DTWorkerPool) Submit(task func()) {
	p.wg.Add(1)
	p.semaphore <- struct{}{} // Acquire semaphore

	go func() {
		defer p.wg.Done()
		defer func() { <-p.semaphore }() // Release semaphore

		if p.limiter != nil {
			if err := p.limiter.Wait(context.Background()); err != nil {
				// the limiter fails
				return
			}
		}

		task()
	}()
}

// Wait blocks until all submitted tasks are completed
func (p *DTWorkerPool) Wait() {
	p.wg.Wait()
}
