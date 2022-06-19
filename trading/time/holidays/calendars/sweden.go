package calendars

//nolint:gofumpt
import (
	"time"

	"mbg/trading/time/computus"
)

// Sweden implements a generic Swedish exchange holiday calendar.
//
// The holidays (apart from weekends) are:
// New Year's Day, Epiphany Day (Twelfth Night, January 6th), Good Friday,
// Easter Monday, Ascension Day, Whit (Pentecost) Monday, Labour Day,
// National Day (June 6th, it has been debated whether or not this day
// should be declared as a holiday; as of 2002 the Stockholmborsen is open
// that day), Midsummer Eve (Friday between June 19-25), Christmas Eve,
// Christmas Day, Boxing Day, New Year's Eve.
//
// Valid for the following ISO 10383 Market Identifier Codes:
// XSTO, NMTF, XNGM, XNDX, FNSE, XSAT.
//
// See https://www.marketholidays.com/HolidaysByCategory.aspx
//
type Sweden struct{}

//nolint:cyclop,gocognit
// IsHoliday implements Calendarer interface.
func (Sweden) IsHoliday(t time.Time) bool {
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

		// National Day.
		y > 2002 && m == time.June && d == 6,

		// Midsummer Eve (Friday between June 19-25).
		m == time.June && dow == time.Friday && d >= 19 && d <= 25,

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

	return yd == es-2 || // Good Friday.
		yd == es+1 || // Easter Monday.
		yd == es+39 || // Ascension Day, 39 days after Easter.
		yd == es+50 // Whit (Pentecost) Monday, 50 days after Easter.
}
