//nolint:testpackage
package monitorings

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		m    Monitoring
		text string
	}{
		{None, none},
		{Quote, quote},
		{Trade, trade},
		{Bar, bar},
		{QuoteTrade, quoteTrade},
		{QuoteBar, quoteBar},
		{TradeBar, tradeBar},
		{QuoteTradeBar, quoteTradeBar},
		{Monitoring(9999), unknown},
		{Monitoring(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.m.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.m, exp, act)
		}
	}
}

func TestIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		m       Monitoring
		boolean bool
	}{
		{None, true},
		{Quote, true},
		{Trade, true},
		{Bar, true},
		{QuoteTrade, true},
		{QuoteBar, true},
		{TradeBar, true},
		{QuoteTradeBar, true},
		{Monitoring(9999), false},
		{Monitoring(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.m.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.m, exp, act)
		}
	}
}

func TestFlags(t *testing.T) {
	t.Parallel()

	tests := []struct {
		m        Monitoring
		hasQuote bool
		hasTrade bool
		hasBar   bool
	}{
		{None, false, false, false},
		{Quote, true, false, false},
		{Trade, false, true, false},
		{Bar, false, false, true},
		{QuoteTrade, true, true, false},
		{QuoteBar, true, false, true},
		{TradeBar, false, true, true},
		{QuoteTradeBar, true, true, true},
		{Monitoring(9999), false, false, false},
		{Monitoring(-9999), false, false, false},
	}

	for _, tt := range tests {
		exp := tt.hasQuote
		act := tt.m.Quotes()

		if exp != act {
			t.Errorf("'%v'|Quote: expected '%v', actual '%v'", tt.m, exp, act)
		}

		exp = tt.hasTrade
		act = tt.m.Trades()

		if exp != act {
			t.Errorf("'%v'|Trade: expected '%v', actual '%v'", tt.m, exp, act)
		}

		exp = tt.hasBar
		act = tt.m.Bars()

		if exp != act {
			t.Errorf("'%v'|Bar: expected '%v', actual '%v'", tt.m, exp, act)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		m         Monitoring
		json      string
		succeeded bool
	}{
		{None, "\"none\"", true},
		{Quote, "\"quote\"", true},
		{Trade, "\"trade\"", true},
		{Bar, "\"bar\"", true},
		{QuoteTrade, "\"quoteTrade\"", true},
		{QuoteBar, "\"quoteBar\"", true},
		{TradeBar, "\"tradeBar\"", true},
		{QuoteTradeBar, "\"quoteTradeBar\"", true},
		{Monitoring(9999), nilstr, false},
		{Monitoring(-9999), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.m.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %w", tt.m, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.m)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.m, exp, act)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero Monitoring
	tests := []struct {
		m         Monitoring
		json      string
		succeeded bool
	}{
		{None, "\"none\"", true},
		{Quote, "\"quote\"", true},
		{Trade, "\"trade\"", true},
		{Bar, "\"bar\"", true},
		{QuoteTrade, "\"quoteTrade\"", true},
		{QuoteBar, "\"quoteBar\"", true},
		{TradeBar, "\"tradeBar\"", true},
		{QuoteTradeBar, "\"quoteTradeBar\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.m
		bs := []byte(tt.json)

		var m Monitoring

		err := m.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %w", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != m {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, m)
		}
	}
}
