package currencies

import "sync"

// UpdatableConverter is a thread-safe Converter with updatable currency exchange rates.
//
// Use the NewUpdatableConverter to create a properly initialized new instance.
type UpdatableConverter struct {
	mu         sync.RWMutex
	knownRates map[Currency]map[Currency]float64
}

//nolint:exhaustivestruct
// NewUpdatableConverter creates a new empty UpdatableConverter without exchange rates.
func NewUpdatableConverter() *UpdatableConverter {
	uc := UpdatableConverter{
		knownRates: make(map[Currency]map[Currency]float64),
	}

	return &uc
}

// Convert implements Converter.
func (uc *UpdatableConverter) Convert(amount float64, base, term Currency) (converted, rate float64) {
	rate = uc.ExchangeRate(base, term)
	converted = amount * rate

	return
}

// ExchangeRate implements Converter.
func (uc *UpdatableConverter) ExchangeRate(base, term Currency) float64 {
	if base == term {
		return 1
	}

	uc.mu.RLock()
	defer uc.mu.RUnlock()

	if m, ok := uc.knownRates[base]; ok {
		if r, ok := m[term]; ok {
			return r
		}
	}

	return 0
}

// KnownBaseCurrencies implements Converter.
func (uc *UpdatableConverter) KnownBaseCurrencies(term Currency) []Currency {
	uc.mu.RLock()
	defer uc.mu.RUnlock()

	cs := make([]Currency, 0, len(uc.knownRates))

	for bc, m := range uc.knownRates {
		for tc := range m {
			if tc == term {
				cs = append(cs, bc)
			}
		}
	}

	return cs
}

// KnownTermCurrencies implements Converter.
func (uc *UpdatableConverter) KnownTermCurrencies(base Currency) []Currency {
	uc.mu.RLock()
	defer uc.mu.RUnlock()

	if m, ok := uc.knownRates[base]; ok {
		cs := make([]Currency, len(m))
		i := 0

		for k := range m {
			cs[i] = k
			i++
		}

		return cs
	}

	return []Currency{}
}

// Update adds or updates a direct exchange rate from the base currency to the term currency.
//
// X units of the base currency are equal to the X * ExchangeRate units of the term currency.
//
// If USD is the base currency and EUR is the term currency, then the currency pair USDEUR
// gives the required rate.
//
// By definition, exchange rates are positive.
// This method does not validate the value of the rate.
func (uc *UpdatableConverter) Update(base, term Currency, rate float64) {
	if base == term {
		return
	}

	uc.mu.Lock()
	defer uc.mu.Unlock()

	if m, ok := uc.knownRates[base]; ok {
		m[term] = rate
	} else {
		m := make(map[Currency]float64)
		m[term] = rate
		uc.knownRates[base] = m
	}
}
