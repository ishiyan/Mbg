package statistics

import (
	"bytes"
	"errors"
	"fmt"
)

// StandardDeviationOutput describes the outputs of the indicator.
type StandardDeviationOutput int

const (
	// The scalar value of the the standard deviation.
	StandardDeviationValue StandardDeviationOutput = iota + 1
	standardDeviationLast
)

const (
	standardDeviationValue = "value"
)

var errUnknownStandardDeviationOutput = errors.New("unknown standard deviation output")

// String implements the Stringer interface.
func (o StandardDeviationOutput) String() string {
	switch o {
	case StandardDeviationValue:
		return standardDeviationValue
	default:
		return unknown
	}
}

// IsKnown determines if this output is known.
func (o StandardDeviationOutput) IsKnown() bool {
	return o >= StandardDeviationValue && o < standardDeviationLast
}

// MarshalJSON implements the Marshaler interface.
func (o StandardDeviationOutput) MarshalJSON() ([]byte, error) {
	s := o.String()
	if s == unknown {
		return nil, fmt.Errorf(marshalErrFmt, s, errUnknownStandardDeviationOutput)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(s)+extra)
	b = append(b, dqc)
	b = append(b, s...)
	b = append(b, dqc)

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (o *StandardDeviationOutput) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, dqs)
	s := string(d)

	switch s {
	case standardDeviationValue:
		*o = StandardDeviationValue
	default:
		return fmt.Errorf(unmarshalErrFmt, s, errUnknownStandardDeviationOutput)
	}

	return nil
}
