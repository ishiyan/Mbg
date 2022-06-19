// Package output encapsuletes info related to outputs of an indicator.
package output

import (
	"bytes"
	"errors"
	"fmt"
)

// Type identifies a type of an output of an indicator by enumerating all possible output types.
type Type int

const (
	// Holds a time stamp and a value.
	Scalar Type = iota + 1

	// Holds a time stamp and two values representing upper and lower lines of a band.
	Band

	// Holds a time stamp and an array of values representing a heat-map column.
	Heatmap
	last
)

const (
	unknown = "unknown"
	scalar  = "scalar"
	band    = "band"
	heatmap = "heatmap"
)

var errUnknownType = errors.New("unknown indicator output type")

// String implements the Stringer interface.
func (t Type) String() string {
	switch t {
	case Scalar:
		return scalar
	case Band:
		return band
	case Heatmap:
		return heatmap
	default:
		return unknown
	}
}

// IsKnown determines if this output type is known.
func (t Type) IsKnown() bool {
	return t >= Scalar && t < last
}

// MarshalJSON implements the Marshaler interface.
func (t Type) MarshalJSON() ([]byte, error) {
	s := t.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownType)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(s)+extra)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (t *Type) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	s := string(d)

	switch s {
	case scalar:
		*t = Scalar
	case band:
		*t = Band
	case heatmap:
		*t = Heatmap
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", s, errUnknownType)
	}

	return nil
}
