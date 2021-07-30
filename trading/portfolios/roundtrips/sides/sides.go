// Package sides enumerates sides of a round-trip.
package sides

import (
	"bytes"
	"errors"
	"fmt"
)

// Side is side of a round-trip.
type Side int

const (
	// Long is a long round-trip.
	Long Side = iota + 1

	// Short is a short round-trip.
	Short
)

const (
	unknown = "unknown"
	long    = "long"
	short   = "short"
)

var errUnknownSide = errors.New("unknown round-trip side")

// String implements the fmt.Stringer interface.
func (s Side) String() string {
	switch s {
	case Long:
		return long
	case Short:
		return short
	default:
		return unknown
	}
}

// IsKnown determines if this round-trip side is known.
func (s Side) IsKnown() bool {
	return s == Long || s == Short
}

// MarshalJSON implements the Marshaler interface.
func (s Side) MarshalJSON() ([]byte, error) {
	str := s.String()
	if str == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", str, errUnknownSide)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(str)+extra)
	b = append(b, '"')
	b = append(b, str...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (s *Side) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	str := string(d)

	switch str {
	case long:
		*s = Long
	case short:
		*s = Short
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", str, errUnknownSide)
	}

	return nil
}
