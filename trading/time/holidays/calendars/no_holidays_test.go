//nolint:testpackage
package calendars

import (
	"testing"
)

func TestIsHolidayNoHolidays(t *testing.T) {
	t.Parallel()

	c := NoHolidays{}
	d := date(testYearStart, 1, 1)

	for i := 0; i < testDays; i++ {
		verify(t, c, d, 0, false, noHolidays)
		d = d.AddDate(0, 0, 1)
	}
}
