package entities

import "time"

// Trade (also called "time and sales") represents [price, volume] pairs.
type Trade struct {
	Time   time.Time `json:"t"` // The date and time.
	Price  float64   `json:"p"` // The price.
	Volume float64   `json:"v"` // The volume.
}
