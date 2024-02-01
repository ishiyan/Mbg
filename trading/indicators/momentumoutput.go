//nolint:dupl
package indicators

import (
	"bytes"
	"fmt"
)

// MomentumOutput describes the outputs of the indicator.
type MomentumOutput int

const (
	// The scalar value of the the momentum.
	MomentumValue MomentumOutput = iota + 1
	momentumLast
)

const (
	momentumValue   = "value"
	momentumUnknown = "unknown"
)

// String implements the Stringer interface.
func (o MomentumOutput) String() string {
	switch o {
	case MomentumValue:
		return momentumValue
	default:
		return momentumUnknown
	}
}

// IsKnown determines if this output is known.
func (o MomentumOutput) IsKnown() bool {
	return o >= MomentumValue && o < momentumLast
}

// MarshalJSON implements the Marshaler interface.
func (o MomentumOutput) MarshalJSON() ([]byte, error) {
	const (
		errFmt = "cannot marshal '%s': unknown momentum output"
		extra  = 2   // Two bytes for quotes.
		dqc    = '"' // Double quote character.
	)

	s := o.String()
	if s == momentumUnknown {
		return nil, fmt.Errorf(errFmt, s)
	}

	b := make([]byte, 0, len(s)+extra)
	b = append(b, dqc)
	b = append(b, s...)
	b = append(b, dqc)

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (o *MomentumOutput) UnmarshalJSON(data []byte) error {
	const (
		errFmt = "cannot unmarshal '%s': unknown momentum output"
		dqs    = "\"" // Double quote string.
	)

	d := bytes.Trim(data, dqs)
	s := string(d)

	switch s {
	case momentumValue:
		*o = MomentumValue
	default:
		return fmt.Errorf(errFmt, s)
	}

	return nil
}
