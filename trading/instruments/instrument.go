package instruments

import (
	"mbg/trading/currencies"
	"mbg/trading/instruments/status"
	"mbg/trading/instruments/symbology"
	"mbg/trading/instruments/types"
	"mbg/trading/markets/mics"
	"mbg/trading/time/holidays"
)

// Instrument contains mutable properties of a financial instrument.
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

	// PriceFactor is a positive factor by which price must be adjusted to determine
	// the true nominal value of a contract:
	//   Nominal Value = Quantity * Price * PriceMultiplier.
	PriceFactor float64 `json:"factor,omitempty"`

	// Margin is an initial margin in the instrument's currency.
	Margin float64 `json:"margin,omitempty"`
}

func (i *Instrument) Freeze() ImmutableInstrument {
	return &immutableInstrument{i}
}

// ImmutableInstrument contains immutable properties of a financial instrument.
type ImmutableInstrument interface {
	// Name is a short name of the instrument.
	Name() string

	// Description is a textual description of the instrument.
	Description() string

	// Symbol (ticker) is a mnemonic of the instrument.
	Symbol() string

	// ISIN is an ISO6166 (International Securities Identifying Number) code of the instrument.
	ISIN() symbology.ISIN

	// CFI is an ISO 10962 (Classification of Financial Instruments) code of the instrument.
	CFI() string

	// MIC is an ISO 10383 Market Identifier Code where the instrument is traded.
	MIC() mics.MIC

	// Currency is an ISO 4217 three-letter currency code which the price of the instrument is denominated.
	Currency() currencies.Currency

	// Type indicates a type of the instrument.
	Type() types.InstrumentType

	// Status indicates a state of the instrument.
	Status() status.InstrumentStatus

	// HolidayCalendar specifies a holiday calendar of the instrument.
	HolidayCalendar() holidays.Calendar

	// PricePrecision is the number of decimal places in the instruments price.
	PricePrecision() int

	// MinPriceIncrement (tick value) is the minimum price increment of the instrument.
	MinPriceIncrement() float64

	// PriceFactor is a positive factor by which price must be adjusted to determine
	// the true nominal value of a contract:
	//   Nominal Value = Quantity * Price * PriceMultiplier.
	PriceFactor() float64

	// Margin is an initial margin of the instrument.
	Margin() float64
}

type immutableInstrument struct {
	ins *Instrument
}

func (ii *immutableInstrument) Name() string {
	return ii.ins.Name
}

func (ii *immutableInstrument) Description() string {
	return ii.ins.Description
}

func (ii *immutableInstrument) Symbol() string {
	return ii.ins.Symbol
}

func (ii *immutableInstrument) ISIN() symbology.ISIN {
	return ii.ins.ISIN
}

func (ii *immutableInstrument) CFI() string {
	return ii.ins.CFI
}

func (ii *immutableInstrument) MIC() mics.MIC {
	return ii.ins.MIC
}

func (ii *immutableInstrument) Currency() currencies.Currency {
	return ii.ins.Currency
}

func (ii *immutableInstrument) Type() types.InstrumentType {
	return ii.ins.Type
}

func (ii *immutableInstrument) Status() status.InstrumentStatus {
	return ii.ins.Status
}

func (ii *immutableInstrument) HolidayCalendar() holidays.Calendar {
	return ii.ins.HolidayCalendar
}

func (ii *immutableInstrument) PricePrecision() int {
	return ii.ins.PricePrecision
}

func (i *immutableInstrument) MinPriceIncrement() float64 {
	return i.ins.MinPriceIncrement
}

func (i *immutableInstrument) PriceFactor() float64 {
	if i.ins.PriceFactor == 0 {
		return 1
	}

	return i.ins.PriceFactor
}

func (i *immutableInstrument) Margin() float64 {
	return i.ins.Margin
}
