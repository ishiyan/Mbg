// Package status enumerates the states of an instrument.
package status

import (
	"bytes"
	"errors"
	"fmt"
)

// InstrumentStatus enumerates the states of an instrument.
type InstrumentStatus int

const (
	// Active indicates the trading is possible.
	Active InstrumentStatus = iota + 1

	// ActiveClosingOrdersOnly indicates an instrument is active but only orders closing positions are allowed.
	ActiveClosingOrdersOnly

	// Inactive indicates an instrument has previously been active and is now no longer traded but has not expired yet.
	// It may become active again.
	Inactive

	// Suspended indicates an instrument has been temporarily disabled for trading.
	Suspended

	// PendingExpiry indicates an instrument is currently still active but will expire after the current business day.
	//
	// For example, a contract that expires intraday (e.g. at noon time) and is no longer traded
	// but will still show up in the current day's order book with related statistics.
	PendingExpiry

	// Expired indicates an instrument has been expired due to reaching maturity
	// or based on contract definitions or exchange rules.
	Expired

	// PendingDeletion indicates an instrument is awaiting deletion from security reference data.
	PendingDeletion

	// Delisted indicates an instrument has been removed from securities reference data.
	//
	// A delisted instrument would not trade on the exchange but it may still be traded over-the-counter.
	//
	// Delisting rules vary from exchange to exchange, which may include non-compliance of
	// capitalization, revenue, consecutive minimum closing price.
	// The instrument may become listed again once the instrument is back in compliance.
	Delisted

	// KnockedOut indicates an instrument has breached a predefined price threshold.
	KnockedOut

	// KnockOutRevoked indicates an instrument reinstated, i.e. threshold has not been breached.
	KnockOutRevoked
	last
)

const (
	unknown                 = "unknown"
	active                  = "active"
	activeClosingOrdersOnly = "activeClosingOrdersOnly"
	inactive                = "inactive"
	suspended               = "suspended"
	pendingExpiry           = "pendingExpiry"
	expired                 = "expired"
	pendingDeletion         = "pendingDeletion"
	delisted                = "delisted"
	knockedOut              = "knockedOut"
	knockOutRevoked         = "knockOutRevoked"
)

var errUnknownInstrumentStatus = errors.New("unknown instrument status")

// String implements the Stringer interface.
//
//nolint:exhaustive,cyclop
func (i InstrumentStatus) String() string {
	switch i {
	case Active:
		return active
	case ActiveClosingOrdersOnly:
		return activeClosingOrdersOnly
	case Inactive:
		return inactive
	case Suspended:
		return suspended
	case PendingExpiry:
		return pendingExpiry
	case Expired:
		return expired
	case PendingDeletion:
		return pendingDeletion
	case Delisted:
		return delisted
	case KnockedOut:
		return knockedOut
	case KnockOutRevoked:
		return knockOutRevoked
	default:
		return unknown
	}
}

// IsKnown determines if this instrument status is known.
func (i InstrumentStatus) IsKnown() bool {
	return i >= Active && i < last
}

// MarshalJSON implements the Marshaler interface.
func (i InstrumentStatus) MarshalJSON() ([]byte, error) {
	s := i.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownInstrumentStatus)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(s)+extra)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
//
//nolint:cyclop
func (i *InstrumentStatus) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	s := string(d)

	switch s {
	case active:
		*i = Active
	case activeClosingOrdersOnly:
		*i = ActiveClosingOrdersOnly
	case inactive:
		*i = Inactive
	case suspended:
		*i = Suspended
	case pendingExpiry:
		*i = PendingExpiry
	case expired:
		*i = Expired
	case pendingDeletion:
		*i = PendingDeletion
	case delisted:
		*i = Delisted
	case knockedOut:
		*i = KnockedOut
	case knockOutRevoked:
		*i = KnockOutRevoked
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", s, errUnknownInstrumentStatus)
	}

	return nil
}
