// Package actions defines actions performed on account.
package actions

import (
	"bytes"
	"errors"
	"fmt"
)

// Action enumerates actions performed on account.
type Action int

const (
	// Credit is an action to deposit money to an account.
	Credit Action = iota + 1

	// Debit is an action to withdraw money from an account.
	Debit
)

const (
	unknown = "unknown"
	credit  = "credit"
	debit   = "debit"
)

var errUnknownAction = errors.New("unknown account action")

// String implements the Stringer interface.
func (a Action) String() string {
	switch a {
	case Credit:
		return credit
	case Debit:
		return debit
	default:
		return unknown
	}
}

// IsKnown determines if this account action is known.
func (a Action) IsKnown() bool {
	return a == Credit || a == Debit
}

// MarshalJSON implements the Marshaler interface.
func (a Action) MarshalJSON() ([]byte, error) {
	s := a.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownAction)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(s)+extra)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (a *Action) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	str := string(d)

	switch str {
	case credit:
		*a = Credit
	case debit:
		*a = Debit
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", str, errUnknownAction)
	}

	return nil
}
