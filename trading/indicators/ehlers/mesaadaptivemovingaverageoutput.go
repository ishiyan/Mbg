package ehlers

import (
	"bytes"
	"fmt"
)

// MesaAdaptiveMovingAverageOutput describes the outputs of the indicator.
type MesaAdaptiveMovingAverageOutput int

const (
	// The scalar value of the moving average.
	MesaAdaptiveMovingAverageValue MesaAdaptiveMovingAverageOutput = iota + 1
	mesaAdaptiveMovingAverageLast
)

const (
	mesaAdaptiveMovingAverageValue   = "value"
	mesaAdaptiveMovingAverageUnknown = "unknown"
)

// String implements the Stringer interface.
func (o MesaAdaptiveMovingAverageOutput) String() string {
	switch o {
	case MesaAdaptiveMovingAverageValue:
		return mesaAdaptiveMovingAverageValue
	default:
		return mesaAdaptiveMovingAverageUnknown
	}
}

// IsKnown determines if this output is known.
func (o MesaAdaptiveMovingAverageOutput) IsKnown() bool {
	return o >= MesaAdaptiveMovingAverageValue && o < mesaAdaptiveMovingAverageLast
}

// MarshalJSON implements the Marshaler interface.
func (o MesaAdaptiveMovingAverageOutput) MarshalJSON() ([]byte, error) {
	const (
		errFmt = "cannot marshal '%s': unknown mesa adaptive moving average output"
		extra  = 2   // Two bytes for quotes.
		dqc    = '"' // Double quote character.
	)

	s := o.String()
	if s == mesaAdaptiveMovingAverageUnknown {
		return nil, fmt.Errorf(errFmt, s)
	}

	b := make([]byte, 0, len(s)+extra)
	b = append(b, dqc)
	b = append(b, s...)
	b = append(b, dqc)

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (o *MesaAdaptiveMovingAverageOutput) UnmarshalJSON(data []byte) error {
	const (
		errFmt = "cannot unmarshal '%s': unknown mesa adaptive moving average output"
		dqs    = "\"" // Double quote string.
	)

	d := bytes.Trim(data, dqs)
	s := string(d)

	switch s {
	case mesaAdaptiveMovingAverageValue:
		*o = MesaAdaptiveMovingAverageValue
	default:
		return fmt.Errorf(errFmt, s)
	}

	return nil
}
