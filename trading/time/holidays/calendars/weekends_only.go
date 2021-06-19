package calendars

import "time"

// WeekendsOnly implements a weekends-only holiday schedule.
type WeekendsOnly struct{}

// IsHoliday implements Calendarer interface.
func (WeekendsOnly) IsHoliday(t time.Time) bool {
	return checkWeekend(t)
}
