//nolint:testpackage
package calendars

import (
	"testing"
	"time"
)

//nolint:funlen
func TestIsHolidayIceland(t *testing.T) {
	t.Parallel()

	c := Iceland{}

	wellKnown := []struct {
		t time.Time
		s string
	}{
		{date(2016, 1, 1), newYearsDay},
		{date(2016, 3, 24), maundyThursday},
		{date(2016, 3, 25), goodFriday},
		{date(2016, 3, 28), easterMonday},
		{date(2016, 4, 21), firstSummerDay},
		{date(2016, 5, 5), ascensionDay},
		{date(2016, 5, 16), whitMonday},
		{date(2016, 6, 17), independenceDay},
		{date(2016, 8, 1), commerceDay},
		{date(2016, 12, 26), boxingDay},

		{date(2017, 1, 1), newYearsDay},
		{date(2017, 4, 13), maundyThursday},
		{date(2017, 4, 14), goodFriday},
		{date(2017, 4, 17), easterMonday},
		{date(2017, 4, 20), firstSummerDay},
		{date(2017, 5, 1), labourDay},
		{date(2017, 5, 25), ascensionDay},
		{date(2017, 6, 5), whitMonday},
		{date(2017, 8, 7), commerceDay},
		{date(2017, 12, 25), christmasDay},
		{date(2017, 12, 26), boxingDay},

		{date(2018, 1, 1), newYearsDay},
		{date(2018, 3, 29), maundyThursday},
		{date(2018, 3, 30), goodFriday},
		{date(2018, 4, 2), easterMonday},
		{date(2018, 4, 19), firstSummerDay},
		{date(2018, 5, 1), labourDay},
		{date(2018, 5, 10), ascensionDay},
		{date(2018, 5, 21), whitMonday},
		{date(2018, 8, 6), commerceDay},
		{date(2018, 12, 25), christmasDay},
		{date(2018, 12, 26), boxingDay},

		{date(2019, 1, 1), newYearsDay},
		{date(2019, 4, 18), maundyThursday},
		{date(2019, 4, 19), goodFriday},
		{date(2019, 4, 22), easterMonday},
		{date(2019, 4, 25), firstSummerDay},
		{date(2019, 5, 1), labourDay},
		{date(2019, 5, 30), ascensionDay},
		{date(2019, 6, 10), whitMonday},
		{date(2019, 6, 17), independenceDay},
		{date(2019, 8, 5), commerceDay},
		{date(2019, 12, 24), christmasEve},
		{date(2019, 12, 25), christmasDay},
		{date(2019, 12, 26), boxingDay},
		{date(2019, 12, 31), newYearsEve},

		{date(2020, 1, 1), newYearsDay},
		{date(2020, 4, 9), maundyThursday},
		{date(2020, 4, 10), goodFriday},
		{date(2020, 4, 13), easterMonday},
		{date(2020, 5, 1), labourDay},
		{date(2020, 5, 21), ascensionDay},
		{date(2020, 6, 1), whitMonday},
		{date(2020, 6, 17), independenceDay},
		{date(2020, 8, 3), commerceDay},
		{date(2020, 12, 24), christmasEve},
		{date(2020, 12, 25), christmasDay},
		{date(2020, 12, 31), newYearsEve},

		{date(2021, 1, 1), newYearsDay},
		{date(2021, 4, 1), maundyThursday},
		{date(2021, 4, 2), goodFriday},
		{date(2021, 4, 5), easterMonday},
		{date(2021, 4, 22), firstSummerDay},
		{date(2021, 5, 13), ascensionDay},
		{date(2021, 5, 24), whitMonday},
		{date(2021, 6, 17), independenceDay},
		{date(2021, 8, 2), commerceDay},
		{date(2021, 12, 24), christmasEve},
		{date(2021, 12, 31), newYearsEve},
	}

	for _, tt := range wellKnown {
		verify(t, c, tt.t, 0, true, tt.s)
	}

	// Fixed dates.
	verifyFixedDateOrWeekend(t, c, 5, 1, labourDay, always)        // Labour Day.
	verifyFixedDateOrWeekend(t, c, 6, 17, independenceDay, always) // Independence Day.
	verifyFixedDateOrWeekend(t, c, 12, 24, christmasEve, always)   // Christmas Eve.
	verifyFixedDateOrWeekend(t, c, 12, 25, christmasDay, always)   // Christmas Day.
	verifyFixedDateOrWeekend(t, c, 12, 26, boxingDay, always)      // Boxing Day.
	verifyFixedDateOrWeekend(t, c, 12, 31, newYearsEve, always)    // New Year's Eve.

	// Computus.
	verifyComputus(t, c, -3, maundyThursday, always)        // Maundy Thursday (Holy Thursday), 3 days before Easter.
	verifyComputus(t, c, -2, goodFriday, always)            // Good Friday.
	verifyComputus(t, c, 1, easterMonday, always)           // Easter Monday.
	verifyComputusOrWeekend(t, c, 39, ascensionDay, always) // Ascension Day, 39 days after Easter.
	verifyComputus(t, c, 50, whitMonday, always)            // Whit (Pentecost) Monday, 50 days after Easter.

	verifyWorkday(t, c, 10, 2)
}
