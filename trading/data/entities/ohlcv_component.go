package entities

import (
	"errors"
	"fmt"
)

// OhlcvComponent defines a component of the Ohlcv type.
type OhlcvComponent int

// OhlcvFunc defines a function to get a component value from the Ohlcv type.
type OhlcvFunc func(o *Ohlcv) float64

const (
	// OhlcvOpenPrice is the opening price component.
	OhlcvOpenPrice OhlcvComponent = iota

	// OhlcvHighPrice is the highest price component.
	OhlcvHighPrice

	// OhlcvLowPrice is the lowest price component.
	OhlcvLowPrice

	// OhlcvClosePrice is the closing price component.
	OhlcvClosePrice

	// OhlcvVolume is the volume component.
	OhlcvVolume

	// OhlcvMedianPrice is the median price component, calculated as
	//   (low + high) / 2.
	OhlcvMedianPrice

	// OhlcvTypicalPrice is the typical price component, calculated as
	//   (low + high + close) / 3.
	OhlcvTypicalPrice

	// OhlcvWeightedPrice is the weighted price component, calculated as
	//   (low + high + 2*close) / 4.
	OhlcvWeightedPrice

	// OhlcvAveragePrice is the average price component, calculated as
	//   (low + high + open + close) / 4.
	OhlcvAveragePrice
)

var errUnknownOhlcvComponent = errors.New("unknown ohlcv component")

//nolint:cyclop
// OhlcvComponentFunc returns an OhlcvFunc function to get a component value from the Ohlcv type.
func OhlcvComponentFunc(c OhlcvComponent) (OhlcvFunc, error) {
	switch c {
	case OhlcvOpenPrice:
		return func(o *Ohlcv) float64 { return o.Open }, nil
	case OhlcvHighPrice:
		return func(o *Ohlcv) float64 { return o.High }, nil
	case OhlcvLowPrice:
		return func(o *Ohlcv) float64 { return o.Low }, nil
	case OhlcvClosePrice:
		return func(o *Ohlcv) float64 { return o.Close }, nil
	case OhlcvVolume:
		return func(o *Ohlcv) float64 { return o.Volume }, nil
	case OhlcvMedianPrice:
		return func(o *Ohlcv) float64 { return o.Median() }, nil
	case OhlcvTypicalPrice:
		return func(o *Ohlcv) float64 { return o.Typical() }, nil
	case OhlcvWeightedPrice:
		return func(o *Ohlcv) float64 { return o.Weighted() }, nil
	case OhlcvAveragePrice:
		return func(o *Ohlcv) float64 { return o.Average() }, nil
	default:
		return nil, fmt.Errorf("%d: %w", int(c), errUnknownOhlcvComponent)
	}
}

//nolint:cyclop
// String implements the Stringer interface.
func (c OhlcvComponent) String() string {
	switch c {
	case OhlcvOpenPrice:
		return "OpenPrice"
	case OhlcvHighPrice:
		return "HighPrice"
	case OhlcvLowPrice:
		return "LowPrice"
	case OhlcvClosePrice:
		return "ClosePrice"
	case OhlcvVolume:
		return "Volume"
	case OhlcvMedianPrice:
		return "MedianPrice"
	case OhlcvTypicalPrice:
		return "TypicalPrice"
	case OhlcvWeightedPrice:
		return "WeightedPrice"
	case OhlcvAveragePrice:
		return "AveragePrice"
	default:
		return fmt.Errorf("%d: %w", int(c), errUnknownOhlcvComponent).Error()
	}
}
