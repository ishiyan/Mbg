//nolint:dupl
package mulloy

import (
	"bytes"
	"fmt"
)

// TripleExponentialMovingAverageOutput describes the outputs of the indicator.
type TripleExponentialMovingAverageOutput int

const (
	// The scalar value of the the moving average.
	TripleExponentialMovingAverageValue TripleExponentialMovingAverageOutput = iota + 1
	tripleExponentialMovingAverageLast
)

const (
	tripleExponentialMovingAverageValue   = "value"
	tripleExponentialMovingAverageUnknown = "unknown"
)

// String implements the Stringer interface.
func (o TripleExponentialMovingAverageOutput) String() string {
	switch o {
	case TripleExponentialMovingAverageValue:
		return tripleExponentialMovingAverageValue
	default:
		return tripleExponentialMovingAverageUnknown
	}
}

// IsKnown determines if this output is known.
func (o TripleExponentialMovingAverageOutput) IsKnown() bool {
	return o >= TripleExponentialMovingAverageValue && o < tripleExponentialMovingAverageLast
}

// MarshalJSON implements the Marshaler interface.
func (o TripleExponentialMovingAverageOutput) MarshalJSON() ([]byte, error) {
	const (
		errFmt = "cannot marshal '%s': unknown triple exponential moving average output"
		extra  = 2   // Two bytes for quotes.
		dqc    = '"' // Double quote character.
	)

	s := o.String()
	if s == tripleExponentialMovingAverageUnknown {
		return nil, fmt.Errorf(errFmt, s)
	}

	b := make([]byte, 0, len(s)+extra)
	b = append(b, dqc)
	b = append(b, s...)
	b = append(b, dqc)

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (o *TripleExponentialMovingAverageOutput) UnmarshalJSON(data []byte) error {
	const (
		errFmt = "cannot unmarshal '%s': unknown triple exponential moving average output"
		dqs    = "\"" // Double quote string.
	)

	d := bytes.Trim(data, dqs)
	s := string(d)

	switch s {
	case tripleExponentialMovingAverageValue:
		*o = TripleExponentialMovingAverageValue
	default:
		return fmt.Errorf(errFmt, s)
	}

	return nil
}
