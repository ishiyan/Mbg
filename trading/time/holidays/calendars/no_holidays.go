package calendars

import "time"

// NoHolidays implements a physical days (no holidays) schedule.
type NoHolidays struct{}

// IsHoliday implements Calendarer interface.
func (NoHolidays) IsHoliday(t time.Time) bool {
	return false
}
