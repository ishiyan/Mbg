package instruments

import (
	"mbg/trading/currencies"
	"mbg/trading/instruments/status"
	"mbg/trading/instruments/symbology"
	"mbg/trading/instruments/types"
	"mbg/trading/markets/mics"
	"mbg/trading/time/holidays"
)

type Instrument struct {
	// Name is a short name of the instrument.
	Name string `json:"name,omitempty"`

	// Description is a textual description of the instrument.
	Description string `json:"description,omitempty"`

	// Symbol (ticker) is a mnemonic of the instrument.
	Symbol string `json:"symbol,omitempty"`

	// ISIN is an ISO6166 (International Securities Identifying Number) code of the instrument.
	ISIN symbology.ISIN `json:"isin,omitempty"`

	// CFI is an ISO 10962 (Classification of Financial Instruments) code of the instrument.
	CFI string `json:"cfi,omitempty"`

	// MIC is an ISO 10383 Market Identifier Code where the instrument is traded.
	MIC mics.MIC `json:"mic,omitempty"`

	// Currency is an ISO 4217 three-letter currency code which the price of the instrument is denominated.
	Currency currencies.Currency `json:"currency,omitempty"`

	// Type indicates a type of the instrument.
	Type types.InstrumentType `json:"type,omitempty"`

	// Status indicates a state of the instrument.
	Status status.InstrumentStatus `json:"status,omitempty"`

	// HolidayCalendar specifies a holiday calendar of the instrument.
	HolidayCalendar holidays.Calendar `json:"holidayCalendar,omitempty"`

	// PricePrecision is the number of decimal places in the instruments price.
	PricePrecision int `json:"pricePrecision,omitempty"`

	// MinPriceIncrement (tick value) is the minimum price increment of the instrument.
	MinPriceIncrement float64 `json:"minPriceIncrement,omitempty"`

	// Margin is an initial margin of the instrument.
	Margin float64 `json:"margin,omitempty"`
}
