package entities

import (
	"errors"
	"fmt"
)

// BarComponent defines a component of the Bar type.
type BarComponent int

// BarFunc defines a function to get a component value from the Bar type.
type BarFunc func(b *Bar) float64

const (
	// BarOpenPrice is the opening price component.
	BarOpenPrice BarComponent = iota

	// BarHighPrice is the highest price component.
	BarHighPrice

	// BarLowPrice is the lowest price component.
	BarLowPrice

	// BarClosePrice is the closing price component.
	BarClosePrice

	// BarVolume is the volume component.
	BarVolume

	// BarMedianPrice is the median price component, calculated as
	//   (low + high) / 2.
	BarMedianPrice

	// BarTypicalPrice is the typical price component, calculated as
	//   (low + high + close) / 3.
	BarTypicalPrice

	// BarWeightedPrice is the weighted price component, calculated as
	//   (low + high + 2*close) / 4.
	BarWeightedPrice

	// BarAveragePrice is the average price component, calculated as
	//   (low + high + open + close) / 4.
	BarAveragePrice
)

var errUnknownBarComponent = errors.New("unknown bar component")

// BarComponentFunc returns an BarFunc function to get a component value from the Bar type.
func BarComponentFunc(c BarComponent) (BarFunc, error) {
	switch c {
	case BarOpenPrice:
		return func(b *Bar) float64 { return b.Open }, nil
	case BarHighPrice:
		return func(b *Bar) float64 { return b.High }, nil
	case BarLowPrice:
		return func(b *Bar) float64 { return b.Low }, nil
	case BarClosePrice:
		return func(b *Bar) float64 { return b.Close }, nil
	case BarVolume:
		return func(b *Bar) float64 { return b.Volume }, nil
	case BarMedianPrice:
		return func(b *Bar) float64 { return b.Median() }, nil
	case BarTypicalPrice:
		return func(b *Bar) float64 { return b.Typical() }, nil
	case BarWeightedPrice:
		return func(b *Bar) float64 { return b.Weighted() }, nil
	case BarAveragePrice:
		return func(b *Bar) float64 { return b.Average() }, nil
	default:
		return nil, fmt.Errorf("%d: %w", int(c), errUnknownBarComponent)
	}
}

// String implements the Stringer interface.
func (c BarComponent) String() string {
	switch c {
	case BarOpenPrice:
		return "OpenPrice"
	case BarHighPrice:
		return "HighPrice"
	case BarLowPrice:
		return "LowPrice"
	case BarClosePrice:
		return "ClosePrice"
	case BarVolume:
		return "Volume"
	case BarMedianPrice:
		return "MedianPrice"
	case BarTypicalPrice:
		return "TypicalPrice"
	case BarWeightedPrice:
		return "WeightedPrice"
	case BarAveragePrice:
		return "AveragePrice"
	default:
		return fmt.Errorf("%d: %w", int(c), errUnknownBarComponent).Error()
	}
}
