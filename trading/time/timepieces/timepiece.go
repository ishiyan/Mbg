package timepieces

import (
	"time"
)

// Timepiece provides the current time, day- and session events and reminder management.
type Timepiece interface {
	// Notifies at the beginning of a new day.
	// event Action<DateTime> NewDay

	// Notifies at the beginning of a new day session.
	// event Action<DateTime> DaySessionBegin

	// Notifies at the end of a day session.
	// event Action<DateTime> DaySessionEnd

	// Now gets the current date and time.
	Now() time.Time

	// IsHoliday indicates whether the current date is weekend or a holiday.
	IsHoliday() bool

	// AddReminder adds a reminder action at a given absolute time.
	AddReminder(string, func(), time.Time)

	// RemoveReminder removes a first occurrence of a not executed reminder action.
	RemoveReminder(string, func())
}
