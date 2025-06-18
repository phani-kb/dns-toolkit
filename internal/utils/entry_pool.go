package utils

import (
	"sync"

	"github.com/phani-kb/dns-toolkit/internal/constants"
)

// DTEntryPool is a string interning pool that allows for efficient storage and retrieval of DNS entries.
type DTEntryPool struct {
	pool  map[string]string
	mu    sync.RWMutex
	stats PoolStats
}

// PoolStats tracks performance statistics for the entry pool.
type PoolStats struct {
	hits   int64 // Number of successful lookups (strings already in pool)
	misses int64 // Number of strings added to the pool
}

// NewDTEntryPool creates a new string interning pool.
func NewDTEntryPool() *DTEntryPool {
	return &DTEntryPool{
		pool: make(map[string]string),
	}
}

// Intern adds a string to the pool if it is not already present and returns the interned string.
func (p *DTEntryPool) Intern(s string) string {
	if len(s) < constants.EntryMinCharLength || len(s) > constants.MaxEntryLength {
		return s
	}

	p.mu.RLock()
	interned, ok := p.pool[s]
	p.mu.RUnlock()

	if ok {
		return interned
	}

	p.mu.Lock()
	// Double-check after acquiring write lock
	interned, ok = p.pool[s]
	if ok {
		p.mu.Unlock()
		return interned
	}
	p.pool[s] = s
	p.mu.Unlock()
	return s
}

// InternMany interns multiple strings in a batch
func (p *DTEntryPool) InternMany(strings []string) []string {
	result := make([]string, len(strings))
	for i, s := range strings {
		result[i] = p.Intern(s)
	}
	return result
}

// Size returns the number of strings in the pool.
func (p *DTEntryPool) Size() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return len(p.pool)
}

// Clear empties the pool.
func (p *DTEntryPool) Clear() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pool = make(map[string]string)
	p.stats.hits = 0
	p.stats.misses = 0
}

// Stats return the current hit/miss statistics for the pool.
func (p *DTEntryPool) Stats() PoolStats {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.stats
}
