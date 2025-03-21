package calendars

import (
	"time"
)

// Norway implements a generic Norwegian exchange holiday calendar.
//
// The holidays (apart from weekends) are:
// New Year's Day, Maundy Thursday, Good Friday, Easter Monday, Labour Day,
// Constitution (National Independence), Ascension Day, Whit (Pentecost) Monday,
// Christmas Eve, Christmas Day, Boxing Day, New Year's Eve.
//
// Valid for the following ISO 10383 Market Identifier Codes:
// XOSL, MERK, XOAS, XOBD, NOTC, NORX, NEXO.
//
// Not valid for XIMA, FISH, NORX-EUR.
//
// See https://www.marketholidays.com/HolidaysByCategory.aspx
//
type Norway struct{}

//nolint:cyclop
// IsHoliday implements Calendarer interface.
func (Norway) IsHoliday(t time.Time) bool {
	if checkWeekend(t) {
		return true
	}

	y, m, d := t.Date()

	switch {
	case
		// New Year's Day.
		m == time.January && d == 1,

		// Labour Day, Constitution (National Independence) Day.
		m == time.May && (d == 1 || d == 17),

		// Christmas Eve, Christmas Day, Boxing Day, New Year's Eve.
		m == time.December && (d == 24 || d == 25 || d == 26 || d == 31):
		return true
	}

	// Maundy Thursday, Good Friday, Easter Monday, Ascension Day, Whit (Pentecost) Monday.
	return computusMaundyFridayMondayAscensionPentecost(y, t.YearDay())
}
