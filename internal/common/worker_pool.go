package common

import (
	"runtime"
	"sync"
)

// DTWorkerPool provides a fixed-size pool of workers for parallel task processing
type DTWorkerPool struct {
	wg        sync.WaitGroup
	semaphore chan struct{} // Used to limit concurrency
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

// Submit submits a task to the pool and blocks if the pool is at capacity
func (p *DTWorkerPool) Submit(task func()) {
	p.wg.Add(1)
	p.semaphore <- struct{}{} // Acquire semaphore

	go func() {
		defer p.wg.Done()
		defer func() { <-p.semaphore }() // Release semaphore

		task()
	}()
}

// Wait blocks until all submitted tasks are completed
func (p *DTWorkerPool) Wait() {
	p.wg.Wait()
}
