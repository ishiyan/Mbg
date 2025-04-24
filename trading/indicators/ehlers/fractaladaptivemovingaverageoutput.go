package ehlers

import (
	"bytes"
	"fmt"
)

// FractalAdaptiveMovingAverageOutput describes the outputs of the indicator.
type FractalAdaptiveMovingAverageOutput int

const (
	// The scalar value of the moving average.
	FractalAdaptiveMovingAverageValue FractalAdaptiveMovingAverageOutput = iota + 1
	FractalAdaptiveMovingAverageValueFdim
	fractalAdaptiveMovingAverageLast
)

const (
	fractalAdaptiveMovingAverageValue     = "value"
	fractalAdaptiveMovingAverageValueFdim = "fdim"
	fractalAdaptiveMovingAverageUnknown   = "unknown"
)

// String implements the Stringer interface.
func (o FractalAdaptiveMovingAverageOutput) String() string {
	switch o {
	case FractalAdaptiveMovingAverageValue:
		return fractalAdaptiveMovingAverageValue
	case FractalAdaptiveMovingAverageValueFdim:
		return fractalAdaptiveMovingAverageValueFdim
	default:
		return fractalAdaptiveMovingAverageUnknown
	}
}

// IsKnown determines if this output is known.
func (o FractalAdaptiveMovingAverageOutput) IsKnown() bool {
	return o >= FractalAdaptiveMovingAverageValue && o < fractalAdaptiveMovingAverageLast
}

// MarshalJSON implements the Marshaler interface.
func (o FractalAdaptiveMovingAverageOutput) MarshalJSON() ([]byte, error) {
	const (
		errFmt = "cannot marshal '%s': unknown fractal adaptive moving average output"
		extra  = 2   // Two bytes for quotes.
		dqc    = '"' // Double quote character.
	)

	s := o.String()
	if s == fractalAdaptiveMovingAverageUnknown {
		return nil, fmt.Errorf(errFmt, s)
	}

	b := make([]byte, 0, len(s)+extra)
	b = append(b, dqc)
	b = append(b, s...)
	b = append(b, dqc)

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (o *FractalAdaptiveMovingAverageOutput) UnmarshalJSON(data []byte) error {
	const (
		errFmt = "cannot unmarshal '%s': unknown fractal adaptive moving average output"
		dqs    = "\"" // Double quote string.
	)

	d := bytes.Trim(data, dqs)
	s := string(d)

	switch s {
	case fractalAdaptiveMovingAverageValue:
		*o = FractalAdaptiveMovingAverageValue
	case fractalAdaptiveMovingAverageValueFdim:
		*o = FractalAdaptiveMovingAverageValueFdim
	default:
		return fmt.Errorf(errFmt, s)
	}

	return nil
}
