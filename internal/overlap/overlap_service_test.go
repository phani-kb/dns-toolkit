package overlap

import (
	"testing"

	c "github.com/phani-kb/dns-toolkit/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestComputeCompactSummaryFromPairs_DedupeOverlapCounts(t *testing.T) {
	pairs := []c.OverlapPair{
		{
			Source:  c.OverlapFileInfo{Name: "sourceA", Count: 100, ListType: "blocklist", Type: "domain"},
			Target:  c.OverlapFileInfo{Name: "targetX", Count: 50, ListType: "blocklist", Type: "domain"},
			Overlap: 10,
		},
		{
			Source:  c.OverlapFileInfo{Name: "targetX", Count: 50, ListType: "blocklist", Type: "domain"},
			Target:  c.OverlapFileInfo{Name: "sourceA", Count: 100, ListType: "blocklist", Type: "domain"},
			Overlap: 8,
		},
		{
			Source:  c.OverlapFileInfo{Name: "sourceA", Count: 100, ListType: "blocklist", Type: "domain"},
			Target:  c.OverlapFileInfo{Name: "targetY", Count: 40, ListType: "blocklist", Type: "domain"},
			Overlap: 5,
		},
	}

	cs := computeCompactSummaryFromPairs("domain", "sourceA", pairs)
	assert.NotNil(t, cs)
	assert.Equal(t, 2, cs.TargetsCount)
	// For targetX, the highest overlap=10, so total overlap = 10 + 5 = 15
	assert.Equal(t, 15, func() int { return cs.Count - cs.Unique }())
	// Unique should be Count - 15 = 85
	assert.Equal(t, 85, cs.Unique)
}
