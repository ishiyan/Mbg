package orders

//nolint:gofumpt
import (
	"time"

	"mbg/trading/currencies"
	"mbg/trading/orders/reports"
	"mbg/trading/orders/status"
)

// OrderSingleExecutionReport is a report event for an order in a single instrument.
type OrderSingleExecutionReport interface {
	// Order is the underlying order for this execution report.
	// If there were any successful order replacements, this will be the most recent version.
	Order() OrderSingle

	// TransactionTime is the date and time when the business
	// represented by this report occurred.
	TransactionTime() time.Time

	// Status is the current state of an order as understood by the broker.
	Status() status.OrderStatus

	// ReportType identifies an action of this report.
	ReportType() reports.OrderReportType

	// ID is a unique identifier of this report as assigned by the sell-side.
	ID() string

	// Note is a free-format text that accompany this report.
	Note() string

	// ReplaceSourceOrder is the replace source order.
	// Filled when report type is Replaced or ReplaceRejected.
	ReplaceSourceOrder() OrderSingle

	// ReplaceTargetOrder is the replace target order.
	// Filled when report type is Replaced or ReplaceRejected.
	ReplaceTargetOrder() OrderSingle

	// LastFillPrice is the price (in order instrument's currency) of the last fill.
	LastFillPrice() float64

	// AveragePrice is an average price (in order instrument's currency) of all fills.
	AveragePrice() float64

	// LastFillQuantity is the quantity bought or sold on the last fill.
	LastFillQuantity() float64

	// LeavesQuantity is the quantity open for further execution.
	//
	// If the order status is Canceled, Expired or Rejected (in which case
	// the order is no longer active) then this could be 0, otherwise
	//   Order.Quantity - CumulativeQuantity.
	LeavesQuantity() float64

	// CumulativeQuantity is the total quantity filled.
	CumulativeQuantity() float64

	// LastFillCommission is the commission (in commission currency) of the last fill.
	LastFillCommission() float64

	// CumulativeCommission is the total commission (in commission currency) for all fills.
	CumulativeCommission() float64

	// CommissionCurrency is a commission currency.
	CommissionCurrency() currencies.Currency
}
