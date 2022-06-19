package output

// Metadata describes a single indicator output.
type Metadata struct {
	// Kind is an identification of this indicator output.
	// It is an integer representation of an output enumeration of a related indicator.
	Kind int `json:"kind"`

	// Type describes a data type of this indicator output.
	Type Type `json:"type"`

	// Name is a name of this indicator output.
	Name string `json:"name"`

	// Description is a description of this indicator output.
	Description string `json:"description"`
}
