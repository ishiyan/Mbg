//nolint:testpackage
package calendars

import (
	"testing"
	"time"
)

//nolint:funlen
func TestIsHolidayEuroNext(t *testing.T) {
	t.Parallel()

	c := EuroNext{}

	wellKnown := []struct {
		t time.Time
		s string
	}{
		{date(1999, 1, 1), newYearsDay},
		{date(1999, 4, 2), goodFriday},
		{date(1999, 4, 5), easterMonday},
		{date(1999, 5, 24), whitMonday},

		{date(2000, 4, 21), goodFriday},
		{date(2000, 4, 24), easterMonday},
		{date(2000, 5, 1), labourDay},
		{date(2000, 6, 12), whitMonday},
		{date(2000, 12, 25), christmasDay},
		{date(2000, 12, 26), boxingDay},

		{date(2001, 1, 1), newYearsDay},
		{date(2001, 4, 13), goodFriday},
		{date(2001, 4, 16), easterMonday},
		{date(2001, 5, 1), labourDay},
		{date(2001, 6, 4), whitMonday},
		{date(2001, 12, 25), christmasDay},
		{date(2001, 12, 26), boxingDay},
		{date(2001, 12, 31), "Last trading day - change over to euro"},

		{date(2002, 1, 1), newYearsDay},
		{date(2002, 3, 29), goodFriday},
		{date(2002, 4, 1), easterMonday},
		{date(2002, 5, 1), labourDay},
		{date(2002, 12, 25), christmasDay},
		{date(2002, 12, 26), boxingDay},

		{date(2003, 1, 1), newYearsDay},
		{date(2003, 4, 18), goodFriday},
		{date(2003, 4, 21), easterMonday},
		{date(2003, 5, 1), labourDay},
		{date(2003, 12, 25), christmasDay},
		{date(2003, 12, 26), boxingDay},

		{date(2004, 1, 1), newYearsDay},
		{date(2004, 4, 9), goodFriday},
		{date(2004, 4, 12), easterMonday},

		{date(2005, 3, 25), goodFriday},
		{date(2005, 3, 28), easterMonday},
		{date(2005, 12, 26), boxingDay},

		{date(2006, 4, 14), goodFriday},
		{date(2006, 4, 17), easterMonday},
		{date(2006, 5, 1), labourDay},
		{date(2006, 12, 25), christmasDay},
		{date(2006, 12, 26), boxingDay},

		{date(2007, 1, 1), newYearsDay},
		{date(2007, 4, 6), goodFriday},
		{date(2007, 4, 9), easterMonday},
		{date(2007, 5, 1), labourDay},
		{date(2007, 12, 25), christmasDay},
		{date(2007, 12, 26), boxingDay},

		{date(2008, 1, 1), newYearsDay},
		{date(2008, 3, 21), goodFriday},
		{date(2008, 3, 24), easterMonday},
		{date(2008, 5, 1), labourDay},
		{date(2008, 12, 25), christmasDay},
		{date(2008, 12, 26), boxingDay},

		{date(2009, 1, 1), newYearsDay},
		{date(2009, 4, 10), goodFriday},
		{date(2009, 4, 13), easterMonday},
		{date(2009, 5, 1), labourDay},
		{date(2009, 12, 25), christmasDay},

		{date(2010, 1, 1), newYearsDay},
		{date(2010, 4, 2), goodFriday},
		{date(2010, 4, 5), easterMonday},

		{date(2011, 4, 22), goodFriday},
		{date(2011, 4, 25), easterMonday},
		{date(2011, 12, 26), boxingDay},

		{date(2012, 4, 6), goodFriday},
		{date(2012, 4, 9), easterMonday},
		{date(2012, 5, 1), labourDay},
		{date(2012, 12, 25), christmasDay},
		{date(2012, 12, 26), boxingDay},

		{date(2013, 1, 1), newYearsDay},
		{date(2013, 3, 29), goodFriday},
		{date(2013, 4, 1), easterMonday},
		{date(2013, 5, 1), labourDay},
		{date(2013, 12, 25), christmasDay},
		{date(2013, 12, 26), boxingDay},

		{date(2014, 1, 1), newYearsDay},
		{date(2014, 4, 18), goodFriday},
		{date(2014, 4, 21), easterMonday},
		{date(2014, 5, 1), labourDay},
		{date(2014, 12, 25), christmasDay},
		{date(2014, 12, 26), boxingDay},

		{date(2015, 1, 1), newYearsDay},
		{date(2015, 4, 3), goodFriday},
		{date(2015, 4, 6), easterMonday},
		{date(2015, 5, 1), labourDay},
		{date(2015, 12, 25), christmasDay},

		{date(2016, 1, 1), newYearsDay},
		{date(2016, 3, 25), goodFriday},
		{date(2016, 3, 28), easterMonday},
		{date(2016, 12, 26), boxingDay},

		{date(2017, 4, 14), goodFriday},
		{date(2017, 4, 17), easterMonday},
		{date(2017, 5, 1), labourDay},
		{date(2017, 12, 25), christmasDay},
		{date(2017, 12, 26), boxingDay},

		{date(2018, 1, 1), newYearsDay},
		{date(2018, 3, 30), goodFriday},
		{date(2018, 4, 2), easterMonday},
		{date(2018, 5, 1), labourDay},
		{date(2018, 12, 25), christmasDay},
		{date(2018, 12, 26), boxingDay},

		{date(2019, 1, 1), newYearsDay},
		{date(2019, 4, 19), goodFriday},
		{date(2019, 4, 22), easterMonday},
		{date(2019, 5, 1), labourDay},
		{date(2019, 12, 25), christmasDay},
		{date(2019, 12, 26), boxingDay},

		{date(2020, 1, 1), newYearsDay},
		{date(2020, 4, 10), goodFriday},
		{date(2020, 4, 13), easterMonday},
		{date(2020, 5, 1), labourDay},
		{date(2020, 12, 25), christmasDay},

		{date(2021, 1, 1), newYearsDay},
		{date(2021, 4, 2), goodFriday},
		{date(2021, 4, 5), easterMonday},
	}

	for _, tt := range wellKnown {
		verify(t, c, tt.t, 0, true, tt.s)
	}

	// Fixed dates.
	verifyFixedDateOrWeekend(t, c, 1, 1, newYearsDay, always)    // New Year's Day.
	verifyFixedDateOrWeekend(t, c, 5, 1, labourDay, always)      // Labour Day.
	verifyFixedDateOrWeekend(t, c, 12, 25, christmasDay, always) // Christmas Day.
	verifyFixedDateOrWeekend(t, c, 12, 26, boxingDay, always)    // Boxing Day.
	verifyFixedDateOrWeekend(t, c, 12, 31, newYearsEve,          // New Year's Eve.
		func(y int) bool { return y == 2001 })

	// Computus.
	verifyComputus(t, c, -2, goodFriday, always)  // Good Friday.
	verifyComputus(t, c, 1, easterMonday, always) // Easter Monday.
	verifyComputusOrWeekend(t, c, 50, whitMonday, // Whit (Pentecost) Monday, 50 days after Easter.
		func(y int) bool { return y == 1999 || y == 2000 || y == 2001 })

	verifyWorkday(t, c, 10, 2)
	verifyWorkday(t, c, 10, 3)
}
