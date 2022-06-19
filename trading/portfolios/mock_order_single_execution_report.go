package portfolios

//nolint:gofumpt
import (
	"time"

	"mbg/trading/currencies"
	"mbg/trading/orders"
	"mbg/trading/orders/reports"
	"mbg/trading/orders/status"
)

// mockOrderSingleExecutionReport is a mock report event for an order in a single instrument.
type mockOrderSingleExecutionReport struct {
	order                orders.OrderSingle
	transactionTime      time.Time
	status               status.OrderStatus
	reportType           reports.OrderReportType
	id                   string
	note                 string
	replaceSourceOrder   orders.OrderSingle
	replaceTargetOrder   orders.OrderSingle
	lastFillPrice        float64
	averagePrice         float64
	lastFillQuantity     float64
	leavesQuantity       float64
	cumulativeQuantity   float64
	lastFillCommission   float64
	cumulativeCommission float64
	commissionCurrency   currencies.Currency
}

// Order is the underlying order for this execution report.
func (m *mockOrderSingleExecutionReport) Order() orders.OrderSingle {
	return m.order
}

// TransactionTime is the date and time when the business represented by this report occurred.
func (m *mockOrderSingleExecutionReport) TransactionTime() time.Time {
	return m.transactionTime
}

// Status is the current state of an order as understood by the broker.
func (m *mockOrderSingleExecutionReport) Status() status.OrderStatus {
	return m.status
}

// ReportType identifies an action of this report.
func (m *mockOrderSingleExecutionReport) ReportType() reports.OrderReportType {
	return m.reportType
}

// ID is a unique identifier of this report as assigned by the sell-side.
func (m *mockOrderSingleExecutionReport) ID() string {
	return m.id
}

// Note is a free-format text that accompany this report.
func (m *mockOrderSingleExecutionReport) Note() string {
	return m.note
}

// ReplaceSourceOrder is the replace source order.
// Filled when report type is Replaced or ReplaceRejected.
func (m *mockOrderSingleExecutionReport) ReplaceSourceOrder() orders.OrderSingle {
	return m.replaceSourceOrder
}

// ReplaceTargetOrder is the replace target order.
// Filled when report type is Replaced or ReplaceRejected.
func (m *mockOrderSingleExecutionReport) ReplaceTargetOrder() orders.OrderSingle {
	return m.replaceTargetOrder
}

// LastFillPrice is the price (in order instrument's currency) of the last fill.
func (m *mockOrderSingleExecutionReport) LastFillPrice() float64 {
	return m.lastFillPrice
}

// AveragePrice is an average price (in order instrument's currency) of all fills.
func (m *mockOrderSingleExecutionReport) AveragePrice() float64 {
	return m.averagePrice
}

// LastFillQuantity is the quantity bought or sold on the last fill.
func (m *mockOrderSingleExecutionReport) LastFillQuantity() float64 {
	return m.lastFillQuantity
}

// LeavesQuantity is the quantity open for further execution.
//
// If the order status is Canceled, Expired or Rejected (in which case
// the order is no longer active) then this could be 0, otherwise
//   Order.Quantity - CumulativeQuantity.
func (m *mockOrderSingleExecutionReport) LeavesQuantity() float64 {
	return m.leavesQuantity
}

// CumulativeQuantity is the total quantity filled.
func (m *mockOrderSingleExecutionReport) CumulativeQuantity() float64 {
	return m.cumulativeQuantity
}

// LastFillCommission is the commission (in commission currency) of the last fill.
func (m *mockOrderSingleExecutionReport) LastFillCommission() float64 {
	return m.lastFillCommission
}

// CumulativeCommission is the total commission (in commission currency) for all fills.
func (m *mockOrderSingleExecutionReport) CumulativeCommission() float64 {
	return m.cumulativeCommission
}

// CommissionCurrency is a commission currency.
func (m *mockOrderSingleExecutionReport) CommissionCurrency() currencies.Currency {
	return m.commissionCurrency
}
