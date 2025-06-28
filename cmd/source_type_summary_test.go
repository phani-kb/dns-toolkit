package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSourceTypesSummaryCommand(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, sourceTypesSummaryCmd)
	assert.Equal(t, "sts", sourceTypesSummaryCmd.Use)
	assert.Contains(t, sourceTypesSummaryCmd.Short, "source types summary")
	assert.NotNil(t, sourceTypesSummaryCmd.Run)
}
