//nolint:testpackage
package portfolios

//nolint:gci
import (
	"math"
	"mbg/trading/currencies"
	"mbg/trading/instruments"
	"mbg/trading/orders"
	"mbg/trading/orders/sides"
	pside "mbg/trading/portfolios/positions/sides"
	"mbg/trading/portfolios/roundtrips/matchings"
	"testing"
	"time"
)

//nolint:funlen,gocognit
func TestPosition(t *testing.T) {
	t.Parallel()

	const (
		fmtVal            = "%v(): expected %v, actual %v"
		equalityThreshold = 1e-13
	)

	notEqual := func(a, b float64) bool {
		return math.Abs(a-b) > equalityThreshold
	}

	type expected struct {
		instrument instruments.Instrument
		currency   currencies.Currency
	}

	converter := currencies.NewUpdatableConverter()
	converter.Update(currencies.EUR, currencies.USD, 2)
	account := newAccount("foo", currencies.EUR, converter)

	/*
		eur := instruments.MutableInstrument{Currency: currencies.EUR}
		usd := instruments.MutableInstrument{Currency: currencies.USD}
		usdWithFactor := instruments.MutableInstrument{Currency: currencies.USD, PriceFactor: 2}
		usdWithMargin := instruments.MutableInstrument{Currency: currencies.USD, Margin: 1000}
	*/
	usdWithFactorAndMargin := instruments.MutableInstrument{Currency: currencies.USD, PriceFactor: 2, Margin: 1000}

	t.Run("create new position from execution", func(t *testing.T) {
		t.Parallel()

		check := func(t *testing.T, instr instruments.Instrument, ex *Execution, pos *Position) {
			if pos.Instrument() != instr {
				t.Errorf(fmtVal, "Instrument", instr, pos.Instrument())
			}

			if pos.Currency() != instr.Currency() {
				t.Errorf(fmtVal, "Currency", instr.Currency(), pos.Currency())
			}

			if notEqual(ex.Margin(), pos.Margin()) {
				t.Errorf(fmtVal, "EntryPrice", ex.Margin(), pos.Margin())
			}

			if notEqual(ex.Debt(), pos.Debt()) {
				t.Errorf(fmtVal, "Debt", ex.Debt(), pos.Debt())
			}

			leverage := 0.
			if ex.Margin() != 0 {
				leverage = ex.Amount() / ex.Margin()
			}

			if notEqual(leverage, pos.Leverage()) {
				t.Errorf(fmtVal, "Leverage", leverage, pos.Leverage())
			}

			if notEqual(ex.Price(), pos.Price()) {
				t.Errorf(fmtVal, "Price", ex.Price(), pos.Price())
			}

			qty := 0.
			if ex.Side() == sides.Buy || ex.Side() == sides.BuyMinus {
				qty = ex.Quantity()
			}

			if notEqual(qty, pos.QuantityBought()) {
				t.Errorf(fmtVal, "QuantityBought", qty, pos.QuantityBought())
			}

			qty = 0.
			if ex.Side() == sides.Sell || ex.Side() == sides.SellPlus {
				qty = ex.Quantity()
			}

			if notEqual(qty, pos.QuantitySold()) {
				t.Errorf(fmtVal, "QuantitySold", qty, pos.QuantitySold())
			}

			qty = 0.
			if ex.Side() == sides.SellShort || ex.Side() == sides.SellShortExempt {
				qty = ex.Quantity()
			}

			if notEqual(qty, pos.QuantitySoldShort()) {
				t.Errorf(fmtVal, "QuantitySoldShort", qty, pos.QuantitySoldShort())
			}

			side := pside.Long
			if ex.quantitySign < 0 {
				side = pside.Short
			}

			if side != pos.Side() {
				t.Errorf(fmtVal, "Side", side, pos.Side())
			}

			if notEqual(ex.Quantity(), pos.Quantity()) {
				t.Errorf(fmtVal, "Quantity", ex.Quantity(), pos.Quantity())
			}
		}

		instr := usdWithFactorAndMargin.Instrument()
		oser := &mockOrderSingleExecutionReport{
			id:                 "1",
			transactionTime:    time.Now(),
			lastFillPrice:      6,
			lastFillQuantity:   2,
			lastFillCommission: 3,
			commissionCurrency: currencies.USD,
			order:              orders.OrderSingle{Instrument: instr, Side: sides.Buy},
		}
		ex := newExecutionOrderSingle(oser, converter)

		pos := newPosition(instr, ex, account, matchings.FirstInFirstOut)
		check(t, instr, ex, pos)
	})
}
