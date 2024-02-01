//nolint:dupl
package chande

import (
	"bytes"
	"fmt"
)

// ChandeMomentumOscillatorOutput describes the outputs of the indicator.
type ChandeMomentumOscillatorOutput int

const (
	// The scalar value of the the Chande momentum oscillator.
	ChandeMomentumOscillatorValue ChandeMomentumOscillatorOutput = iota + 1
	chandeMomentumOscillatorLast
)

const (
	chandeMomentumOscillatorValue   = "value"
	chandeMomentumOscillatorUnknown = "unknown"
)

// String implements the Stringer interface.
func (o ChandeMomentumOscillatorOutput) String() string {
	switch o {
	case ChandeMomentumOscillatorValue:
		return chandeMomentumOscillatorValue
	default:
		return chandeMomentumOscillatorUnknown
	}
}

// IsKnown determines if this output is known.
func (o ChandeMomentumOscillatorOutput) IsKnown() bool {
	return o >= ChandeMomentumOscillatorValue && o < chandeMomentumOscillatorLast
}

// MarshalJSON implements the Marshaler interface.
func (o ChandeMomentumOscillatorOutput) MarshalJSON() ([]byte, error) {
	const (
		errFmt = "cannot marshal '%s': unknown Chande momentum oscillator output"
		extra  = 2   // Two bytes for quotes.
		dqc    = '"' // Double quote character.
	)

	s := o.String()
	if s == chandeMomentumOscillatorUnknown {
		return nil, fmt.Errorf(errFmt, s)
	}

	b := make([]byte, 0, len(s)+extra)
	b = append(b, dqc)
	b = append(b, s...)
	b = append(b, dqc)

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (o *ChandeMomentumOscillatorOutput) UnmarshalJSON(data []byte) error {
	const (
		errFmt = "cannot unmarshal '%s': unknown Chande momentum oscillator output"
		dqs    = "\"" // Double quote string.
	)

	d := bytes.Trim(data, dqs)
	s := string(d)

	switch s {
	case chandeMomentumOscillatorValue:
		*o = ChandeMomentumOscillatorValue
	default:
		return fmt.Errorf(errFmt, s)
	}

	return nil
}
