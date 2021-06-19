package calendars

import (
	"time"
)

// TARGET is the 'Trans-european Automated Real-time Gross settlement Express Transfer' holiday calendar.
//
// The holidays (apart from weekends) are:
// New Year's Day, Good Friday, Easter Monday, Labour Day, Christmas Day, Boxing Day.
type TARGET struct{}

//nolint:cyclop,gomnd
// IsHoliday implements Calendarer interface.
func (TARGET) IsHoliday(t time.Time) bool {
	if checkWeekend(t) {
		return true
	}

	y, m, d := t.Date()

	switch {
	case
		// New Year's Day.
		d == 1 && m == time.January,

		// Labour Day.
		d == 1 && m == time.May && y >= 2000,

		// Christmas Day.
		d == 25 && m == time.December,

		// Boxing Day.
		d == 26 && m == time.December && y >= 2000,

		// December 31st, 1998, 1999, and 2001 only.
		d == 31 && m == time.December && (y == 1998 || y == 1999 || y == 2001):
		return true
	}

	// Good Friday and Easter Monday are holidays only after 1999.
	if y < 2000 {
		return false
	}

	return computusFridayMonday(y, t.YearDay())
}
