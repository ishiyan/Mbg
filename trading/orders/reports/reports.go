// Package reports enumerates an order report event types.
package reports

import (
	"bytes"
	"errors"
	"fmt"
)

// OrderReportType enumerates an order report event types.
type OrderReportType int

const (
	// PendingNew reports a transition to the "pending new" order status.
	PendingNew OrderReportType = iota + 1

	// New reports a transition to the "new" order status.
	New

	// Rejected reports a transition to the "rejected" order status.
	Rejected

	// PartiallyFilled reports a transition to the "partially filled" order status.
	PartiallyFilled

	// Filled reports a transition to the "filled" order status.
	Filled

	// Expired reports a transition to the "expired" order status.
	Expired

	// PendingReplace reports a transition to the "pending replace" order status.
	PendingReplace

	// Replaced reports that an order has been replaced.
	Replaced

	// ReplaceRejected reports that an order replacement has been rejected.
	ReplaceRejected

	// PendingCancel reports a transition to the "pending cancel" order status.
	PendingCancel

	// Canceled reports a transition to the "canceled" order status.
	Canceled

	// CancelRejected reports that an order cancellation has been rejected.
	CancelRejected

	// OrderStatus reports an order status.
	OrderStatus
)

const (
	unknown         = "unknown"
	pendingNew      = "pendingNew"
	new             = "new" //nolint
	rejected        = "rejected"
	partiallyFilled = "partiallyFilled"
	filled          = "filled"
	expired         = "expired"
	pendingReplace  = "pendingReplace"
	replaced        = "replaced"
	replaceRejected = "replaceRejected"
	pendingCancel   = "pendingCancel"
	canceled        = "canceled"
	cancelRejected  = "cancelRejected"
	orderStatus     = "orderStatus"
)

var errUnknownOrderReportType = errors.New("unknown order report type")

//nolint:cyclop
// String implements the Stringer interface.
func (t OrderReportType) String() string {
	switch t {
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
	case Replaced:
		return replaced
	case ReplaceRejected:
		return replaceRejected
	case PendingCancel:
		return pendingCancel
	case Canceled:
		return canceled
	case CancelRejected:
		return cancelRejected
	case OrderStatus:
		return orderStatus
	default:
		return unknown
	}
}

// IsKnown determines if this order report type is known.
func (t OrderReportType) IsKnown() bool {
	return t >= PendingNew && t <= OrderStatus
}

// MarshalJSON implements the Marshaler interface.
func (t OrderReportType) MarshalJSON() ([]byte, error) {
	s := t.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownOrderReportType)
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
func (t *OrderReportType) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	str := string(d)

	switch str {
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
	case replaced:
		*t = Replaced
	case replaceRejected:
		*t = ReplaceRejected
	case pendingCancel:
		*t = PendingCancel
	case canceled:
		*t = Canceled
	case cancelRejected:
		*t = CancelRejected
	case orderStatus:
		*t = OrderStatus
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", str, errUnknownOrderReportType)
	}

	return nil
}
