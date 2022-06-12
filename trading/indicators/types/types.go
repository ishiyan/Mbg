// Package types enumerates the types of indicators.
package types

import (
	"bytes"
	"errors"
	"fmt"
)

// IndicatorType Identifies an indicator.
type IndicatorType int

const (
	// SimpleMovingAverage identifies the Simple Moving Average (SMA) indicator.
	SimpleMovingAverage IndicatorType = iota + 1

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

var errUnknownIndicatorType = errors.New("unknown indicator type")

//nolint:exhaustive,cyclop
// String implements the Stringer interface.
func (i IndicatorType) String() string {
	switch i {
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

// IsKnown determines if this I=indicator type is known.
func (i IndicatorType) IsKnown() bool {
	return i >= SimpleMovingAverage && i < last
}

// MarshalJSON implements the Marshaler interface.
func (i IndicatorType) MarshalJSON() ([]byte, error) {
	s := i.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownIndicatorType)
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
func (i *IndicatorType) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	s := string(d)

	switch s {
	case simpleMovingAverage:
		*i = SimpleMovingAverage
	case exponentialMovingAverage:
		*i = ExponentialMovingAverage
	case bollingerBands:
		*i = BollingerBands
	case variance:
		*i = Variance
	case standardDeviation:
		*i = StandardDeviation
	case goertzelSpectrum:
		*i = GoertzelSpectrum
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", s, errUnknownIndicatorType)
	}

	return nil
}
