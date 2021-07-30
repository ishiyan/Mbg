// Package groupings enumerates algorithms used to group order executions into round-trips.
package groupings

import (
	"bytes"
	"errors"
	"fmt"
)

// Grouping enumerates algorithms used to group order executions into round-trips.
type Grouping int

const (
	// FillToFill is a round-trip defined by
	// (1) an order execution that establishes or increases a position and
	// (2) an offsetting execution that reduces the position size.
	FillToFill Grouping = iota + 1

	// FlatToFlat is a round-trip defined by a sequence of order executions,
	// from a flat position to a non-zero position which may increase or decrease
	// in quantity, and back to a flat position.
	FlatToFlat

	// FlatToReduced is a round-trip defined by a sequence of order executions,
	// from a flat position to a non-zero position and an offsetting execution
	// that reduces the position size.
	FlatToReduced
)

const (
	unknown       = "unknown"
	fillToFill    = "fillToFill"
	flatToFlat    = "flatToFlat"
	flatToReduced = "flatToReduced"
)

var errUnknownGrouping = errors.New("unknown round trip grouping")

// String implements the fmt.Stringer interface.
func (g Grouping) String() string {
	switch g {
	case FillToFill:
		return fillToFill
	case FlatToFlat:
		return flatToFlat
	case FlatToReduced:
		return flatToReduced
	default:
		return unknown
	}
}

// IsKnown determines if this round-trip grouping is known.
func (g Grouping) IsKnown() bool {
	return g >= FillToFill && g <= FlatToReduced
}

// MarshalJSON implements the Marshaler interface.
func (g Grouping) MarshalJSON() ([]byte, error) {
	s := g.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownGrouping)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(s)+extra)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (g *Grouping) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	str := string(d)

	switch str {
	case fillToFill:
		*g = FillToFill
	case flatToFlat:
		*g = FlatToFlat
	case flatToReduced:
		*g = FlatToReduced
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", str, errUnknownGrouping)
	}

	return nil
}
