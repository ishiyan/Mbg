//nolint:dupl
package indicators

import (
	"bytes"
	"fmt"
)

// RateOfChangePercentOutput describes the outputs of the indicator.
type RateOfChangePercentOutput int

const (
	// The scalar value of the the rate of change.
	RateOfChangePercentValue RateOfChangePercentOutput = iota + 1
	rateOfChangePercentLast
)

const (
	rateOfChangePercentValue   = "value"
	rateOfChangePercentUnknown = "unknown"
)

// String implements the Stringer interface.
func (o RateOfChangePercentOutput) String() string {
	switch o {
	case RateOfChangePercentValue:
		return rateOfChangePercentValue
	default:
		return rateOfChangePercentUnknown
	}
}

// IsKnown determines if this output is known.
func (o RateOfChangePercentOutput) IsKnown() bool {
	return o >= RateOfChangePercentValue && o < rateOfChangePercentLast
}

// MarshalJSON implements the Marshaler interface.
func (o RateOfChangePercentOutput) MarshalJSON() ([]byte, error) {
	const (
		errFmt = "cannot marshal '%s': unknown rate of change percent output"
		extra  = 2   // Two bytes for quotes.
		dqc    = '"' // Double quote character.
	)

	s := o.String()
	if s == rateOfChangePercentUnknown {
		return nil, fmt.Errorf(errFmt, s)
	}

	b := make([]byte, 0, len(s)+extra)
	b = append(b, dqc)
	b = append(b, s...)
	b = append(b, dqc)

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (o *RateOfChangePercentOutput) UnmarshalJSON(data []byte) error {
	const (
		errFmt = "cannot unmarshal '%s': unknown rate of change percent output"
		dqs    = "\"" // Double quote string.
	)

	d := bytes.Trim(data, dqs)
	s := string(d)

	switch s {
	case rateOfChangePercentValue:
		*o = RateOfChangePercentValue
	default:
		return fmt.Errorf(errFmt, s)
	}

	return nil
}
