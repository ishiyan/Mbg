// Package matchings enumerates algorithm used to match the offsetting order executions in a round-trip.
package matchings

import (
	"bytes"
	"errors"
	"fmt"
)

// Matching enumerates algorithms used to match the offsetting order executions in a round-trip.
type Matching int

const (
	// FirstInFirstOut means offsetting order executions will be matched in FIFO order.
	FirstInFirstOut Matching = iota + 1

	// LastInFirstOut means offsetting order executions will be matched in LIFO order.
	LastInFirstOut
)

const (
	unknown         = "unknown"
	firstInFirstOut = "firstInFirstOut"
	lastInFirstOut  = "lastInFirstOut"
)

var errUnknownMatching = errors.New("unknown round-trip matching")

// String implements the fmt.Stringer interface.
func (m Matching) String() string {
	switch m {
	case FirstInFirstOut:
		return firstInFirstOut
	case LastInFirstOut:
		return lastInFirstOut
	default:
		return unknown
	}
}

// IsKnown determines if this round-trip matching is known.
func (m Matching) IsKnown() bool {
	return m == FirstInFirstOut || m == LastInFirstOut
}

// MarshalJSON implements the Marshaler interface.
func (m Matching) MarshalJSON() ([]byte, error) {
	str := m.String()
	if str == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", str, errUnknownMatching)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(str)+extra)
	b = append(b, '"')
	b = append(b, str...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (m *Matching) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	str := string(d)

	switch str {
	case firstInFirstOut:
		*m = FirstInFirstOut
	case lastInFirstOut:
		*m = LastInFirstOut
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", str, errUnknownMatching)
	}

	return nil
}
