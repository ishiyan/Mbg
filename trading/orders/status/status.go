// Package status enumerates the states an order runs through during its lifetime.
package status

import (
	"bytes"
	"errors"
	"fmt"
)

// OrderStatus enumerates the states an order runs through during its lifetime.
type OrderStatus int

const (
	// Accepted indicates the order has been received by the broker and is being evaluated.
	//
	// The order will proceed to the PendingNew status.
	Accepted OrderStatus = iota + 1

	// PendingNew indicates the order has been accepted by the broker but not yet acknowledged
	// for execution.
	//
	// The order will proceed to the either New or the Rejected status.
	PendingNew

	// New indicates the order has been acknowledged by the broker and becomes the
	// outstanding order with no executions.
	//
	// The order can proceed to the Filled, the PartiallyFilled, the Expired,
	// the PendingCancel, the PendingReplace, or to the Rejected status.
	New

	// Rejected indicates the order has been rejected by the broker. No executions were done.
	//
	// This is a terminal state of an order, no further changes are allowed.
	Rejected

	// PartiallyFilled indicates the order has been partially filled and has remaining quantity.
	//
	// The order can proceed to the Filled, the PendingCancel, or to the
	// PendingReplace status.
	PartiallyFilled

	// Filled indicates the order has been completely filled.
	//
	// This is a terminal state of an order, no further changes are allowed.
	Filled

	// Expired indicates the order (with or without executions) has been canceled
	// in broker’s system due to time in force instructions.
	//
	// The only exceptions are FillOrKill and ImmediateOrCancel
	// orders that have Canceled as terminal order state.
	//
	// This is a terminal state of an order, no further changes are allowed.
	Expired

	// PendingReplace indicates a replace request has been sent to the broker, but the broker
	// hasn't replaced the order yet.
	//
	// The order will proceed back to the previous status.
	PendingReplace

	// PendingCancel indicates a cancel request has been sent to the broker, but
	// the broker hasn't canceled the order yet.
	//
	// The order will proceed to the either Canceled
	// or back to the previous status.
	PendingCancel

	// Canceled indicates the order (with or without executions)
	// has been canceled by the broker.
	//
	// The order may still be partially filled.
	// This is a terminal state of an order, no further changes are allowed.
	Canceled
)

const (
	unknown         = "unknown"
	accepted        = "accepted"
	pendingNew      = "pendingNew"
	new             = "new" //nolint
	rejected        = "rejected"
	partiallyFilled = "partiallyFilled"
	filled          = "filled"
	expired         = "expired"
	pendingReplace  = "pendingReplace"
	pendingCancel   = "pendingCancel"
	canceled        = "canceled"
)

var errUnknownOrderStatus = errors.New("unknown order status")

//nolint:cyclop
// String implements the Stringer interface.
func (t OrderStatus) String() string {
	switch t {
	case Accepted:
		return accepted
	case PendingNew:
		return pendingNew
	case New:
		return new
	case Rejected:
		return rejected
	case PartiallyFilled:
		return partiallyFilled
	case Filled:
		return filled
	case Expired:
		return expired
	case PendingReplace:
		return pendingReplace
	case PendingCancel:
		return pendingCancel
	case Canceled:
		return canceled
	default:
		return unknown
	}
}

// IsKnown determines if this order status is known.
func (t OrderStatus) IsKnown() bool {
	return t >= Accepted && t <= Canceled
}

// MarshalJSON implements the Marshaler interface.
func (t OrderStatus) MarshalJSON() ([]byte, error) {
	s := t.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownOrderStatus)
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
func (t *OrderStatus) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	str := string(d)

	switch str {
	case accepted:
		*t = Accepted
	case pendingNew:
		*t = PendingNew
	case new:
		*t = New
	case rejected:
		*t = Rejected
	case partiallyFilled:
		*t = PartiallyFilled
	case filled:
		*t = Filled
	case expired:
		*t = Expired
	case pendingReplace:
		*t = PendingReplace
	case pendingCancel:
		*t = PendingCancel
	case canceled:
		*t = Canceled
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", str, errUnknownOrderStatus)
	}

	return nil
}
