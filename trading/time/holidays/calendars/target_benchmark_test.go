//nolint:testpackage
package calendars

import (
	"testing"
)

func BenchmarkIsHolidayTARGET(b *testing.B) {
	minDate := date(1999, 1, 1)
	maxDate := date(2022, 1, 1)
	d := minDate
	c := TARGET{}

	for i := 0; i < b.N; i++ {
		_ = c.IsHoliday(d)

		if d == maxDate {
			d = minDate
		} else {
			d = d.AddDate(0, 0, 1)
		}
	}
}
