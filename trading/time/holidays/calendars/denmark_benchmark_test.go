//nolint:testpackage
package calendars

import (
	"testing"
)

func BenchmarkIsHolidayDenmark(b *testing.B) {
	minDate := date(1980, 1, 1)
	maxDate := date(2050, 1, 1)
	d := minDate
	c := Denmark{}

	for i := 0; i < b.N; i++ {
		_ = c.IsHoliday(d)

		if d == maxDate {
			d = minDate
		} else {
			d = d.AddDate(0, 0, 1)
		}
	}
}

func BenchmarkIsHolidayDenmarkWorkday(b *testing.B) {
	c := Denmark{}
	d := date(2021, 10, 5)

	for i := 0; i < b.N; i++ {
		_ = c.IsHoliday(d)
	}
}
