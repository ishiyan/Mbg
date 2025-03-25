package hilberttransformer

import (
	"bytes"
	"fmt"
)

// CycleEstimatorType enumerates types of techniques to estimate
// an instantaneous period using a Hilbert transformer.
type CycleEstimatorType int

const (
	// HomodyneDiscriminator identifies an instantaneous period estimation
	// based on the homodyne discriminator technique.
	HomodyneDiscriminator CycleEstimatorType = iota + 1

	// HomodyneDiscriminatorTaLib identifies an instantaneous period estimation
	// based on the homodyne discriminator technique (TA-Lib implementation with unrolled loops).
	HomodyneDiscriminatorUnrolled

	// PhaseAccumulation identifies an instantaneous period estimation
	// based on the phase accumulation technique.
	PhaseAccumulator

	// DualDifferentiator identifies an instantaneous period estimation
	// based on the dual differentiation technique.
	DualDifferentiator
	last
)

const (
	unknown                       = "unknown"
	homodyneDiscriminator         = "homodyneDiscriminator"
	homodyneDiscriminatorUnrolled = "homodyneDiscriminatorUnrolled"
	phaseAccumulator              = "phaseAccumulator"
	dualDifferentiator            = "dualDifferentiator"
)

// String implements the Stringer interface.
func (t CycleEstimatorType) String() string {
	switch t {
	case HomodyneDiscriminator:
		return homodyneDiscriminator
	case HomodyneDiscriminatorUnrolled:
		return homodyneDiscriminatorUnrolled
	case PhaseAccumulator:
		return phaseAccumulator
	case DualDifferentiator:
		return dualDifferentiator
	default:
		return unknown
	}
}

// IsKnown determines if this cycle estimator type is known.
func (t CycleEstimatorType) IsKnown() bool {
	return t >= HomodyneDiscriminator && t < last
}

// MarshalJSON implements the Marshaler interface.
func (t CycleEstimatorType) MarshalJSON() ([]byte, error) {
	const (
		errFmt = "cannot marshal '%s': unknown cycle estimator type"
		extra  = 2   // Two bytes for quotes.
		dqc    = '"' // Double quote character.
	)

	s := t.String()
	if s == unknown {
		return nil, fmt.Errorf(errFmt, s)
	}

	b := make([]byte, 0, len(s)+extra)
	b = append(b, dqc)
	b = append(b, s...)
	b = append(b, dqc)

	return b, nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (t *CycleEstimatorType) UnmarshalJSON(data []byte) error {
	const (
		errFmt = "cannot unmarshal '%s': unknown cycle estimator type"
		dqs    = "\"" // Double quote string.
	)

	d := bytes.Trim(data, dqs)
	s := string(d)

	switch s {
	case homodyneDiscriminator:
		*t = HomodyneDiscriminator
	case homodyneDiscriminatorUnrolled:
		*t = HomodyneDiscriminatorUnrolled
	case phaseAccumulator:
		*t = PhaseAccumulator
	case dualDifferentiator:
		*t = DualDifferentiator
	default:
		return fmt.Errorf(errFmt, s)
	}

	return nil
}
