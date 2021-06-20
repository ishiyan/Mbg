package calendars

//nolint:gci
import (
	"mbg/trading/time/computus"
	"time"
)

// Denmark implements a generic Danish exchange holiday calendar.
//
// The holidays (apart from weekends) are:
// New Year's Day, Maundy (Holy) Thursday, Good Friday, Easter Monday,
// General Prayer Day (26 days after Easter), Ascension Day,
// Whit (Pentecost) Monday, Constitution Day (June 5th),
// Christmas Eve, Christmas Day, Boxing Day, New Year's Eve.
//
// Valid for the following ISO 10383 Market Identifier Codes:
// DKTC, XCSE, FNDK.
//
// See https://www.marketholidays.com/HolidaysByCategory.aspx
//
type Denmark struct{}

//nolint:cyclop,gocognit
// IsHoliday implements Calendarer interface.
func (Denmark) IsHoliday(t time.Time) bool {
	if checkWeekend(t) {
		return true
	}

	y, m, d := t.Date()

	switch {
	case
		// New Year's Day.
		m == time.January && d == 1,

		// Constitution Day.
		m == time.June && d == 5,

		// Christmas Eve, Christmas Day, Boxing Day, New Year's Eve.
		m == time.December && (d == 24 || d == 25 || d == 26 || d == 31):
		return true
	}

	es, err := computus.EasterSundayYearDay(y)
	if err != nil {
		// This can only happen if year is less then 1583.
		return false
	}

	yd := t.YearDay()

	return yd == es-3 || // Maundy Thursday (Holy Thursday), 3 days before Easter.
		yd == es-2 || // Good Friday.
		yd == es+1 || // Easter Monday.
		yd == es+26 || // General Prayer Day, 26 days after Easter.
		yd == es+39 || // Ascension Day, 39 days after Easter.
		(yd == es+40 && y > 2008) || // Day after Ascension Day, 40 days after Easter.
		yd == es+50 // Whit (Pentecost) Monday, 50 days after Easter.
}
