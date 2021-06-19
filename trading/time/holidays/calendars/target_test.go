//nolint:testpackage
package calendars

import (
	"testing"
)

func TestIsHolidayTARGET(t *testing.T) {
	t.Parallel()

	c := TARGET{}
	from2000 := func(y int) bool { return y >= 2000 }

	// Fixed dates.
	verifyFixedDateOrWeekend(t, c, 1, 1, newYearsDay, always)    // New Year's Day.
	verifyFixedDateOrWeekend(t, c, 5, 1, labourDay, from2000)    // Labour Day.
	verifyFixedDateOrWeekend(t, c, 12, 25, christmasDay, always) // Christmas Day.
	verifyFixedDateOrWeekend(t, c, 12, 26, boxingDay, from2000)  // Boxing Day.
	verifyFixedDateOrWeekend(t, c, 12, 31, newYearsEve,          // New Year's Eve.
		func(y int) bool { return y == 1998 || y == 1999 || y == 2001 })

	// Computus.
	verifyComputusOrWeekend(t, c, -2, goodFriday, from2000)  // Good Friday.
	verifyComputusOrWeekend(t, c, 1, easterMonday, from2000) // Easter Monday.
}
