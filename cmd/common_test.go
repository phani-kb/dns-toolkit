package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendSummary(t *testing.T) {
	var summaries []string

	appendSummary(&summaries, "test1", func(s string) bool { return len(s) > 0 })
	assert.Equal(t, 1, len(summaries))
	assert.Equal(t, "test1", summaries[0])

	appendSummary(&summaries, "", func(s string) bool { return len(s) > 0 })
	assert.Equal(t, 1, len(summaries))

	var intSummaries []int
	appendSummary(&intSummaries, 42, func(i int) bool { return i > 0 })
	assert.Equal(t, 1, len(intSummaries))
	assert.Equal(t, 42, intSummaries[0])

	appendSummary(&intSummaries, -1, func(i int) bool { return i > 0 })
	assert.Equal(t, 1, len(intSummaries))
}
