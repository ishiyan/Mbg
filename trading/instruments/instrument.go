package instruments

import (
	"mbg/trading/currencies"
	"mbg/trading/instruments/status"
	"mbg/trading/instruments/symbology"
	"mbg/trading/instruments/types"
	"mbg/trading/markets/mics"
	"mbg/trading/time/holidays"
)

// Instrument contains properties of a financial instrument.
type Instrument interface {
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

	// PriceFactor is a positive multiplier by which price must be adjusted to determine
	// the true nominal value of a contract:
	//   Nominal Value = Quantity * Price * PriceFactor.
	PriceFactor() float64

	// Margin is an initial margin of the instrument.
	Margin() float64
}
