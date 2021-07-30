package portfolios

//nolint:gci
import (
	"sync"
)

// RoundtripPerformance is a roundtrip performance.
type RoundtripPerformance struct {
	mu sync.RWMutex
}

// NewRoundtripPerformance creates a new roundtrip portfolio performance.
// This is the only correct way to create a roundtrip performance instance.
func NewRoundtripPerformance() *RoundtripPerformance {
	p := &RoundtripPerformance{}

	return p
}
