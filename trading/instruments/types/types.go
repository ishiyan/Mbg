// Package types enumerates types of an instrument.
package types

import (
	"bytes"
	"errors"
	"fmt"
)

// InstrumentType enumerates types of an instrument.
type InstrumentType int

const (
	// Undefined instrument type.
	Undefined InstrumentType = iota + 1

	// Stock is a security that denote an ownership in a public company.
	Stock

	// Index tracks the performance of a group of assets in a standardized way.
	// Indexes typically measure the performance of a basket of securities
	// intended to replicate a certain area of the market.
	Index

	// INAV is an intraday indicative net asset value of an ETF or ETV
	// based on the market values of its underlying constituents.
	INAV

	// ETF is an exchange traded fund, a security that tracks a basket of assets.
	ETF

	// ETC is an exchange traded commodity, a security thet tracks the price of a commodity or a commodity bucket.
	// It is backed by an underwritten note, but that note is collateralized by physical commodities.
	ETC

	// Forex is a currency instrument.
	Forex

	// Crypto is a crypto currency instrument.
	Crypto
	last
)

const (
	unknown   = "unknown"
	undefined = "undefined"
	stock     = "stock"
	index     = "index"
	inav      = "inav"
	etf       = "etf"
	etc       = "etc"
	forex     = "forex"
	crypto    = "crypto"
)

var errUnknownInstrumentType = errors.New("unknown instrument type")

//nolint:exhaustive
// String implements the Stringer interface.
func (t InstrumentType) String() string {
	switch t {
	case Undefined:
		return undefined
	case Stock:
		return stock
	case Index:
		return index
	case INAV:
		return inav
	case ETF:
		return etf
	case ETC:
		return etc
	case Forex:
		return forex
	case Crypto:
		return crypto
	default:
		return unknown
	}
}

// IsKnown determines if this instrument type is known.
func (t InstrumentType) IsKnown() bool {
	return t > Undefined && t < last
}

// MarshalJSON implements the Marshaler interface.
func (t InstrumentType) MarshalJSON() ([]byte, error) {
	s := t.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownInstrumentType)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(s)+extra)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (t *InstrumentType) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	s := string(d)

	switch s {
	case stock:
		*t = Stock
	case index:
		*t = Index
	case inav:
		*t = INAV
	case etf:
		*t = ETF
	case etc:
		*t = ETC
	case forex:
		*t = Forex
	case crypto:
		*t = Crypto
	case undefined:
		*t = Undefined
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", s, errUnknownInstrumentType)
	}

	return nil
}
