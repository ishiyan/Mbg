package timepieces

import (
	"fmt"
	"time"
)

// SynchronizedTimepiece is an externally synchronized step-time timepiece.
type SynchronizedTimepiece struct {
	// Now gets the current date and time.
	now             time.Time
	bod             time.Time
	mon             int
	day             int
	sec             int
	nsec            int
	secSessionStart int
	secSessionEnd   int
	durSessionStart time.Duration
	durSessionEnd   time.Duration
}

func (st SynchronizedTimepiece) Synchronize(t time.Time) {
	if !t.After(st.now) {
		return
	}

	y, m, day := t.Date()
	mon := int(m) + y*12 //nolint:gomnd

	sec := daySeconds(t)
	nsec := t.Nanosecond()

	if mon > st.mon || day > st.day { //nolint:nestif
		// delta := time.Sub(st.now)
		if st.sec <= st.secSessionEnd {
			if st.sec < st.secSessionStart {
				st.sessionStarted(st.bod.Add(st.durSessionStart))
			}

			st.sessionEnded(st.bod.Add(st.durSessionEnd))
		}
	} else if sec > st.sec || nsec > st.nsec {
		// Within the same day.
		if st.sec < st.secSessionStart {
			if sec >= st.secSessionStart {
				st.sessionStarted(st.bod.Add(st.durSessionStart))
				if sec > st.secSessionEnd {
					st.sessionEnded(st.bod.Add(st.durSessionEnd))
				}
			}
		} else if st.sec <= st.secSessionEnd && sec > st.secSessionEnd {
			st.sessionEnded(st.bod.Add(st.durSessionEnd))
		}

		st.now = t
		st.bod = beginOfDay(t)
		st.mon = mon
		st.day = day
		st.sec = sec
		st.nsec = nsec
	}
}

func (st SynchronizedTimepiece) sessionStarted(t time.Time) {
	fmt.Printf("session started: %v\n", t) //nolint:forbidigo
}

func (st SynchronizedTimepiece) sessionEnded(t time.Time) {
	fmt.Printf("session ended: %v\n", t) //nolint:forbidigo
}

func beginOfDay(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func daySeconds(t time.Time) int {
	h, m, s := t.Clock()

	return s + m*60 + h*3600
}
