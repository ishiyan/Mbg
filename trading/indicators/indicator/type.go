package indicator

import (
	"bytes"
	"errors"
	"fmt"
)

// Type Identifies an indicator by enumerating all implemented indicators.
type Type int

const (
	// SimpleMovingAverage identifies the Simple Moving Average (SMA) indicator.
	SimpleMovingAverage Type = iota + 1

	// ExponentialMovingAverage identifies the Exponential Moving Average (EMA) indicator.
	ExponentialMovingAverage

	// BollingerBands identifies the Bollinger Bands (BB) indicator.
	BollingerBands

	// Variance identifies the Variance (VAR) indicator.
	Variance

	// StandardDeviation identifies the Standard Deviation (STDEV) indicator.
	StandardDeviation

	// GoertzelSpectrum identifies the Goertzel power spectrum (GOERTZEL) indicator.
	GoertzelSpectrum
	last
)

const (
	unknown                  = "unknown"
	simpleMovingAverage      = "simpleMovingAverage"
	exponentialMovingAverage = "exponentialMovingAverage"
	bollingerBands           = "bollingerBands"
	variance                 = "variance"
	standardDeviation        = "standardDeviation"
	goertzelSpectrum         = "goertzelSpectrum"
)

var errUnknownType = errors.New("unknown indicator type")

//nolint:exhaustive,cyclop
// String implements the Stringer interface.
func (t Type) String() string {
	switch t {
	case SimpleMovingAverage:
		return simpleMovingAverage
	case ExponentialMovingAverage:
		return exponentialMovingAverage
	case BollingerBands:
		return bollingerBands
	case Variance:
		return variance
	case StandardDeviation:
		return standardDeviation
	case GoertzelSpectrum:
		return goertzelSpectrum
	default:
		return unknown
	}
}

// IsKnown determines if this indicator type is known.
func (t Type) IsKnown() bool {
	return t >= SimpleMovingAverage && t < last
}

// MarshalJSON implements the Marshaler interface.
func (t Type) MarshalJSON() ([]byte, error) {
	s := t.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownType)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(s)+extra)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')

	return b, nil
}

//nolint:cyclop
// UnmarshalJSON implements the Unmarshaler interface.
func (t *Type) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	s := string(d)

	switch s {
	case simpleMovingAverage:
		*t = SimpleMovingAverage
	case exponentialMovingAverage:
		*t = ExponentialMovingAverage
	case bollingerBands:
		*t = BollingerBands
	case variance:
		*t = Variance
	case standardDeviation:
		*t = StandardDeviation
	case goertzelSpectrum:
		*t = GoertzelSpectrum
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", s, errUnknownType)
	}

	return nil
}
