package entities

import (
	"fmt"
	"time"
)

// Ohlcv represents an [open, high, low, close, volume] price bar.
type Ohlcv struct {
	Time   time.Time `json:"t"` // The date and time of the closing price.
	Open   float64   `json:"o"` // The opening price.
	High   float64   `json:"h"` // The highest price.
	Low    float64   `json:"l"` // The lowest price.
	Close  float64   `json:"c"` // The closing price.
	Volume float64   `json:"v"` // The volume.
}

// IsRising indicates whether this is a rising bar, i.e. the opening price is less than the closing price.
func (o *Ohlcv) IsRising() bool {
	return o.Open < o.Close
}

// IsFalling indicates whether this is a falling bar, i.e. the closing price is less than the opening price.
func (o *Ohlcv) IsFalling() bool {
	return o.Close < o.Open
}

// Median is the median price, calculated as
//   (low + high) / 2.
func (o *Ohlcv) Median() float64 {
	return (o.Low + o.High) / 2 //nolint:gomnd
}

// Typical is the typical price, calculated as
//   (low + high + close) / 3.
func (o *Ohlcv) Typical() float64 {
	return (o.Low + o.High + o.Close) / 3 //nolint:gomnd
}

// Weighted is the weighted price, calculated as
//   (low + high + 2*close) / 4.
func (o *Ohlcv) Weighted() float64 {
	return (o.Low + o.High + o.Close + o.Close) / 4 //nolint:gomnd
}

// Average is the weighted price, calculated as
//   (low + high + open + close) / 4.
func (o *Ohlcv) Average() float64 {
	return (o.Low + o.High + o.Open + o.Close) / 4 //nolint:gomnd
}

// String implements the Stringer interface.
func (o *Ohlcv) String() string {
	return fmt.Sprintf("Ohlcv(%v, %f, %f, %f, %f, %f)", o.Time, o.Open, o.High, o.Low, o.Close, o.Volume)
}
