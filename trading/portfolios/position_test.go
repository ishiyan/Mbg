//nolint:testpackage
package portfolios

//nolint:gofumpt
import (
	"math"
	"testing"
	"time"

	"mbg/trading/currencies"
	"mbg/trading/instruments"
	"mbg/trading/orders"
	"mbg/trading/orders/sides"
	pside "mbg/trading/portfolios/positions/sides"
	"mbg/trading/portfolios/roundtrips/matchings"
)

//nolint:funlen,gocognit,maintidx
func TestPosition(t *testing.T) {
	t.Parallel()

	const (
		fmtVal            = "%v(): expected %v, actual %v"
		fmtLen            = "%v(): expected length %v, actual %v"
		fmtElem           = "%v()[%d]: expected %v, actual %v"
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

	/*
		eur := instruments.MutableInstrument{Currency: currencies.EUR}
		usd := instruments.MutableInstrument{Currency: currencies.USD}
		usdWithFactor := instruments.MutableInstrument{Currency: currencies.USD, PriceFactor: 2}
		usdWithMargin := instruments.MutableInstrument{Currency: currencies.USD, Margin: 1000}
	*/
	usdWithFactorAndMargin := instruments.MutableInstrument{Currency: currencies.USD, PriceFactor: 2, Margin: 1000}

	t.Run("create new position from execution", func(t *testing.T) {
		t.Parallel()

		check := func(t *testing.T, instr instruments.Instrument, ex *Execution, pos *Position, acc *Account, tranCnt int) {
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

			cf := ex.cashFlow - ex.commissionConverted
			if notEqual(cf, pos.CashFlow()) {
				t.Errorf(fmtVal, "CashFlow", cf, pos.CashFlow())
			}

			if notEqual(ex.Amount(), pos.Amount()) {
				t.Errorf(fmtVal, "Amount", ex.Amount(), pos.Amount())
			}

			his := pos.AmountHistory()
			if len(his) != 1 {
				t.Errorf(fmtLen, "AmountHistory", 1, len(his))
			}

			if notEqual(ex.Amount(), his[0].Value) {
				t.Errorf(fmtElem, "Value AmountHistory", 0, ex.Amount(), his[0].Value)
			}

			if ex.reportTime != his[0].Time {
				t.Errorf(fmtElem, "Time AmountHistory", 0, ex.reportTime, his[0].Time)
			}

			ehis := pos.ExecutionHistory()
			if len(ehis) != 1 {
				t.Errorf(fmtLen, "ExecutionHistory", 1, len(ehis))
			}

			if *ex != ehis[0] {
				t.Errorf(fmtElem, "ExecutionHistory", 0, *ex, ehis[0])
			}

			tranhis := acc.TransactionHistory()
			if len(tranhis) != tranCnt {
				t.Errorf(fmtLen, "account TransactionHistory", tranCnt, len(tranhis))
			}
		}

		instr := usdWithFactorAndMargin.Instrument()

		t.Run("buy", func(t *testing.T) {
			t.Parallel()

			ex := newExecutionOrderSingle(&mockOrderSingleExecutionReport{
				id:                 "1",
				transactionTime:    time.Now(),
				lastFillPrice:      6,
				lastFillQuantity:   2,
				lastFillCommission: 3,
				commissionCurrency: currencies.USD,
				order:              orders.OrderSingle{Instrument: instr, Side: sides.Buy},
			}, converter)

			account := newAccount("foo", currencies.EUR, converter)
			pos := newPosition(instr, ex, account, matchings.FirstInFirstOut)
			check(t, instr, ex, pos, account, 2)
		})

		t.Run("sell", func(t *testing.T) {
			t.Parallel()

			ex := newExecutionOrderSingle(&mockOrderSingleExecutionReport{
				id:                 "1",
				transactionTime:    time.Now(),
				lastFillPrice:      6,
				lastFillQuantity:   2,
				lastFillCommission: 3,
				commissionCurrency: currencies.USD,
				order:              orders.OrderSingle{Instrument: instr, Side: sides.Sell},
			}, converter)

			account := newAccount("foo", currencies.EUR, converter)
			pos := newPosition(instr, ex, account, matchings.FirstInFirstOut)
			check(t, instr, ex, pos, account, 2)
		})

		t.Run("sell short", func(t *testing.T) {
			t.Parallel()

			ex := newExecutionOrderSingle(&mockOrderSingleExecutionReport{
				id:                 "1",
				transactionTime:    time.Now(),
				lastFillPrice:      6,
				lastFillQuantity:   2,
				lastFillCommission: 3,
				commissionCurrency: currencies.USD,
				order:              orders.OrderSingle{Instrument: instr, Side: sides.SellShort},
			}, converter)

			account := newAccount("foo", currencies.EUR, converter)
			pos := newPosition(instr, ex, account, matchings.FirstInFirstOut)
			check(t, instr, ex, pos, account, 2)
		})

		t.Run("same currency", func(t *testing.T) {
			t.Parallel()

			ex := newExecutionOrderSingle(&mockOrderSingleExecutionReport{
				id:                 "1",
				transactionTime:    time.Now(),
				lastFillPrice:      6,
				lastFillQuantity:   2,
				lastFillCommission: 3,
				commissionCurrency: currencies.USD,
				order:              orders.OrderSingle{Instrument: instr, Side: sides.Buy},
			}, converter)

			account := newAccount("foo", currencies.USD, converter)
			pos := newPosition(instr, ex, account, matchings.FirstInFirstOut)
			check(t, instr, ex, pos, account, 2)
		})

		t.Run("zero commission", func(t *testing.T) {
			t.Parallel()

			ex := newExecutionOrderSingle(&mockOrderSingleExecutionReport{
				id:                 "1",
				transactionTime:    time.Now(),
				lastFillPrice:      6,
				lastFillQuantity:   2,
				lastFillCommission: 0,
				commissionCurrency: currencies.USD,
				order:              orders.OrderSingle{Instrument: instr, Side: sides.Buy},
			}, converter)

			account := newAccount("foo", currencies.EUR, converter)
			pos := newPosition(instr, ex, account, matchings.FirstInFirstOut)
			check(t, instr, ex, pos, account, 1)
		})
	})
	t.Run("update price after creation", func(t *testing.T) {
		t.Parallel()

		instr := usdWithFactorAndMargin.Instrument()
		ex := newExecutionOrderSingle(&mockOrderSingleExecutionReport{
			id:                 "1",
			transactionTime:    time.Now(),
			lastFillPrice:      6,
			lastFillQuantity:   2,
			lastFillCommission: 3,
			commissionCurrency: currencies.USD,
			order:              orders.OrderSingle{Instrument: instr, Side: sides.Buy},
		}, converter)

		account := newAccount("foo", currencies.EUR, converter)
		pos := newPosition(instr, ex, account, matchings.FirstInFirstOut)

		tim := ex.ReportTime().Add(2 * time.Minute)
		pri := ex.Price() + 1
		pos.updatePrice(tim, pri)

		if notEqual(pri, pos.Price()) {
			t.Errorf(fmtVal, "Price", pri, pos.Price())
		}

		amt := pri * pos.priceFactor * pos.quantitySigned
		if notEqual(amt, pos.Amount()) {
			t.Errorf(fmtVal, "Amount", amt, pos.Amount())
		}
	})
	t.Run("add execution", func(t *testing.T) {
		t.Parallel()

		instr := usdWithFactorAndMargin.Instrument()
		ex1 := newExecutionOrderSingle(&mockOrderSingleExecutionReport{
			id:                 "1",
			transactionTime:    time.Now(),
			lastFillPrice:      6,
			lastFillQuantity:   2,
			lastFillCommission: 3,
			commissionCurrency: currencies.USD,
			order:              orders.OrderSingle{Instrument: instr, Side: sides.Buy},
		}, converter)
		ex2 := newExecutionOrderSingle(&mockOrderSingleExecutionReport{
			id:                 "2",
			transactionTime:    ex1.reportTime.Add(time.Hour),
			lastFillPrice:      7,
			lastFillQuantity:   3,
			lastFillCommission: 4,
			commissionCurrency: currencies.USD,
			order:              orders.OrderSingle{Instrument: instr, Side: sides.Buy},
		}, converter)

		account := newAccount("foo", currencies.EUR, converter)
		pos := newPosition(instr, ex1, account, matchings.FirstInFirstOut)
		_ = pos.add(ex2, account)
	})
}
