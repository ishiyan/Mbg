//nolint:dupl
package indicators

import (
	"bytes"
	"fmt"
)

// RateOfChangeOutput describes the outputs of the indicator.
type RateOfChangeOutput int

const (
	// The scalar value of the the rate of change.
	RateOfChangeValue RateOfChangeOutput = iota + 1
	rateOfChangeLast
)

const (
	rateOfChangeValue   = "value"
	rateOfChangeUnknown = "unknown"
)

// String implements the Stringer interface.
func (o RateOfChangeOutput) String() string {
	switch o {
	case RateOfChangeValue:
		return rateOfChangeValue
	default:
		return rateOfChangeUnknown
	}
}

// IsKnown determines if this output is known.
func (o RateOfChangeOutput) IsKnown() bool {
	return o >= RateOfChangeValue && o < rateOfChangeLast
}

// MarshalJSON implements the Marshaler interface.
func (o RateOfChangeOutput) MarshalJSON() ([]byte, error) {
	const (
		errFmt = "cannot marshal '%s': unknown rate of change output"
		extra  = 2   // Two bytes for quotes.
		dqc    = '"' // Double quote character.
	)

	s := o.String()
	if s == rateOfChangeUnknown {
		return nil, fmt.Errorf(errFmt, s)
	}

	b := make([]byte, 0, len(s)+extra)
	b = append(b, dqc)
	b = append(b, s...)
	b = append(b, dqc)

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (o *RateOfChangeOutput) UnmarshalJSON(data []byte) error {
	const (
		errFmt = "cannot unmarshal '%s': unknown rate of change output"
		dqs    = "\"" // Double quote string.
	)

	d := bytes.Trim(data, dqs)
	s := string(d)

	switch s {
	case rateOfChangeValue:
		*o = RateOfChangeValue
	default:
		return fmt.Errorf(errFmt, s)
	}

	return nil
}
