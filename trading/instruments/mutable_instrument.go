package instruments

import (
	"mbg/trading/currencies"
	"mbg/trading/instruments/status"
	"mbg/trading/instruments/symbology"
	"mbg/trading/instruments/types"
	"mbg/trading/markets/mics"
	"mbg/trading/time/holidays"
)

// MutableInstrument contains mutable properties of a financial instrument.
type MutableInstrument struct {
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

	// PriceFactor is a positive multiplier by which price must be adjusted to determine
	// the true nominal value of a contract:
	//   Nominal Value = Quantity * Price * PriceFactor.
	PriceFactor float64 `json:"factor,omitempty"`

	// Margin is an initial margin in the instrument's currency.
	Margin float64 `json:"margin,omitempty"`
}

// Instrument returns an immutable interface for this mutable instrument.
func (mi *MutableInstrument) Instrument() Instrument {
	return &immutableInstrument{mi}
}

type immutableInstrument struct {
	mi *MutableInstrument
}

// Name is a short name of the instrument.
func (ii *immutableInstrument) Name() string {
	return ii.mi.Name
}

// Description is a textual description of the instrument.
func (ii *immutableInstrument) Description() string {
	return ii.mi.Description
}

// Symbol (ticker) is a mnemonic of the instrument.
func (ii *immutableInstrument) Symbol() string {
	return ii.mi.Symbol
}

// ISIN is an ISO6166 (International Securities Identifying Number) code of the instrument.
func (ii *immutableInstrument) ISIN() symbology.ISIN {
	return ii.mi.ISIN
}

// CFI is an ISO 10962 (Classification of Financial Instruments) code of the instrument.
func (ii *immutableInstrument) CFI() string {
	return ii.mi.CFI
}

// MIC is an ISO 10383 Market Identifier Code where the instrument is traded.
func (ii *immutableInstrument) MIC() mics.MIC {
	return ii.mi.MIC
}

// Currency is an ISO 4217 three-letter currency code which the price of the instrument is denominated.
func (ii *immutableInstrument) Currency() currencies.Currency {
	return ii.mi.Currency
}

// Type indicates a type of the instrument.
func (ii *immutableInstrument) Type() types.InstrumentType {
	return ii.mi.Type
}

// Status indicates a state of the instrument.
func (ii *immutableInstrument) Status() status.InstrumentStatus {
	return ii.mi.Status
}

// HolidayCalendar specifies a holiday calendar of the instrument.
func (ii *immutableInstrument) HolidayCalendar() holidays.Calendar {
	return ii.mi.HolidayCalendar
}

// PricePrecision is the number of decimal places in the instruments price.
func (ii *immutableInstrument) PricePrecision() int {
	return ii.mi.PricePrecision
}

// MinPriceIncrement (tick value) is the minimum price increment of the instrument.
func (ii *immutableInstrument) MinPriceIncrement() float64 {
	return ii.mi.MinPriceIncrement
}

// PriceFactor is a positive multiplier by which price must be adjusted to determine
// the true nominal value of a contract:
//   Nominal Value = Quantity * Price * PriceFactor.
func (ii *immutableInstrument) PriceFactor() float64 {
	return ii.mi.PriceFactor
}

// Margin is an initial margin of the instrument.
func (ii *immutableInstrument) Margin() float64 {
	return ii.mi.Margin
}
