package statistics

import (
	"bytes"
	"errors"
	"fmt"
)

// VarianceOutput describes the outputs of the indicator.
type VarianceOutput int

const (
	// The scalar value of the the variance.
	VarianceValue VarianceOutput = iota + 1
	varianceLast
)

const (
	varianceValue = "value"
)

var errUnknownVarianceOutput = errors.New("unknown variance output")

// String implements the Stringer interface.
func (o VarianceOutput) String() string {
	switch o {
	case VarianceValue:
		return varianceValue
	default:
		return unknown
	}
}

// IsKnown determines if this output is known.
func (o VarianceOutput) IsKnown() bool {
	return o >= VarianceValue && o < varianceLast
}

// MarshalJSON implements the Marshaler interface.
func (o VarianceOutput) MarshalJSON() ([]byte, error) {
	s := o.String()
	if s == unknown {
		return nil, fmt.Errorf(marshalErrFmt, s, errUnknownVarianceOutput)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(s)+extra)
	b = append(b, dqc)
	b = append(b, s...)
	b = append(b, dqc)

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (o *VarianceOutput) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, dqs)
	s := string(d)

	switch s {
	case varianceValue:
		*o = VarianceValue
	default:
		return fmt.Errorf(unmarshalErrFmt, s, errUnknownVarianceOutput)
	}

	return nil
}
