package indicator

import (
	"mbg/trading/indicators/indicator/output"
)

// Metadata describes a type and outputs of an indicator.
type Metadata struct {
	// Type identifies a type this indicator.
	Type Type `json:"type"`

	// Outputs is a slice of metadata for individual outputs.
	Outputs []output.Metadata `json:"outputs"`
}
