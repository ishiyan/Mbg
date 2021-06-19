//nolint:testpackage
package calendars

import (
	"testing"
	"time"
)

func TestIsHolidayWeenendsOnly(t *testing.T) {
	t.Parallel()

	c := WeekendsOnly{}
	d := date(testYearStart, 1, 1)

	for i := 0; i < testDays; i++ {
		dow := d.Weekday()
		verify(t, c, d, 0, dow == time.Saturday || dow == time.Sunday, weekend)
		d = d.AddDate(0, 0, 1)
	}
}
