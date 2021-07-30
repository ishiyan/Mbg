// Package orders defines order properties and functionality to create and track orders.
package orders

//nolint:gci
import (
	"mbg/trading/instruments"
	"mbg/trading/orders/sides"
	"mbg/trading/orders/tif"
	"mbg/trading/orders/types"
	"time"
)

// OrderSingle is a request to place an order in a single instrument.
type OrderSingle struct {
	// Instrument specifies a single instrument associated with this order.
	Instrument instruments.Instrument

	// Type specifies an order type associated with this order.
	Type types.OrderType

	// Side specifies an order side associated with this order.
	Side sides.Side

	// TimeInForce specifies a time in force associated with this order.
	TimeInForce tif.OrderTimeInForce

	// Quantity is a total order quantity (in units) to execute.
	// Zero if not set.
	Quantity float64

	// MinimumQuantity is a minimum quantity (in units) of an order to be executed.
	// Zero if not set.
	MinimumQuantity float64

	// LimitPrice is the limit price in instrument's currency per unit of quantity.
	// Zero if not set.
	//
	// Required for limit order types. For FX orders, should be the "all-in"
	// rate (spot rate adjusted for forward points). Can be used to specify
	// a limit price for a pegged order, previously indicated, etc.
	LimitPrice float64

	// StopPrice is the stop price in instrument's currency per unit of quantity.
	// Zero if not set.
	StopPrice float64

	// TrailingDistance is the trailing distance. Zero if not set.
	TrailingDistance float64

	// CreationTime is the date and time this order request was created
	// by a trader, trading system, or intermediary.
	CreationTime time.Time

	// ExpirationTime is the order expiration date and time for the orders
	// with the GoodTillDate TimeInForce value.
	ExpirationTime time.Time

	// Note is a free-format text with notes on the order.
	Note string
}
