// Package holidays provides various holiday schedules for a specific exchange or a country.
package holidays

import (
	"time"
)

// Calendarer provides an abstraction to check if a given date is a holiday.
type Calendarer interface {
	// IsHoliday checks if a given date is a holiday.
	IsHoliday(time.Time) bool
}
