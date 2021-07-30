package portfolios

//nolint:gci
import (
	"sync"
)

// Performance is a portfolio performance.
type Performance struct {
	mu sync.RWMutex
}

// NewPerformance creates a new portfolio performance.
// This is the only correct way to create a performance instance.
func NewPerformance() *Performance {
	p := &Performance{}

	return p
}
