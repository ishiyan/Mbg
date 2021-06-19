package entities

import (
	"errors"
	"fmt"
)

// TradeComponent defines a component of the Trade type.
type TradeComponent int

// TradeFunc defines a function to get a component value from the Trade type.
type TradeFunc func(t *Trade) float64

const (
	// TradePrice is the price component.
	TradePrice TradeComponent = iota

	// TradeVolume is the volume component.
	TradeVolume
)

var errUnknownTradeComponent = errors.New("unknown trade component")

// TradeComponentFunc returns a TradeFunc function to get a component value from the Trade type.
func TradeComponentFunc(c TradeComponent) (TradeFunc, error) {
	switch c {
	case TradePrice:
		return func(t *Trade) float64 { return t.Price }, nil
	case TradeVolume:
		return func(t *Trade) float64 { return t.Volume }, nil
	default:
		return nil, fmt.Errorf("%d: %w", int(c), errUnknownTradeComponent)
	}
}

// String implements the Stringer interface.
func (c TradeComponent) String() string {
	switch c {
	case TradePrice:
		return "Price"
	case TradeVolume:
		return "Volume"
	default:
		return fmt.Errorf("%d: %w", int(c), errUnknownTradeComponent).Error()
	}
}
