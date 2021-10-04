//nolint:testpackage
package portfolios

//nolint:gci
import (
	"mbg/trading/currencies"
	"mbg/trading/instruments"
	"mbg/trading/orders"
	"mbg/trading/orders/sides"
	"testing"
	"time"
)

//nolint:funlen,gocognit
func TestExecution(t *testing.T) {
	t.Parallel()

	converter := currencies.NewUpdatableConverter()
	converter.Update(currencies.EUR, currencies.USD, 2)

	verify := func(exp orders.OrderSingleExecutionReport, act *Execution) {
		t.Helper()

		if exp.ID() != act.ReportID() {
			t.Errorf("ReportID(): expected %v, actual %v", exp.ID(), act.ReportID())
		}

		if exp.TransactionTime() != act.ReportTime() {
			t.Errorf("ReportTime(): expected %v, actual %v", exp.TransactionTime(), act.ReportTime())
		}

		if exp.Order().Side != act.Side() {
			t.Errorf("Side(): expected %v, actual %v", exp.Order().Side, act.Side())
		}

		if exp.LastFillQuantity() != act.Quantity() {
			t.Errorf("Quantity(): expected %v, actual %v", exp.LastFillQuantity(), act.Quantity())
		}

		if exp.Order().Instrument.Currency() != act.Currency() {
			t.Errorf("Currency(): expected %v, actual %v", exp.Order().Instrument.Currency(), act.Currency())
		}

		if exp.CommissionCurrency() != act.CommissionCurrency() {
			t.Errorf("CommissionCurrency(): expected %v, actual %v", exp.CommissionCurrency(), act.CommissionCurrency())
		}

		if exp.LastFillCommission() != act.Commission() {
			t.Errorf("Commission(): expected %v, actual %v", exp.LastFillCommission(), act.Commission())
		}

		c, r := converter.Convert(act.Commission(), act.CommissionCurrency(), act.Currency())
		if r != act.ConversionRate() {
			t.Errorf("ConversionRate(): expected %v, actual %v", r, act.ConversionRate())
		}

		if c != act.CommissionConverted() {
			t.Errorf("CommissionConverted(): expected %v, actual %v", c, act.CommissionConverted())
		}

		if -c != act.PnL() {
			t.Errorf("PnL(): expected %v, actual %v", -c, act.PnL())
		}

		if 0 != act.RealizedPnL() {
			t.Errorf("RealizedPnL(): expected %v, actual %v", 0, act.RealizedPnL())
		}

		if exp.LastFillPrice() != act.Price() {
			t.Errorf("Price(): expected %v, actual %v", exp.LastFillPrice(), act.Price())
		}

		p := exp.Order().Instrument.PriceFactor() * exp.LastFillPrice()
		if p == 0 {
			p = exp.LastFillPrice()
		}

		sign := 1.
		if exp.Order().Side.IsSell() {
			sign = -1.
		}

		a := p * exp.LastFillQuantity()
		if a != act.Amount() {
			t.Errorf("Amount(): expected %v, actual %v", a, act.Amount())
		}

		if -sign*a != act.CashFlow() {
			t.Errorf("CashFlow(): expected %v, actual %v", -sign*a, act.CashFlow())
		}

		m := exp.Order().Instrument.Margin() * exp.LastFillQuantity()
		if m != act.Margin() {
			t.Errorf("Margin(): expected %v, actual %v", m, act.Margin())
		}

		d := a - m
		if m == 0 {
			d = 0
		}

		if d != act.Debt() {
			t.Errorf("Debt(): expected %v, actual %v", d, act.Debt())
		}
	}

	eur := instruments.MutableInstrument{Currency: currencies.EUR}
	usd := instruments.MutableInstrument{Currency: currencies.USD}
	usdWithFactor := instruments.MutableInstrument{Currency: currencies.USD, PriceFactor: 2}
	usdWithMargin := instruments.MutableInstrument{Currency: currencies.USD, Margin: 1000}
	usdFull := instruments.MutableInstrument{Currency: currencies.USD, PriceFactor: 2, Margin: 1000}

	t.Run("buy without commission conversion", func(t *testing.T) {
		t.Parallel()

		oser := &mockOrderSingleExecutionReport{
			id: "1", transactionTime: time.Now(),
			lastFillPrice: 6, lastFillQuantity: 2, lastFillCommission: 3, commissionCurrency: currencies.EUR,
			order: orders.OrderSingle{Instrument: eur.Instrument(), Side: sides.Buy},
		}

		ex := newExecutionOrderSingle(oser, converter)
		verify(oser, ex)
	})

	t.Run("buy with commission conversion", func(t *testing.T) {
		t.Parallel()

		oser := &mockOrderSingleExecutionReport{
			id: "2", transactionTime: time.Now(),
			lastFillPrice: 6, lastFillQuantity: 2, lastFillCommission: 3, commissionCurrency: currencies.EUR,
			order: orders.OrderSingle{Instrument: usd.Instrument(), Side: sides.Buy},
		}

		ex := newExecutionOrderSingle(oser, converter)
		verify(oser, ex)
	})

	t.Run("buy with commission conversion and price factor", func(t *testing.T) {
		t.Parallel()

		oser := &mockOrderSingleExecutionReport{
			id: "3", transactionTime: time.Now(),
			lastFillPrice: 6, lastFillQuantity: 2, lastFillCommission: 3, commissionCurrency: currencies.EUR,
			order: orders.OrderSingle{Instrument: usdWithFactor.Instrument(), Side: sides.Buy},
		}

		ex := newExecutionOrderSingle(oser, converter)
		verify(oser, ex)
	})

	t.Run("buy with commission conversion and margin", func(t *testing.T) {
		t.Parallel()

		oser := &mockOrderSingleExecutionReport{
			id: "4", transactionTime: time.Now(),
			lastFillPrice: 6, lastFillQuantity: 2, lastFillCommission: 3, commissionCurrency: currencies.EUR,
			order: orders.OrderSingle{Instrument: usdWithMargin.Instrument(), Side: sides.Buy},
		}

		ex := newExecutionOrderSingle(oser, converter)
		verify(oser, ex)
	})

	t.Run("buy", func(t *testing.T) {
		t.Parallel()

		oser := &mockOrderSingleExecutionReport{
			id: "5", transactionTime: time.Now(),
			lastFillPrice: 6, lastFillQuantity: 2, lastFillCommission: 3, commissionCurrency: currencies.EUR,
			order: orders.OrderSingle{Instrument: usdFull.Instrument(), Side: sides.Buy},
		}

		ex := newExecutionOrderSingle(oser, converter)
		verify(oser, ex)
	})

	t.Run("buy minus", func(t *testing.T) {
		t.Parallel()

		oser := &mockOrderSingleExecutionReport{
			id: "6", transactionTime: time.Now(),
			lastFillPrice: 6, lastFillQuantity: 2, lastFillCommission: 3, commissionCurrency: currencies.EUR,
			order: orders.OrderSingle{Instrument: usdFull.Instrument(), Side: sides.BuyMinus},
		}

		ex := newExecutionOrderSingle(oser, converter)
		verify(oser, ex)
	})

	t.Run("sell", func(t *testing.T) {
		t.Parallel()

		oser := &mockOrderSingleExecutionReport{
			id: "7", transactionTime: time.Now(),
			lastFillPrice: 6, lastFillQuantity: 2, lastFillCommission: 3, commissionCurrency: currencies.EUR,
			order: orders.OrderSingle{Instrument: usdFull.Instrument(), Side: sides.Sell},
		}

		ex := newExecutionOrderSingle(oser, converter)
		verify(oser, ex)
	})

	t.Run("sell plus", func(t *testing.T) {
		t.Parallel()

		oser := &mockOrderSingleExecutionReport{
			id: "8", transactionTime: time.Now(),
			lastFillPrice: 6, lastFillQuantity: 2, lastFillCommission: 3, commissionCurrency: currencies.EUR,
			order: orders.OrderSingle{Instrument: usdFull.Instrument(), Side: sides.SellPlus},
		}

		ex := newExecutionOrderSingle(oser, converter)
		verify(oser, ex)
	})

	t.Run("sell short", func(t *testing.T) {
		t.Parallel()

		oser := &mockOrderSingleExecutionReport{
			id: "9", transactionTime: time.Now(),
			lastFillPrice: 6, lastFillQuantity: 2, lastFillCommission: 3, commissionCurrency: currencies.EUR,
			order: orders.OrderSingle{Instrument: usdFull.Instrument(), Side: sides.SellShort},
		}

		ex := newExecutionOrderSingle(oser, converter)
		verify(oser, ex)
	})

	t.Run("sell short exempt", func(t *testing.T) {
		t.Parallel()

		oser := &mockOrderSingleExecutionReport{
			id: "10", transactionTime: time.Now(),
			lastFillPrice: 6, lastFillQuantity: 2, lastFillCommission: 3, commissionCurrency: currencies.EUR,
			order: orders.OrderSingle{Instrument: usdFull.Instrument(), Side: sides.SellShortExempt},
		}

		ex := newExecutionOrderSingle(oser, converter)
		verify(oser, ex)
	})
}
