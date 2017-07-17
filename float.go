package logisticmap

import (
	"math"
	"math/big"
)

// Float compatible with normal float64, or math/big.Float if exact is true.
type Float struct {
	exact bool
	value float64
	float *big.Float
}

// NewFloat allocates and returns a new Float set to float64 value.
func NewFloat(value float64) *Float {
	return &Float{
		exact: false,
		value: value,
	}
}

// NewExactFloat allocates and returns a new Float set to value,
// with precision 53 and rounding mode ToNearestEven.
func NewExactFloat(value float64) *Float {
	return &Float{
		exact: true,
		value: value,
		float: big.NewFloat(value),
	}
}

func newFloat(exact bool, value float64) *Float {
	if exact {
		return NewExactFloat(value)
	}
	return NewFloat(value)
}

// IsExact returns true if uses of math/big.
func (f Float) IsExact() bool {
	return f.exact && f.float != nil
}

// Copy allocates and returns a new same Float.
func (f Float) Copy() *Float {
	if f.IsExact() {
		return NewExactFloat(f.value)
	}
	return NewFloat(f.value)
}

// Add returns f + d2.
func (f Float) Add(d2 *Float) *Float {
	if d2 == nil {
		return f.Copy()
	}
	if f.IsExact() && d2.IsExact() {
		ef := NewExactFloat(0)
		ef.float.Add(f.float, d2.float)
		return ef
	}
	return NewFloat(f.value + d2.value)
}

// Sub returns f - d2.
func (f Float) Sub(d2 *Float) *Float {
	if d2 == nil {
		return f.Copy()
	}
	if f.IsExact() && d2.IsExact() {
		ef := NewExactFloat(0)
		ef.float.Sub(f.float, d2.float)
		return ef
	}
	return NewFloat(f.value - d2.value)
}

// Mul returns f * d2.
func (f Float) Mul(d2 *Float) *Float {
	if d2 == nil {
		return f.Copy()
	}
	if f.IsExact() && d2.IsExact() {
		ef := NewExactFloat(0)
		ef.float.Mul(f.float, d2.float)
		return ef
	}
	return NewFloat(f.value * d2.value)
}

// Quo returns f / d2.
func (f Float) Quo(d2 *Float) *Float {
	if d2 == nil {
		return f.Copy()
	}
	if d2.value == 0 {
		if f.IsExact() && d2.IsExact() {
			return NewExactFloat(math.MaxFloat64)
		}
		return NewFloat(math.MaxFloat64)
	}
	if f.IsExact() && d2.IsExact() {
		ef := NewExactFloat(0)
		ef.float.Quo(f.float, d2.float)
		return ef
	}
	return NewFloat(f.value / d2.value)
}

// Float64 returns the nearest float64 value for f exactly.
func (f Float) Float64() float64 {
	if f.IsExact() {
		f.value, _ = f.float.Float64()
	}
	return f.value
}
