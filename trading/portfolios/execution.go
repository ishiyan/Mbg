package portfolios

//nolint:gci
import (
	"math"
	"mbg/trading/currencies"
	"mbg/trading/orders"
	"mbg/trading/orders/sides"
	"time"
)

// Execution contains properties of a fill or a partial fill of an order.
type Execution struct {
	reportID            string
	reportTime          time.Time
	side                sides.Side
	quantity            float64
	currency            currencies.Currency
	commissionCurrency  currencies.Currency
	conversionRate      float64
	commission          float64
	commissionConverted float64
	price               float64
	amount              float64
	margin              float64
	debt                float64
	pnl                 float64
	realizedPnL         float64
	netCashFlow         float64
	cashFlow            float64
	roundtripQuantity   float64
	roundtripPriceHigh  float64
	roundtripPriceLow   float64
}

// newExecutionOrderSingle creates an execution from a Filled or PartiallyFilled
// execution report of an order in a single instrument.
func newExecutionOrderSingle(report orders.OrderSingleExecutionReport, converter currencies.Converter) *Execution {
	qtyAbs := math.Abs(report.LastFillQuantity())
	qtySigned := qtyAbs

	side := report.Order().Side
	if side.IsSell() {
		qtySigned = -qtyAbs
	}

	price := report.LastFillPrice()
	priceFactored := price
	instrument := report.Order().Instrument

	if instrument.PriceFactor != 0 {
		priceFactored *= instrument.PriceFactor
	}

	margin := instrument.Margin * qtyAbs
	amount := priceFactored * qtyAbs
	netCashFlow := -qtySigned * priceFactored

	debt := amount - margin
	if margin == 0 {
		debt = 0
	}

	rate := 1.0
	conv := report.LastFillCommission()

	if report.CommissionCurrency() != instrument.Currency {
		conv, rate = converter.Convert(conv, report.CommissionCurrency(), instrument.Currency)
	}

	return &Execution{
		reportID:            report.ID(),
		reportTime:          report.TransactionTime(),
		side:                side,
		quantity:            qtyAbs,
		currency:            instrument.Currency,
		commissionCurrency:  report.CommissionCurrency(),
		conversionRate:      rate,
		commission:          report.LastFillCommission(),
		commissionConverted: conv,
		price:               price,
		amount:              amount,
		margin:              margin,
		debt:                debt,
		pnl:                 -conv, // Will be updated when added to position.
		realizedPnL:         0,     // Will be updated when added to position.
		netCashFlow:         netCashFlow,
		cashFlow:            netCashFlow - conv,
		roundtripQuantity:   qtyAbs,
		roundtripPriceHigh:  price,
		roundtripPriceLow:   price,
	}
}

func (e *Execution) quantitySign() float64 {
	if e.side.IsSell() {
		return -1
	}

	return 1
}

// ReportID is a unique transaction identifier of an associated
// execution report as assigned by the sell-side.
func (e *Execution) ReportID() string {
	return e.reportID
}

// ReportTime is the date and time of an associated
// execution report as assigned by the broker.
func (e *Execution) ReportTime() time.Time {
	return e.reportTime
}

// Side is the execution order side,
// which determines the sign of the quantity.
func (e *Execution) Side() sides.Side {
	return e.side
}

// Quantity is the unsigned execution quantity,
// the sign is determined by the side.
func (e *Execution) Quantity() float64 {
	return e.quantity
}

// Currency is the instrument's currency.
func (e *Execution) Currency() currencies.Currency {
	return e.currency
}

// CommissionCurrency is a commission currency.
func (e *Execution) CommissionCurrency() currencies.Currency {
	return e.commissionCurrency
}

// ConversionRate is an exchange rate
// from the commission currency to the instrument's currency.
func (e *Execution) ConversionRate() float64 {
	return e.conversionRate
}

// Commission is the execution commission amount
// in the commission currency.
func (e *Execution) Commission() float64 {
	return e.commission
}

// CommissionConverted is the execution commission amount
// in instrument's currency.
func (e *Execution) CommissionConverted() float64 {
	return e.commissionConverted
}

// Price is the execution price in instrument's currency.
func (e *Execution) Price() float64 {
	return e.price
}

// Amount is the unsigned execution value in instrument's currency
// (factored price times quantity).
func (e *Execution) Amount() float64 {
	return e.amount
}

// Margin is the execution margin in instrument's currency
// (instrument margin times quantity).
func (e *Execution) Margin() float64 {
	return e.margin
}

// Debt is the execution debt in instrument's currency
// (amount minus margin).
func (e *Execution) Debt() float64 {
	return e.debt
}

// PnL is the Profit and Loss of this execution in instrument's currency.
func (e *Execution) PnL() float64 {
	return e.pnl
}

// RealizedPnL is the realized Profit and Loss of this execution
// in instrument's currency.
func (e *Execution) RealizedPnL() float64 {
	return e.realizedPnL
}

// NetCashFlow is the execution net cash flow in instrument's currency
// (factored price times negative signed quantity).
func (e *Execution) NetCashFlow() float64 {
	return e.netCashFlow
}

// CashFlow is the execution cash flow in instrument's currency
// (net cash flow minus commission).
func (e *Execution) CashFlow() float64 {
	return e.cashFlow
}
