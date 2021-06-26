// Package sides enumerates sides of an order.
package sides

import (
	"bytes"
	"errors"
	"fmt"
)

// OrderSide enumerates sides of an order.
type OrderSide int

const (
	// Buy order side refers to the buying of a security.
	Buy OrderSide = iota + 1

	// Sell order side refers to the selling of a security.
	Sell

	// BuyMinus (buy "minus") is a buy order provided that the price is not higher than
	// the last sale if the last sale was a “minus” or “zero minus” tick and not higher
	// than the last sale minus the minimum fractional change in the security if the last
	// sale was a “plus” or “zero plus” tick.
	//
	// “Minus tick” is a trade executed at a lower price than the preceding trade.
	//
	// “Zero minus tick” is a trade executed at the same price as the preceding
	// trade, but at a lower price than the last trade of a different price.
	//
	// For example, if a succession of trades occurs at $10.25, $10.00 and $10.00 again,
	// the second trade is a "minus tick" trade and the latter trade is a “zero minus tick”
	// or “zero downtick” trade.
	BuyMinus

	// SellPlus (sell “plus”) is a sell order provided that the price is not lower than
	// the last sale if the last sale was a “plus” or “zero plus” tick and not lower
	// than the last sale minus the minimum fractional change in the security if the last
	// sale was a “minus” or “zero minus” tick.
	//
	// “Plus tick” is a trade that is executed at a higher price than the preceding trade.
	//
	// “Zero plus tick” is a trade that is executed at the same price as the preceding
	// trade but at a higher price than the last trade of a different price.
	//
	// For example, if a succession of trades occurs at $10.00, $10.25 and $10.25 again,
	// the second trade is a "plus tick" trade and the latter trade is a “zero plus tick”
	// or “zero uptick” trade.
	SellPlus

	// SellShort is an order to sell a security that the seller does not own.
	//
	// Since 1938 there was an “uptick rule” established by the U.S. Securities
	// and Exchange Commission (SEC). The rule stated that securities could be
	// shorted only on an “uptick“ or a “zero plus“ tick, not on a “downtick“
	// or a “zero minus“ tick. This rule was lifted in 2007, allowing short
	// sales to occur on any price tick in the market, whether up or down.
	//
	// However, in 2010 the SEC adopted the alternative uptick rule applies to
	// all securities, which is triggered when the price of a security has dropped
	// by 10% or more from the previous day's close. When the rule is in effect,
	// short selling is permitted if the price is above the current best bid and
	// stays in effect for the rest of the day and the following trading session.
	//
	// “Plus tick” is a trade that is executed at a higher price than the preceding trade.
	//
	// “Zero plus tick” is a trade that is executed at the same price as the preceding
	// trade but at a higher price than the last trade of a different price.
	//
	// For example, if a succession of trades occurs at $10, $10.25 and $10.25 again,
	// the second trade is a "plus tick" trade and the latter trade is a “zero plus tick”
	// or “zero uptick” trade.
	//
	// “Minus tick” is a trade executed at a lower price than the preceding trade.
	//
	// “Zero minus tick” is a trade executed at the same price as the preceding
	// trade, but at a lower price than the last trade of a different price.
	//
	// For example, if a succession of trades occurs at $10.25, $10.00 and $10.00 again,
	// the second trade is a "minus tick" trade and the latter trade is a “zero minus tick”
	// or “zero downtick” trade.
	SellShort

	// SellShortExempt refers to a special trading situation where a short sale is
	// allowed on a minustick.
	SellShortExempt
)

const (
	unknown         = "unknown"
	buy             = "buy"
	sell            = "sell"
	buyMinus        = "buyMinus"
	sellPlus        = "sellPlus"
	sellShort       = "sellShort"
	sellShortExempt = "sellShortExempt"
)

var errUnknownOrderSide = errors.New("unknown order side")

// String implements the Stringer interface.
func (s OrderSide) String() string {
	switch s {
	case Buy:
		return buy
	case Sell:
		return sell
	case BuyMinus:
		return buyMinus
	case SellPlus:
		return sellPlus
	case SellShort:
		return sellShort
	case SellShortExempt:
		return sellShortExempt
	default:
		return unknown
	}
}

// IsKnown determines if this order side is known.
func (s OrderSide) IsKnown() bool {
	return s >= Buy && s <= SellShortExempt
}

// MarshalJSON implements the Marshaler interface.
func (s OrderSide) MarshalJSON() ([]byte, error) {
	str := s.String()
	if str == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", str, errUnknownOrderSide)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(str)+extra)
	b = append(b, '"')
	b = append(b, str...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (s *OrderSide) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	str := string(d)

	switch str {
	case buy:
		*s = Buy
	case sell:
		*s = Sell
	case buyMinus:
		*s = BuyMinus
	case sellPlus:
		*s = SellPlus
	case sellShort:
		*s = SellShort
	case sellShortExempt:
		*s = SellShortExempt
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", str, errUnknownOrderSide)
	}

	return nil
}
