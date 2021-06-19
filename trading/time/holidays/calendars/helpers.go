package calendars

//nolint:gci
import (
	"mbg/trading/time/computus"
	"time"
)

// Checks if a date is a weekend.
func checkWeekend(t time.Time) bool {
	dow := t.Weekday()

	return dow == time.Saturday || dow == time.Sunday
}

// Checks if a date is the Good Friday or Easter Monday.
// Returns false before 1583.
func computusFridayMonday(y, doy int) bool {
	es, err := computus.EasterSundayYearDay(y)
	if err != nil {
		// This can only happen if year is less then 1583.
		return false
	}

	return doy == es-2 || doy == es+1
}

// Checks for Maundy Thursday, Good Friday, Easter Monday, Ascension Day, Whit (Pentecost) Monday.
// Returns false before 1583.
func computusMaundyFridayMondayAscensionPentecost(y, yd int) bool {
	es, err := computus.EasterSundayYearDay(y)
	if err != nil {
		// This can only happen if year is less then 1583.
		return false
	}

	return yd == es-3 || // Maundy Thursday (Holy Thursday), 3 days before Easter.
		yd == es-2 || // Good Friday.
		yd == es+1 || // Easter Monday.
		yd == es+39 || // Ascension Day, 39 days after Easter.
		yd == es+50 // Whit (Pentecost) Monday, 50 days after Easter.
}

/*
// Checks if a date is the Good Friday or Easter Monday.
func checkEaasterFridayMonday(y, doy int) (bool, bool, error) {
	es, err := computus.EasterSundayYearDay(y)
	if err != nil {
		return false, false, fmt.Errorf("couldn't calculate Easter Sunday: %w", err)
	}

	return doy == es-2, doy == es+1, nil
}
*/
