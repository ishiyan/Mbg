package kaufman

import (
	"bytes"
	"fmt"
)

// AdaptiveMovingAverageOutput describes the outputs of the indicator.
type AdaptiveMovingAverageOutput int

const (
	// The scalar value of the moving average.
	AdaptiveMovingAverageValue AdaptiveMovingAverageOutput = iota + 1
	AdaptiveMovingAverageValueEr
	adaptiveMovingAverageLast
)

const (
	adaptiveMovingAverageValue   = "value"
	adaptiveMovingAverageValueEr = "er"
	adaptiveMovingAverageUnknown = "unknown"
)

// String implements the Stringer interface.
func (o AdaptiveMovingAverageOutput) String() string {
	switch o {
	case AdaptiveMovingAverageValue:
		return adaptiveMovingAverageValue
	case AdaptiveMovingAverageValueEr:
		return adaptiveMovingAverageValueEr
	default:
		return adaptiveMovingAverageUnknown
	}
}

// IsKnown determines if this output is known.
func (o AdaptiveMovingAverageOutput) IsKnown() bool {
	return o >= AdaptiveMovingAverageValue && o < adaptiveMovingAverageLast
}

// MarshalJSON implements the Marshaler interface.
func (o AdaptiveMovingAverageOutput) MarshalJSON() ([]byte, error) {
	const (
		errFmt = "cannot marshal '%s': unknown kaufman adaptive moving average output"
		extra  = 2   // Two bytes for quotes.
		dqc    = '"' // Double quote character.
	)

	s := o.String()
	if s == adaptiveMovingAverageUnknown {
		return nil, fmt.Errorf(errFmt, s)
	}

	b := make([]byte, 0, len(s)+extra)
	b = append(b, dqc)
	b = append(b, s...)
	b = append(b, dqc)

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (o *AdaptiveMovingAverageOutput) UnmarshalJSON(data []byte) error {
	const (
		errFmt = "cannot unmarshal '%s': unknown kaufman adaptive moving average output"
		dqs    = "\"" // Double quote string.
	)

	d := bytes.Trim(data, dqs)
	s := string(d)

	switch s {
	case adaptiveMovingAverageValue:
		*o = AdaptiveMovingAverageValue
	case adaptiveMovingAverageValueEr:
		*o = AdaptiveMovingAverageValueEr
	default:
		return fmt.Errorf(errFmt, s)
	}

	return nil
}
