package calendars

//nolint:gofumpt
import (
	"time"

	"mbg/trading/time/computus"
)

// Finland implements a generic Finnish exchange holiday calendar.
//
// The holidays (apart from weekends) are:
// New Year's Day, Epiphany Day (Twelfth Night, Driekoningen, January 6th),
// Good Friday, Easter Monday, Ascension Day, Labour Day, Midsummer Eve (Friday
//  between June 19-25), Independence Day (December 6th), Christmas Eve,
// Christmas Day, Boxing Day, New Year's Eve.
//
// Valid for the following ISO 10383 Market Identifier Codes:
// XHEL,FNFE.
//
// See https://www.marketholidays.com/HolidaysByCategory.aspx
//
type Finland struct{}

//nolint:cyclop,gocognit
// IsHoliday implements Calendarer interface.
func (Finland) IsHoliday(t time.Time) bool {
	dow := t.Weekday()
	if dow == time.Saturday || dow == time.Sunday {
		return true
	}

	y, m, d := t.Date()

	switch {
	case
		// New Year's Day, Epiphany.
		m == time.January && (d == 1 || d == 6),

		// Labour Day.
		m == time.May && d == 1,

		// Midsummer Eve (Friday between June 19-25).
		m == time.June && dow == time.Friday && d >= 19 && d <= 25,

		// Independence Day, Christmas Eve, Christmas Day, Boxing Day.
		m == time.December && (d == 6 || d == 24 || d == 25 || d == 26 || d == 31):
		return true
	}

	es, err := computus.EasterSundayYearDay(y)
	if err != nil {
		// This can only happen if year is less then 1583.
		return false
	}

	yd := t.YearDay()

	return yd == es-2 || // Good Friday.
		yd == es+1 || // Easter Monday.
		yd == es+39 // Ascension Day, 39 days after Easter.
}
