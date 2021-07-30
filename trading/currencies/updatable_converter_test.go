//nolint:testpackage
package currencies

import (
	"testing"
)

const (
	amount = 123.456
	gbpEur = 1.14
	gbpChf = 1.73
	eurChf = 1.51
	chfEur = 0.66
)

func TestWithoutRates(t *testing.T) {
	t.Parallel()

	var c Converter = NewUpdatableConverter()

	if rate := c.ExchangeRate(GBP, EUR); rate != 0 {
		t.Errorf("ExchangeRate(GBP, EUR): expected zero rate, actual %v", rate)
	}

	if rate := c.ExchangeRate(GBP, GBP); rate != 1 {
		t.Errorf("ExchangeRate(GBP, GBP): expected rate 1, actual %v", rate)
	}

	if conv, rate := c.Convert(amount, GBP, EUR); conv != 0 || rate != 0 {
		t.Errorf("Convert(%v, GBP, EUR): expected zero rate and amount, actual %v and %v",
			amount, rate, conv)
	}

	if known := c.KnownBaseCurrencies(GBP); len(known) != 0 {
		t.Errorf("KnownBaseCurrencies(): expected empty collection, actual length %v", len(known))
	}

	if known := c.KnownTermCurrencies(EUR); len(known) != 0 {
		t.Errorf("KnownTermCurrencies(): expected empty collection, actual length %v", len(known))
	}
}

func TestWithRates(t *testing.T) {
	t.Parallel()

	uc := NewUpdatableConverter()
	uc.Update(GBP, EUR, gbpEur)
	uc.Update(GBP, CHF, gbpChf)
	uc.Update(CHF, EUR, chfEur)
	uc.Update(EUR, CHF, eurChf)

	var c Converter = uc

	if rate := c.ExchangeRate(GBP, CHF); rate != gbpChf {
		t.Errorf("ExchangeRate(GBP, CHF): expected rate %v, actual %v", gbpChf, rate)
	}

	uc.Update(CHF, CHF, 123.456)

	if rate := c.ExchangeRate(CHF, CHF); rate != 1 {
		t.Errorf("ExchangeRate(CHF, CHF): expected rate 1, actual %v", rate)
	}

	if conv, rate := c.Convert(amount, GBP, EUR); rate != gbpEur || conv != gbpEur*amount {
		t.Errorf("Convert(%v, GBP, EUR): expected rate %v and amount %v, actual %v and %v",
			amount, gbpEur, gbpEur*amount, rate, conv)
	}

	if known := c.KnownBaseCurrencies(CHF); !hasTwo(known, GBP, EUR) {
		t.Errorf("KnownBaseCurrencies(CHF): expected {GBP, EUR}, actual %v", known)
	}

	if known := c.KnownTermCurrencies(GBP); !hasTwo(known, EUR, CHF) {
		t.Errorf("KnownTermCurrencies(GBP): expected {EUR, CHF}, actual %v", known)
	}
}

func hasTwo(s []Currency, c1, c2 Currency) bool {
	if len(s) == 2 {
		if s[0] == c1 {
			return s[1] == c2
		} else if s[0] == c2 {
			return s[1] == c1
		}
	}

	return false
}
