package calendars

import (
	"time"
)

// Iceland implements a generic Icelandic exchange holiday calendar.
//
// The holidays (apart from weekends) are:
// New Year's Day (possibly moved to Monday), Maundy (Holy) Thursday, Good Friday,
// Easter Monday, First day of Summer (third or fourth Thursday in April),
// Labour Day, Ascension Day, Whit (Pentecost) Monday, Independence Day (June 17th),
// Commerce Day (first Monday in August), Christmas Eve, Christmas Day, Boxing Day,
// New Year's Eve.
//
// Valid for the following ISO 10383 Market Identifier Codes:
// XICE, FNIS.
//
// See https://www.marketholidays.com/HolidaysByCategory.aspx
//
type Iceland struct{}

//nolint:cyclop,gocognit
// IsHoliday implements Calendarer interface.
func (Iceland) IsHoliday(t time.Time) bool {
	dow := t.Weekday()
	if dow == time.Saturday || dow == time.Sunday {
		return true
	}

	y, m, d := t.Date()

	switch {
	case
		// New Year's Day (possibly moved to Monday).
		m == time.January && (d == 1 || ((d == 2 || d == 3) && dow == time.Monday)),

		// First day of Summer (third or fourth Thursday in April).
		m == time.April && (d >= 19 && d <= 25 && dow == time.Thursday),

		// Labour Day.
		m == time.May && d == 1,

		// Independence Day (June 17th).
		m == time.June && d == 17,

		// Commerce Day.
		m == time.August && d <= 7 && dow == time.Monday,

		// Christmas Eve, Christmas Day, Boxing Day, New Year's Eve.
		m == time.December && (d == 24 || d == 25 || d == 26 || d == 31):
		return true
	}

	return computusMaundyFridayMondayAscensionPentecost(y, t.YearDay())
}
