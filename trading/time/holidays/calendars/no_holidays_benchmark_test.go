//nolint:testpackage
package calendars

import (
	"testing"
)

func BenchmarkIsHolidayNoHolidays(b *testing.B) {
	minDate := date(1999, 1, 1)
	maxDate := date(2022, 1, 1)
	d := minDate
	c := NoHolidays{}

	for i := 0; i < b.N; i++ {
		_ = c.IsHoliday(d)

		if d == maxDate {
			d = minDate
		} else {
			d = d.AddDate(0, 0, 1)
		}
	}
}
