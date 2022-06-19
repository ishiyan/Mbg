// Package statistics implements various statistical indicators.
package statistics

import "math"

const (
	unknown         = "unknown"
	dqs             = "\""
	dqc             = '"'
	marshalErrFmt   = "cannot marshal '%s': %w"
	unmarshalErrFmt = "cannot unmarshal '%s': %w"
)

var nan = math.NaN()
