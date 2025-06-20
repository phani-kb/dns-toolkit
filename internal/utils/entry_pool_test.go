package utils

import (
	"strings"
	"sync"
	"testing"

	"github.com/phani-kb/dns-toolkit/internal/constants"
	"github.com/stretchr/testify/assert"
)

func TestNewDTEntryPool(t *testing.T) {
	t.Parallel()

	pool := NewDTEntryPool()
	assert.NotNil(t, pool)
	assert.Equal(t, 0, pool.Size())

	stats := pool.Stats()
	assert.Equal(t, int64(0), stats.hits)
	assert.Equal(t, int64(0), stats.misses)
}

func TestDTEntryPool_Intern(t *testing.T) {
	t.Parallel()

	pool := NewDTEntryPool()

	str1 := "example.com"
	interned1 := pool.Intern(str1)
	assert.Equal(t, str1, interned1)
	assert.Equal(t, 1, pool.Size())

	interned2 := pool.Intern(str1)
	assert.Equal(t, str1, interned2)
	assert.Equal(t, 1, pool.Size()) // Size shouldn't increase

	str2 := "google.com"
	interned3 := pool.Intern(str2)
	assert.Equal(t, str2, interned3)
	assert.Equal(t, 2, pool.Size())

	interned4 := pool.Intern(str1)
	assert.Equal(t, interned1, interned4)
}

func TestDTEntryPool_Intern_EdgeCases(t *testing.T) {
	t.Parallel()

	pool := NewDTEntryPool()

	shortStr := "a"
	interned := pool.Intern(shortStr)
	assert.Equal(t, shortStr, interned)
	// Should not be added to pool if below minimum length

	longStr := strings.Repeat("a", 1000)
	interned = pool.Intern(longStr)
	assert.Equal(t, longStr, interned)
	// Should not be added to pool if above maximum length

	emptyStr := ""
	interned = pool.Intern(emptyStr)
	assert.Equal(t, emptyStr, interned)
}

func TestDTEntryPool_InternMany(t *testing.T) {
	t.Parallel()

	pool := NewDTEntryPool()

	strings := []string{"example.com", "google.com", "example.com", "github.com"}
	interned := pool.InternMany(strings)

	assert.Len(t, interned, 4)
	assert.Equal(t, "example.com", interned[0])
	assert.Equal(t, "google.com", interned[1])
	assert.Equal(t, "example.com", interned[2])
	assert.Equal(t, "github.com", interned[3])

	assert.Equal(t, interned[0], interned[2]) // Same string should be same object

	assert.Equal(t, 3, pool.Size()) // Should have 3 unique strings
}

func TestDTEntryPool_Size(t *testing.T) {
	t.Parallel()

	pool := NewDTEntryPool()
	assert.Equal(t, 0, pool.Size())

	pool.Intern("test123") // 7 characters
	assert.Equal(t, 1, pool.Size())

	pool.Intern("example.com") // 11 characters
	assert.Equal(t, 2, pool.Size())

	pool.Intern("test123")          // Duplicate
	assert.Equal(t, 2, pool.Size()) // Size shouldn't change
}

func TestDTEntryPool_Clear(t *testing.T) {
	t.Parallel()

	pool := NewDTEntryPool()

	pool.Intern("test123")
	pool.Intern("example.com")
	pool.Intern("github.com")
	assert.Equal(t, 3, pool.Size())

	pool.Clear()
	assert.Equal(t, 0, pool.Size())

	stats := pool.Stats()
	assert.Equal(t, int64(0), stats.hits)
	assert.Equal(t, int64(0), stats.misses)
}

func TestDTEntryPool_Stats(t *testing.T) {
	t.Parallel()

	pool := NewDTEntryPool()

	stats := pool.Stats()
	assert.Equal(t, int64(0), stats.hits)
	assert.Equal(t, int64(0), stats.misses)

	pool.Intern("test1")
	stats = pool.Stats()

	pool.Intern("test1")
	stats = pool.Stats()

	pool.Intern("test2")
	stats = pool.Stats()
}

func TestDTEntryPool_ConcurrentAccess(t *testing.T) {
	t.Parallel()

	pool := NewDTEntryPool()
	const numGoroutines = 10
	const numOperations = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				pool.Intern("common.com")
				pool.Intern("unique" + string(rune(id)) + string(rune(j)) + ".com")
			}
		}(i)
	}

	wg.Wait()

	assert.True(t, pool.Size() > 0)

	result := pool.Intern("final-test.com")
	assert.Equal(t, "final-test.com", result)
}

func TestDTEntryPool_WithConstants(t *testing.T) {
	t.Parallel()

	pool := NewDTEntryPool()

	normalStr := "example.com" // This should be within normal length bounds
	interned := pool.Intern(normalStr)
	assert.Equal(t, normalStr, interned)

	strings := []string{"google.com", "github.com", "stackoverflow.com"}
	for _, s := range strings {
		interned := pool.Intern(s)
		assert.Equal(t, s, interned)
	}

	assert.True(t, pool.Size() > 0)
}

func TestConstants(t *testing.T) {
	t.Parallel()

	assert.Greater(t, constants.MaxEntryLength, 0)
	assert.Greater(t, constants.EntryMinCharLength, 0)
	assert.LessOrEqual(t, constants.EntryMinCharLength, constants.MaxEntryLength)
}
