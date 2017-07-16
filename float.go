package logisticmap

import (
	"math"
	"math/big"
)

var (
	ONE       = NewFloat(1)
	ONE_EXACT = NewExactFloat(1)
)

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

// NewFloat allocates and returns a new Float set to value,
// with precision 53 and rounding mode ToNearestEven.
func NewExactFloat(value float64) *Float {
	return &Float{
		exact: true,
		value: value,
		float: big.NewFloat(value),
	}
}

func FloatOne(exact bool) *Float {
	if exact {
		return ONE_EXACT
	}
	return ONE
}

func (f Float) IsExact() bool {
	return f.exact && f.float != nil
}

func (f Float) Copy() *Float {
	if f.exact {
		return NewExactFloat(f.value)
	}
	return NewFloat(f.value)
}

// Add returns f + d2.
func (f Float) Add(d2 *Float) *Float {
	if d2 == nil {
		return f.Copy()
	}
	if f.exact && d2.exact {
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
	if f.exact && d2.exact {
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
	if f.exact && d2.exact {
		ef := NewExactFloat(0)
		ef.float.Mul(f.float, d2.float)
		return ef
	}
	return NewFloat(f.value * d2.value)
}

// Quo returns f / d2.
func (f Float) Quo(d2 *Float) *Float {
	if d2.value == 0 {
		if f.exact && d2.exact {
			return NewExactFloat(math.MaxFloat64)
		}
		return NewFloat(math.MaxFloat64)
	}
	if d2 == nil {
		return f.Copy()
	}
	if f.exact && d2.exact {
		ef := NewExactFloat(0)
		ef.float.Quo(f.float, d2.float)
		return ef
	}
	return NewFloat(f.value / d2.value)
}

// Float64 returns the nearest float64 value for f and a bool if exactly.
func (f Float) Float64() (float64, bool) {
	if f.IsExact() {
		f.value, _ = f.float.Float64()
	}
	return f.value, f.exact
}