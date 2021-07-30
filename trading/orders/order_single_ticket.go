package orders

import (
	"mbg/trading/orders/status"
)

// OrderSingleTicket tracks a new order in a single instrument.
type OrderSingleTicket interface {
	// Order is the underlying order for this ticket. If there were any
	// successful order replacements, this will be the most recent version.
	Order() OrderSingle

	// ClientOrderID is a unique identifier for an order as assigned by the
	// buy-side (institution, broker, intermediary etc.).
	ClientOrderID() string

	// OrderID is a unique identifier for an order as assigned by the sell-side.
	OrderID() string

	// Status is the current state of an order as understood by the broker.
	Status() status.OrderStatus

	// LastReport is the last order report, nil if not any.
	LastReport() OrderSingleExecutionReport

	// Reports provides a collection of all order reports in the chronological order.
	Reports() []OrderSingleExecutionReport

	// CancelReplace is used to change the parameters of an existing order.
	//
	// Allowed modifications to an order include:
	//  - reducing or increasing order quantity
	//  - changing a limit order to a market order
	//  - changing the limit price
	//  - changing time in force
	//
	// Modifications cannot include:
	//  - changing order side
	//  - changing series
	//  - reducing quantity to zero (canceling the order)
	//  - re-opening a filled order by increasing quantity
	//
	// Unchanging attributes to be carried over from the original order
	// must be specified in the replacement.
	//
	// If the order has been completed (successfully or not), does nothing.
	//
	// Produces an execution report on completion.
	CancelReplace(replacementOrder OrderSingle)

	// Cancel cancels this order.
	//
	// If order has been already completed (successfully or not), does nothing.
	//
	// Produces an execution report on completion.
	Cancel()
}
