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
	reportID                   string
	reportTime                 time.Time
	side                       sides.Side
	quantity                   float64
	quantitySign               float64
	currency                   currencies.Currency
	commissionCurrency         currencies.Currency
	conversionRate             float64
	commission                 float64
	commissionConverted        float64
	commissionConvertedPerUnit float64
	price                      float64
	amount                     float64
	margin                     float64
	debt                       float64
	pnl                        float64
	realizedPnL                float64
	cashFlow                   float64
	unrealizedQuantity         float64
	unrealizedPriceHigh        float64
	unrealizedPriceLow         float64
}

// newExecutionOrderSingle creates an execution from a Filled or PartiallyFilled
// execution report of an order in a single instrument.
func newExecutionOrderSingle(report orders.OrderSingleExecutionReport, converter currencies.Converter) *Execution {
	qtyAbs := math.Abs(report.LastFillQuantity())
	side := report.Order().Side

	qtySign := 1.
	if side.IsSell() {
		qtySign = -1. //nolint:gomnd
	}

	price := report.LastFillPrice()
	priceFactored := price
	instrument := report.Order().Instrument

	if instrument.PriceFactor() != 0 {
		priceFactored *= instrument.PriceFactor()
	}

	marginAbs := instrument.Margin() * qtyAbs
	amountAbs := priceFactored * qtyAbs
	cashFlow := -qtySign * amountAbs

	debt := amountAbs - marginAbs
	if marginAbs == 0 {
		debt = 0
	}

	rate := 1.0
	conv := report.LastFillCommission()

	if report.CommissionCurrency() != instrument.Currency() {
		conv, rate = converter.Convert(conv, report.CommissionCurrency(), instrument.Currency())
	}

	return &Execution{
		reportID:                   report.ID(),
		reportTime:                 report.TransactionTime(),
		side:                       side,
		quantity:                   qtyAbs,
		quantitySign:               qtySign,
		currency:                   instrument.Currency(),
		commissionCurrency:         report.CommissionCurrency(),
		conversionRate:             rate,
		commission:                 report.LastFillCommission(),
		commissionConverted:        conv,
		commissionConvertedPerUnit: conv / qtyAbs,
		price:                      price,
		amount:                     amountAbs,
		margin:                     marginAbs,
		debt:                       debt,
		pnl:                        -conv, // Will be updated when adding to position.
		realizedPnL:                0,     // Will be updated when adding to position.
		cashFlow:                   cashFlow,
		unrealizedQuantity:         qtyAbs,
		unrealizedPriceHigh:        price,
		unrealizedPriceLow:         price,
	}
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

// Side is the execution order side, which determines the sign of the quantity.
func (e *Execution) Side() sides.Side {
	return e.side
}

// Quantity is the unsigned execution quantity, the sign is determined by the side.
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

// Commission is the execution commission amount in the commission currency.
func (e *Execution) Commission() float64 {
	return e.commission
}

// CommissionConverted is the execution commission amount in the instrument's currency.
func (e *Execution) CommissionConverted() float64 {
	return e.commissionConverted
}

// Price is the execution price in the instrument's currency.
func (e *Execution) Price() float64 {
	return e.price
}

// Amount is the unsigned execution value in the instrument's currency
// (factored price times quantity).
func (e *Execution) Amount() float64 {
	return e.amount
}

// Margin is the unsigned execution margin in instrument's currency
// (instrument margin times quantity).
func (e *Execution) Margin() float64 {
	return e.margin
}

// Debt is the execution debt in the instrument's currency
// (amount minus margin).
func (e *Execution) Debt() float64 {
	return e.debt
}

// PnL is the Profit and Loss in the instrument's currency.
func (e *Execution) PnL() float64 {
	return e.pnl
}

// RealizedPnL is the realized Profit and Loss in the instrument's currency.
func (e *Execution) RealizedPnL() float64 {
	return e.realizedPnL
}

// CashFlow is the execution cash flow in the instrument's currency
// (factored price times negative signed quantity).
func (e *Execution) CashFlow() float64 {
	return e.cashFlow
}
