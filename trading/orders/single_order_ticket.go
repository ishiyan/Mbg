package orders

import (
	"mbg/trading/orders/status"
)

// SingleOrderTicket is a ticket for an order in a single instrument.
type SingleOrderTicket struct {

	// Notifies when a report has been received.
	// event Action<ISingleOrderTicket, SingleOrderReport> OrderReport;

	// Notifies when an order has been completed.
	// This is called after the order has been moved to the terminal state.
	// event Action<ISingleOrderTicket> OrderCompleted;

	// Order is the underlying order for this ticket. If there were any
	// successful order replacements, this will be the most recent version.
	Order SingleOrder

	// ClientOrderID is a unique identifier for an order as assigned by the
	// buy-side (institution, broker, intermediary etc.).
	ClientOrderID string

	// OrderID is a unique identifier for an order as assigned by the sell-side.
	OrderID string

	// Status is the current state of an order as understood by the broker.
	Status status.OrderStatus

	// LastFillPrice is the price of the last fill.
	// Zero if not set.
	LastFillPrice float64

	// LastFillQuantity is the quantity bought or sold on the last fill.
	// Zero if not set.
	LastFillQuantity float64

	// LeavesQuantity is the quantity open for further execution.
	// If the order status is Canceled, Expired or Rejected (in which case
	// the order is no longer active) then this could be 0, otherwise
	//  Order.Quantity - CumulativeQuantity.
	LeavesQuantity float64

	// CumulativeQuantity is the total quantity filled.
	CumulativeQuantity float64

	// AveragePrice is an average price (in instrument's currency) of all fills.
	AveragePrice float64

	// LastFillCommission is the commission (in commission currency) of the last fill.
	LastFillCommission float64

	// CumulativeCommission is the total commission (in commission currency) for all fills.
	CumulativeCommission float64

	// LastReport is the last order report, nil if not any.
	LastReport SingleOrderReport

	// Reports provides a collection of all order reports in the chronological order.
	Reports []SingleOrderReport

	// Replace replaces a pending order. If the order has been completed
	// (successfully or not), does nothing.
	// func Replace(replacementOrder SingleOrder)

	// Cancel cancels this order. If order has been already completed
	// (successfully or not), does nothing.
	// func Cancel()
}
