package entities

import (
	"errors"
	"fmt"
)

// QuoteComponent defines a component of the Quote type.
type QuoteComponent int

// QuoteFunc defines a function to get a component value from the Quote type.
type QuoteFunc func(q *Quote) float64

const (
	// QuoteBidPrice is the bid price component.
	QuoteBidPrice QuoteComponent = iota

	// QuoteAskPrice is the ask price component.
	QuoteAskPrice

	// QuoteBidSize is the bid size component.
	QuoteBidSize

	// QuoteAskSize is the ask size component.
	QuoteAskSize

	// QuoteMidPrice is the med-price component, calculated as
	//   (ask + bid) / 2.
	QuoteMidPrice

	// QuoteWeightedPrice is the weighted price component, calculated as
	//   (ask*askSize + bid*bidSize) / (askSize + bidSize).
	QuoteWeightedPrice

	// QuoteWeightedMidPrice is the weighted mid-price component (sometimes called micro-price), calculated as
	//   (ask*bidSize + bid*askSize) / (askSize + bidSize).
	QuoteWeightedMidPrice

	// QuoteSpreadBp is the spread in basis points (100 basis points = 1%) component, calculated as
	//   10000 * (ask - bid) / mid.
	QuoteSpreadBp
)

var errUnknownQuoteComponent = errors.New("unknown quote component")

// QuoteComponentFunc returns a QuoteFunc function to get a component value from the Quote type.
func QuoteComponentFunc(c QuoteComponent) (QuoteFunc, error) {
	switch c {
	case QuoteBidPrice:
		return func(q *Quote) float64 { return q.Bid }, nil
	case QuoteAskPrice:
		return func(q *Quote) float64 { return q.Ask }, nil
	case QuoteBidSize:
		return func(q *Quote) float64 { return q.BidSize }, nil
	case QuoteAskSize:
		return func(q *Quote) float64 { return q.AskSize }, nil
	case QuoteMidPrice:
		return func(q *Quote) float64 { return q.Mid() }, nil
	case QuoteWeightedPrice:
		return func(q *Quote) float64 { return q.Weighted() }, nil
	case QuoteWeightedMidPrice:
		return func(q *Quote) float64 { return q.WeightedMid() }, nil
	case QuoteSpreadBp:
		return func(q *Quote) float64 { return q.SpreadBp() }, nil
	default:
		return nil, fmt.Errorf("%d: %w", int(c), errUnknownQuoteComponent)
	}
}

// String implements the Stringer interface.
func (c QuoteComponent) String() string {
	switch c {
	case QuoteBidPrice:
		return "BidPrice"
	case QuoteAskPrice:
		return "AskPrice"
	case QuoteBidSize:
		return "BidSize"
	case QuoteAskSize:
		return "AskSize"
	case QuoteMidPrice:
		return "MidPrice"
	case QuoteWeightedPrice:
		return "WeightedPrice"
	case QuoteWeightedMidPrice:
		return "WeightedMidPrice"
	case QuoteSpreadBp:
		return "SpreadBp"
	default:
		return fmt.Errorf("%d: %w", int(c), errUnknownQuoteComponent).Error()
	}
}
