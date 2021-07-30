// Package monitorings enumerates portfolio position sell price monitoring combinations.
package monitorings

import (
	"bytes"
	"errors"
	"fmt"
)

// Monitoring enumerates portfolio position sell price monitoring combinations.
type Monitoring int

const (
	// None means no monitoring.
	None Monitoring = iota // 000

	// Quote means monitoring on quotes.
	Quote // 001

	// Trade means monitoring on trades.
	Trade // 010

	// QuoteTrade means monitoring on quotes and trades.
	QuoteTrade // 011

	// Bar means monitoring on bars.
	Bar // 100

	// QuoteBar means monitoring on quotes and bars.
	QuoteBar // 101

	// TradeBar means monitoring on trades and bars.
	TradeBar // 110

	// QuoteTradeBar means monitoring on quotes, trades and bars.
	QuoteTradeBar // 111
)

const (
	unknown       = "unknown"
	none          = "none"
	quote         = "quote"
	trade         = "trade"
	bar           = "bar"
	quoteTrade    = "quoteTrade"
	quoteBar      = "quoteBar"
	tradeBar      = "tradeBar"
	quoteTradeBar = "quoteTradeBar"
)

var errUnknownMonitoring = errors.New("unknown sell price monitoring")

// String implements the fmt.Stringer interface.
func (m Monitoring) String() string {
	switch m {
	case None:
		return none
	case Quote:
		return quote
	case Trade:
		return trade
	case Bar:
		return bar
	case QuoteTrade:
		return quoteTrade
	case QuoteBar:
		return quoteBar
	case TradeBar:
		return tradeBar
	case QuoteTradeBar:
		return quoteTradeBar
	default:
		return unknown
	}
}

// IsKnown determines if this sell price monitoring is known.
func (m Monitoring) IsKnown() bool {
	return m >= None && m <= QuoteTradeBar
}

//nolint:exhaustive
// Quotes determines if this sell price monitoring monitors quotes.
func (m Monitoring) Quotes() bool {
	switch m {
	case Quote, QuoteTrade, QuoteBar, QuoteTradeBar:
		return true
	default:
		return false
	}
}

//nolint:exhaustive
// Trades determines if this sell price monitoring monitors trades.
func (m Monitoring) Trades() bool {
	switch m {
	case Trade, QuoteTrade, TradeBar, QuoteTradeBar:
		return true
	default:
		return false
	}
}

//nolint:exhaustive
// Bars determines if this sell price monitoring monitors bars.
func (m Monitoring) Bars() bool {
	switch m {
	case Bar, TradeBar, QuoteBar, QuoteTradeBar:
		return true
	default:
		return false
	}
}

// MarshalJSON implements the Marshaler interface.
func (m Monitoring) MarshalJSON() ([]byte, error) {
	str := m.String()
	if str == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", str, errUnknownMonitoring)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(str)+extra)
	b = append(b, '"')
	b = append(b, str...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (m *Monitoring) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	str := string(d)

	switch str {
	case none:
		*m = None
	case quote:
		*m = Quote
	case trade:
		*m = Trade
	case bar:
		*m = Bar
	case quoteTrade:
		*m = QuoteTrade
	case quoteBar:
		*m = QuoteBar
	case tradeBar:
		*m = TradeBar
	case quoteTradeBar:
		*m = QuoteTradeBar
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", str, errUnknownMonitoring)
	}

	return nil
}
