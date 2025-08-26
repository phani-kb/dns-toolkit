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
	if cs == nil {
		t.FailNow()
	}
	assert.Equal(t, 2, cs.TargetsCount)
	// For targetX, the highest overlap=10, so total overlap = 10 + 5 = 15
	assert.Equal(t, 15, func() int { return cs.Count - cs.Unique }())
	// Unique should be Count - 15 = 85
	assert.Equal(t, 85, cs.Unique)
	// No cross-list overlaps in this test
	assert.Equal(t, 0, cs.Conflicts)
}

func TestComputeCompactSummaryFromPairs_ListTypeDiffers(t *testing.T) {
	pairs := []c.OverlapPair{
		{
			Source:  c.OverlapFileInfo{Name: "sourceA", Count: 100, ListType: "blocklist", Type: "domain"},
			Target:  c.OverlapFileInfo{Name: "targetX", Count: 50, ListType: "blocklist", Type: "domain"},
			Overlap: 10,
		},
		{
			Source:  c.OverlapFileInfo{Name: "sourceA", Count: 100, ListType: "blocklist", Type: "domain"},
			Target:  c.OverlapFileInfo{Name: "targetX", Count: 20, ListType: "allowlist", Type: "domain"},
			Overlap: 3,
		},
	}

	cs := computeCompactSummaryFromPairs("domain", "sourceA", pairs)
	assert.NotNil(t, cs)
	if cs == nil {
		t.FailNow()
	}
	assert.Equal(t, 2, cs.TargetsCount)
	// total overlap = 10 + 3 = 13
	assert.Equal(t, 13, func() int { return cs.Count - cs.Unique }())
	// Cross-list overlaps should be 3
	assert.Equal(t, 3, cs.Conflicts)
}

func TestComputeCompactSummaryFromPairs_TypeDiffers(t *testing.T) {
	pairs := []c.OverlapPair{
		{
			Source:  c.OverlapFileInfo{Name: "sourceA", Count: 100, ListType: "blocklist", Type: "domain"},
			Target:  c.OverlapFileInfo{Name: "targetX", Count: 50, ListType: "blocklist", Type: "domain"},
			Overlap: 10,
		},
		{
			Source:  c.OverlapFileInfo{Name: "sourceA", Count: 100, ListType: "blocklist", Type: "domain"},
			Target:  c.OverlapFileInfo{Name: "targetX", Count: 30, ListType: "blocklist", Type: "ipv4"},
			Overlap: 4,
		},
	}

	cs := computeCompactSummaryFromPairs("domain", "sourceA", pairs)
	assert.NotNil(t, cs)
	if cs == nil {
		t.FailNow()
	}
	assert.Equal(t, 2, cs.TargetsCount)
	// total overlap = 10 + 4 = 14
	assert.Equal(t, 14, func() int { return cs.Count - cs.Unique }())
	// Cross-type overlaps count as cross-list-type for conflicts
	assert.Equal(t, 4, cs.Conflicts)
}
