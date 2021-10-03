package data

//nolint:gci
import (
	"mbg/trading/data/entities"
	"sort"
	"time"
)

// ScalarTimeSeries contains a time series of values.
// The implementation is not thread-safe and should not be exposed directly.
type ScalarTimeSeries struct {
	scalar *entities.Scalar
	series []*entities.Scalar
}

var empty []entities.Scalar = []entities.Scalar{} //nolint:gochecknoglobals

// Current returns a current (the most recent) value of the time series or zero if time series is empty.
func (sts *ScalarTimeSeries) Current() float64 {
	if v := sts.scalar; v != nil {
		return v.Value
	}

	return 0
}

// At returns a value of the time series at or immediately before the given time.
// If the time is after the last time series sample, the most recent value is returned.
// Returns zero if the time is before the first sample or the time series is empty.
func (sts *ScalarTimeSeries) At(t time.Time) float64 {
	if h := sts.series; h != nil {
		l := len(h)
		i := sort.Search(l, func(i int) bool { return h[i].Time.After(t) })

		if i > 0 {
			return h[i-1].Value
		}
	}

	return 0
}

// History returns a copy of the time series or an empty slice if the time series is empty.
func (sts *ScalarTimeSeries) History() []entities.Scalar {
	if h := sts.series; h != nil {
		if l := len(h); l > 0 {
			v := make([]entities.Scalar, l)
			for i, s := range h {
				v[i] = *s
			}

			return v
		}
	}

	return empty
}

// Add appends a new sample to the time series if the new sample time is after the last time of the time series.
// If the new sample time is equal to the last time, the last time series value will be updated.
// Otherwise (if the new sample time is less than the last time), the sammple will be inserted.
func (sts *ScalarTimeSeries) Add(t time.Time, v float64) {
	switch {
	case sts.scalar == nil || sts.scalar.Time.Before(t): // The very first or a next sample.
		s := &entities.Scalar{Time: t, Value: v}
		sts.series = append(sts.series, s)
		sts.scalar = s
	case sts.scalar.Time.Equal(t): // The sample time is equal to the last time.
		sts.scalar.Value = v
	default: // The sample time is less than the last time.
		l := len(sts.series)
		i := sort.Search(l, func(i int) bool { return !sts.series[i].Time.Before(t) })

		if sts.series[i].Time.Equal(t) {
			sts.series[i].Value = v
		} else {
			s := &entities.Scalar{Time: t, Value: v}
			sts.series = append(sts.series[:i+1], sts.series[i:]...)
			sts.series[i] = s
		}
	}
}

// Accumulate appends a new sample to the time series adding the current value to the value of the new sample.
// If the new sample time is equal to the last time, the last time series value will be updated.
// Otherwise (if the new sample time is less than the last time), the sammple will be inserted
// and the values will be updated started from the newly inserted one.
func (sts *ScalarTimeSeries) Accumulate(t time.Time, v float64) {
	switch {
	case sts.scalar == nil: // The very first sample.
		s := &entities.Scalar{Time: t, Value: v}
		sts.series = append(sts.series, s)
		sts.scalar = s
	case sts.scalar.Time.Before(t): // Next sample.
		s := &entities.Scalar{Time: t, Value: v + sts.scalar.Value}
		sts.series = append(sts.series, s)
		sts.scalar = s
	case sts.scalar.Time.Equal(t): // Sample time is equal to the last time.
		sts.scalar.Value += v
	default: // Sample time is less than the last time.
		h := sts.series
		l := len(h)
		i := sort.Search(l, func(i int) bool { return !h[i].Time.Before(t) })

		if h[i].Time.Equal(t) {
			h[i].Value += v
			for i++; i < l; i++ {
				h[i].Value += v
			}
		} else {
			s := &entities.Scalar{Time: t, Value: v}
			if i > 0 {
				s.Value += h[i-1].Value
			}

			h = append(h[:i+1], h[i:]...)
			h[i] = s
			for i++; i <= l; i++ {
				h[i].Value += v
			}

			sts.series = h
		}
	}
}
