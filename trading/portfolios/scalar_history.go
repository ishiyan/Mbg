package portfolios

//nolint:gci
import (
	"mbg/trading/data/entities"
	"time"
)

// scalarHistory contains a time series of values.
// The implementation is not thread-safe and should not be exposed directly.
type scalarHistory struct {
	scalar  *entities.Scalar
	history []*entities.Scalar
}

// Current returns a current (the latest) value of the time series or zero if not initialized.
func (sh *scalarHistory) Current() float64 {
	if v := sh.scalar; v != nil {
		return v.Value
	}

	return 0
}

// History returns a copy of the time series or an empty slice if not initialized.
func (sh *scalarHistory) History() []entities.Scalar {
	if sh.history == nil || len(sh.history) == 0 {
		return []entities.Scalar{}
	}

	v := make([]entities.Scalar, len(sh.history))
	for i, s := range sh.history {
		v[i] = *s
	}

	return v
}

// add adds a new sample to the time series
// if the new sample time is later the last time of the time series.
// Otherwise (if the new sample time is less or equal to the last time),
// the last time series value will be updated.
func (sh *scalarHistory) add(time time.Time, value float64) {
	if sh.scalar == nil || sh.scalar.Time.Before(time) { // The very first or a next sample.
		s := &entities.Scalar{Time: time, Value: value}
		sh.history = append(sh.history, s)
		sh.scalar = s
	} else { // Sample time is less or equal to the last time.
		sh.scalar.Value = value
	}
}

// accumulate adds a new sample to the time series
// adding the current value to the value of the new sample.
// If the new sample time is less or equal to the last time,
// the last time series value will be updated instead.
func (sh *scalarHistory) accumulate(time time.Time, value float64) {
	switch {
	case sh.scalar == nil: // The very first sample.
		s := &entities.Scalar{Time: time, Value: value}
		sh.history = append(sh.history, s)
		sh.scalar = s
	case sh.scalar.Time.Before(time): // Next sample.
		s := &entities.Scalar{Time: time, Value: value + sh.scalar.Value}
		sh.history = append(sh.history, s)
		sh.scalar = s
	default: // Sample time is less or equal to the last time.
		sh.scalar.Value += value
	}
}
