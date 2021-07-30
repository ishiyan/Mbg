//nolint:testpackage
package monitorings

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	act := QuoteTradeBar
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkQuotes(b *testing.B) {
	act := QuoteTradeBar
	for i := 0; i < b.N; i++ {
		_ = act.Quotes()
	}
}

func BenchmarkTrades(b *testing.B) {
	act := QuoteTradeBar
	for i := 0; i < b.N; i++ {
		_ = act.Trades()
	}
}

func BenchmarkBars(b *testing.B) {
	act := QuoteTradeBar
	for i := 0; i < b.N; i++ {
		_ = act.Bars()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := QuoteTradeBar
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var m Monitoring

	bs := []byte("\"quoteTradeBar\"")
	for i := 0; i < b.N; i++ {
		_ = m.UnmarshalJSON(bs)
	}
}
