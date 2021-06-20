//nolint:testpackage
package calendars

import (
	"testing"
)

func BenchmarkIsHolidayEuroNext(b *testing.B) {
	minDate := date(1999, 1, 1)
	maxDate := date(2022, 1, 1)
	d := minDate
	c := EuroNext{}

	for i := 0; i < b.N; i++ {
		_ = c.IsHoliday(d)

		if d == maxDate {
			d = minDate
		} else {
			d = d.AddDate(0, 0, 1)
		}
	}
}

func BenchmarkIsHolidayEuroNextWorkday(b *testing.B) {
	c := EuroNext{}
	d := date(2021, 10, 5)

	for i := 0; i < b.N; i++ {
		_ = c.IsHoliday(d)
	}
}
