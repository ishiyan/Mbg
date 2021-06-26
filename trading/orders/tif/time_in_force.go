// Package tif enumerates time conditions an order is to be traded.
package tif

import (
	"bytes"
	"errors"
	"fmt"
)

// OrderTimeInForce enumerates time conditions an order is to be traded.
type OrderTimeInForce int

const (
	// Day requires an order to be executed within the trading day
	// on which it was entered.
	Day OrderTimeInForce = iota + 1

	// ImmediateOrCancel requires an order to be executed immediately
	// in whole or partially.
	//
	// Any portion not so executed is to be canceled.
	// Not to be confused with FillOrKill.
	ImmediateOrCancel

	// FillOrKill requires an order to be executed immediately in its entirety.
	//
	// If not so executed, the order is to be canceled.
	// Not to be confused with ImmediateOrCancel.
	FillOrKill

	// GoodTillCanceled requires an order to remain in effect until it is
	// either executed or canceled.
	//
	// Typically, GTC orders will be automatically be cancelled if a corporate
	// action on a security results in a stock split (forward or reverse),
	// exchange for shares, or distribution of shares.
	GoodTillCanceled

	// GoodTillDate requires an order, if not executed, to expire at the
	// specified date.
	GoodTillDate

	// AtOpen requires an order to be executed at the opening or not at all.
	//
	// All or part of any order not executed at the opening is treated as canceled.
	AtOpen

	// AtClose requires an order to be executed at the closing or not at all.
	//
	// All or part of any order not executed at the closing is treated as canceled.
	//
	// Indicated price is to be around the closing price, however,
	// not held to the closing price.
	AtClose
)

const (
	unknown           = "unknown"
	day               = "day"
	immediateOrCancel = "immediateOrCancel"
	fillOrKill        = "fillOrKill"
	goodTillCanceled  = "goodTillCanceled"
	goodTillDate      = "goodTillDate"
	atOpen            = "atOpen"
	atClose           = "atClose"
)

var errUnknownOrderTimeInForce = errors.New("unknown order time in force")

// String implements the Stringer interface.
func (t OrderTimeInForce) String() string {
	switch t {
	case Day:
		return day
	case ImmediateOrCancel:
		return immediateOrCancel
	case FillOrKill:
		return fillOrKill
	case GoodTillCanceled:
		return goodTillCanceled
	case GoodTillDate:
		return goodTillDate
	case AtOpen:
		return atOpen
	case AtClose:
		return atClose
	default:
		return unknown
	}
}

// IsKnown determines if this order status is known.
func (t OrderTimeInForce) IsKnown() bool {
	return t >= Day && t <= AtClose
}

// MarshalJSON implements the Marshaler interface.
func (t OrderTimeInForce) MarshalJSON() ([]byte, error) {
	s := t.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownOrderTimeInForce)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(s)+extra)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (t *OrderTimeInForce) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	str := string(d)

	switch str {
	case day:
		*t = Day
	case immediateOrCancel:
		*t = ImmediateOrCancel
	case fillOrKill:
		*t = FillOrKill
	case goodTillCanceled:
		*t = GoodTillCanceled
	case goodTillDate:
		*t = GoodTillDate
	case atOpen:
		*t = AtOpen
	case atClose:
		*t = AtClose
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", str, errUnknownOrderTimeInForce)
	}

	return nil
}
