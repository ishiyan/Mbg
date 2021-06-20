//nolint:testpackage
package calendars

//nolint:gci
import (
	"mbg/trading/time/computus"
	"mbg/trading/time/holidays"
	"testing"
	"time"
)

const (
	testYearStart        = 1581
	testYearEnd          = 2100
	testDays             = (testYearEnd - testYearStart) * 365
	maundyThursday       = "Maundy (Holy) Thursday"
	ascensionDay         = "Ascension Day"
	dayAfterAscensionDay = "Day after Ascension Day"
	whitMonday           = "Whit (Pentecost) Monday"
	weekend              = "Weekend"
	workday              = "Workday"
	goodFriday           = "Good Friday"
	easterMonday         = "Easter Monday"
	labourDay            = "Labour Day"
	firstSummerDay       = "First day of Summer"
	midsummerEve         = "Midsummer Eve"
	nationalDay          = "National Day"
	independenceDay      = "Independence Day"
	commerceDay          = "Commerce Day"
	christmasEve         = "Christmas Eve"
	christmasDay         = "Christmas Day"
	boxingDay            = "Boxing Day"
	newYearsEve          = "New Year's Eve"
	newYearsDay          = "New Year's Day"
	epiphanyDay          = "Epiphany Day (Twelfth Night)"
	noHolidays           = "No holidays"
	constitutionDay      = "Constitution Day"
	commonPrayerDay      = "Common Prayer Day"
	dateFmt              = "Mon, Jan 2, 2006"
)

// Creates a date from a year, a month and a day.
func date(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, &time.Location{})
}

func easterSundayDate(year int) time.Time {
	tm, err := computus.EasterSunday(year)
	if err != nil {
		return date(0, 1, 1)
	}

	return tm
}

func whitMondayDate(year int) time.Time {
	es := easterSundayDate(year)

	// Whit Monday (Pfingstmontag), 50 days after Easter.
	return es.AddDate(0, 0, 50)
}

func verify(t *testing.T, c holidays.Calendarer, tm time.Time, delta int, exp bool, des string) {
	t.Helper()

	if delta != 0 {
		tm = tm.AddDate(0, 0, delta)
	}

	if act := c.IsHoliday(tm); act != exp {
		t.Errorf("IsHoliday('%v'): %v: expected %v, actual %v", tm.Format(dateFmt), des, exp, act)
	}
}

func verifyWorkday(t *testing.T, c holidays.Calendarer, month, day int) {
	t.Helper()

	w := WeekendsOnly{}

	for i := testYearStart; i < testYearEnd; i++ {
		d := date(i, month, day)
		if b := w.IsHoliday(d); !b {
			verify(t, c, d, 0, false, workday)
		} else {
			verify(t, c, d, 0, true, weekend)
		}
	}
}

//nolint:gochecknoglobals
var always = func(y int) bool { return true }

func verifyFixedDateOrWeekend(t *testing.T, c holidays.Calendarer, month, day int, name string,
	condition func(int) bool) {
	t.Helper()

	w := WeekendsOnly{}

	for i := testYearStart; i < testYearEnd; i++ {
		d := date(i, month, day)

		var s string

		var exp bool

		if b := w.IsHoliday(d); b {
			exp = true
			s = weekend
		} else {
			exp = condition(i)
			s = name
		}

		if act := c.IsHoliday(d); act != exp {
			t.Errorf("IsHoliday('%v'): %v: expected %v, actual %v", d.Format(dateFmt), s, exp, act)
		}
	}
}

func verifyComputusOrWeekend(t *testing.T, c holidays.Calendarer, offset int, name string,
	condition func(int) bool) {
	t.Helper()

	w := WeekendsOnly{}

	for i := testYearStart; i < testYearEnd; i++ {
		es := easterSundayDate(i)

		var s string

		var exp bool

		if b := w.IsHoliday(es); b {
			exp = true
			s = weekend
		} else {
			exp = condition(i)
			s = name
		}

		if act := c.IsHoliday(es); act != exp {
			t.Errorf("IsHoliday('%v'): %v: expected %v, actual %v", es.Format(dateFmt), s, exp, act)
		}
	}
}

func verifyComputus(t *testing.T, c holidays.Calendarer, offset int, name string,
	condition func(int) bool) {
	t.Helper()

	for i := testYearStart; i < testYearEnd; i++ {
		es := easterSundayDate(i)
		exp := condition(i)

		if act := c.IsHoliday(es); act != exp {
			t.Errorf("IsHoliday('%v'): %v: expected %v, actual %v", es.Format(dateFmt), name, exp, act)
		}
	}
}
