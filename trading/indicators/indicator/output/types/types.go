// Package types enumerates the output entitity types of indicators.
package types

import (
	"bytes"
	"errors"
	"fmt"
)

// OutputType identifies an output entity type of an indicator.
type OutputType int

const (
	// Holds a time stamp and a value.
	Scalar OutputType = iota + 1

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

var errUnknownOutputType = errors.New("unknown indicator output type")

//nolint:exhaustive,cyclop
// String implements the Stringer interface.
func (i OutputType) String() string {
	switch i {
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
func (i OutputType) IsKnown() bool {
	return i >= Scalar && i < last
}

// MarshalJSON implements the Marshaler interface.
func (i OutputType) MarshalJSON() ([]byte, error) {
	s := i.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownOutputType)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(s)+extra)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')

	return b, nil
}

//nolint:cyclop
// UnmarshalJSON implements the Unmarshaler interface.
func (i *OutputType) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	s := string(d)

	switch s {
	case scalar:
		*i = Scalar
	case band:
		*i = Band
	case heatmap:
		*i = Heatmap
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", s, errUnknownOutputType)
	}

	return nil
}
