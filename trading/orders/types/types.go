// Package types enumerates types of an order.
package types

import (
	"bytes"
	"errors"
	"fmt"
)

// OrderType enumerates types of an order.
type OrderType int

const (
	// Market order is an order to buy (sell) an asset at the
	// ask (bid) price currently available in the marketplace.
	//
	// This order may increase the likelihood of a fill and the
	// speed of execution, but provides no price protection and
	// may fill at a price far lower (higher) than the current
	// bid (ask).
	Market OrderType = iota + 1

	// MarketIfTouched (market-if-touched) order is an order to
	// buy (sell) a stated amount of an asset as soon as the market
	// goes below (above) a preset price, at which point it becomes
	// a market order.
	//
	// This order is similar to a stop order, except that a
	// market-if-touched sell order is placed above the current
	// market price, and a stop sell order is placed below.
	MarketIfTouched

	// Limit order is an order to buy (sell) only at a specified
	// limit price or better, above (below) the limit price.
	//
	// A limit order may not get filled if the price never
	// reaches the specified limit price.
	Limit

	// Stop order is a buy (sell) stop order which becomes
	// a market order when the last traded price is
	// greater (less) -than-or-equal to the stop price.
	//
	// A buy (sell) stop price is always above (below) the
	// current market price.
	//
	// A stop order may not get filled if the price never
	// reaches the siscified stop price.
	Stop

	// StopLimit (stop-limit) order is a buy (sell) order
	// which becomes a limit order when the last traded price
	// is greater (less) -than-or-equal to the stop price.
	//
	// A buy (sell) stop price is always above (below)
	// the current market price.
	//
	// A stop-limit order may not get filled if the price
	// never reaches the specified stop price.
	StopLimit

	// TrailingStop (trailing-stop) order is a buy (sell) order
	// entered with a stop price at a fixed amount above (below)
	// the market price that creates a moving or trailing
	// activation price, hence the name.
	//
	// If the market price falls (rises), the stop loss price
	// rises (falls) by the increased amount, but if the stock
	// price falls, the stop loss price remains the same.
	//
	// The reverse is true for a buy trailing stop order.
	TrailingStop

	// MarketOnClose (market on close) order will execute as a
	// market order as close to the closing price as possible.
	MarketOnClose

	// MarketToLimit (market-to-limit) order is a market order
	// to execute at the current best price.
	//
	// If the entire order does not immediately execute at the
	// market price, the remainder of the order is re-submitted
	// as a limit order with the limit price set to the price at
	// which the market order portion of the order executed.
	MarketToLimit

	// LimitIfTouched (limit-if-touched) order is designed to
	// buy (or sell) a contract below (or above) the market,
	// at the limit price or better.
	LimitIfTouched

	// LimitOnClose (limit-on-close) order will fill at the
	// closing price if that price is at or better than the
	// submitted limit price.
	//
	// Otherwise, the order will be canceled.
	LimitOnClose
)

const (
	unknown         = "unknown"
	market          = "market"
	marketIfTouched = "marketIfTouched"
	limit           = "limit"
	stop            = "stop"
	stopLimit       = "stopLimit"
	trailingStop    = "trailingStop"
	marketOnClose   = "marketOnClose"
	marketToLimit   = "marketToLimit"
	limitIfTouched  = "limitIfTouched"
	limitOnClose    = "limitOnClose"
)

var errUnknownOrderType = errors.New("unknown order type")

//nolint:cyclop
// String implements the Stringer interface.
func (t OrderType) String() string {
	switch t {
	case Market:
		return market
	case MarketIfTouched:
		return marketIfTouched
	case Limit:
		return limit
	case Stop:
		return stop
	case StopLimit:
		return stopLimit
	case TrailingStop:
		return trailingStop
	case MarketOnClose:
		return marketOnClose
	case MarketToLimit:
		return marketToLimit
	case LimitIfTouched:
		return limitIfTouched
	case LimitOnClose:
		return limitOnClose
	default:
		return unknown
	}
}

// IsKnown determines if this order type is known.
func (t OrderType) IsKnown() bool {
	return t >= Market && t <= LimitOnClose
}

// MarshalJSON implements the Marshaler interface.
func (t OrderType) MarshalJSON() ([]byte, error) {
	s := t.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownOrderType)
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
func (t *OrderType) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	s := string(d)

	switch s {
	case market:
		*t = Market
	case marketIfTouched:
		*t = MarketIfTouched
	case limit:
		*t = Limit
	case stop:
		*t = Stop
	case stopLimit:
		*t = StopLimit
	case trailingStop:
		*t = TrailingStop
	case marketOnClose:
		*t = MarketOnClose
	case marketToLimit:
		*t = MarketToLimit
	case limitIfTouched:
		*t = LimitIfTouched
	case limitOnClose:
		*t = LimitOnClose
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", s, errUnknownOrderType)
	}

	return nil
}
