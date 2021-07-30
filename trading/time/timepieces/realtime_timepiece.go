package timepieces

import (
	"time"
)

// RealtimeTimepiece is a local real-time timepiece.
type RealtimeTimepiece struct {
	// Now gets the current date and time.
	now time.Time
}
