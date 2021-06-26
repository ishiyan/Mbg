package orders

//nolint:gci
import (
	"mbg/trading/currencies"
	"mbg/trading/orders/reports"
	"mbg/trading/orders/status"
	"time"
)

// SingleOrderReport is a report event for an order in a single instrument.
type SingleOrderReport struct {
	// TransactionTime is the date and time when the business
	// represented by this report occurred.
	TransactionTime time.Time

	// Status is the current state of an order as understood by the broker.
	Status status.OrderStatus

	// Type identifies an action of this report.
	Type reports.OrderReportType

	// ID is a unique identifier of this report as assigned by the sell-side.
	ID string

	// Note is a free-format text that accompany this report.
	Note string

	// ReplaceSourceOrder is the replace source order.
	// Filled when report type is Replaced or ReplaceRejected.
	ReplaceSourceOrder SingleOrder

	// ReplaceTargetOrder is the replace target order.
	// Filled when report type is Replaced or ReplaceRejected.
	ReplaceTargetOrder SingleOrder

	// LastPrice is the price of the last fill.
	// Zero if not set.
	LastPrice float64

	// LastQuantity is the quantity bought or sold on the last fill.
	// Zero if not set.
	LastQuantity float64

	// LeavesQuantity is the quantity open for further execution.
	LeavesQuantity float64

	// CumulativeQuantity is the total quantity filled.
	CumulativeQuantity float64

	// AveragePrice is an average price (in instrument's currency) of all fills.
	AveragePrice float64

	// LastCommission is the commission (in commission currency) of the last fill.
	LastCommission float64

	// CumulativeCommission is the total commission (in commission currency) for all fills.
	CumulativeCommission float64

	// CommissionCurrency is a commission currency.
	CommissionCurrency currencies.Currency
}
