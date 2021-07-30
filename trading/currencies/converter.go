package currencies

// Converter converts base currency units to term currency units.
// Exchange rates may change over time.
type Converter interface {
	// Convert converts the amount in the base currency to the converted amount in the term currency.
	//
	// If either base or term currency is unknown, the converted amount is zero.
	// If both currencies are the same, the amounts are also the same.
	Convert(amount float64, base, term Currency) (converted, rate float64)

	// ExchangeRate returns a direct exchange rate from the base currency to the term currency.
	//
	// X units of the base currency are equal to the X * ExchangeRate units of the term currency.
	//
	// If USD is the base currency and EUR is the term currency, then the currency pair USDEUR
	// gives the required rate.
	//
	// If either base or term currencies are unknown, the exchange rate is zero.
	// If both currencies are the same, the exchange rate is 1.
	ExchangeRate(base, term Currency) float64

	// KnownBaseCurrencies returns a collection of base currencies with known exchange rates
	// for a given term currency.
	KnownBaseCurrencies(term Currency) []Currency

	// KnownTermCurrencies returns a collection of term currencies with known exchange rates
	// for a given base currency.
	KnownTermCurrencies(base Currency) []Currency
}
