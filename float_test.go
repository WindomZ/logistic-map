package logisticmap

import (
	"math"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestNewFloat(t *testing.T) {
	f0 := newFloat(false, 0)
	f1 := newFloat(false, 1)
	f2 := newFloat(false, 2)

	assert.Equal(t, f1.Add(f2).Float64(), float64(3))
	assert.Equal(t, f1.Sub(f2).Float64(), float64(-1))
	assert.Equal(t, f1.Mul(f2).Float64(), float64(2))
	assert.Equal(t, f1.Quo(f2).Float64(), float64(0.5))

	assert.Equal(t, f1.Quo(f0).Float64(), math.MaxFloat64)

	assert.Equal(t, f1.Add(nil).Float64(), float64(1))
	assert.Equal(t, f1.Sub(nil).Float64(), float64(1))
	assert.Equal(t, f1.Mul(nil).Float64(), float64(1))
	assert.Equal(t, f1.Quo(nil).Float64(), float64(1))
}

func TestNewExactFloat(t *testing.T) {
	f0 := newFloat(true, 0)
	f1 := newFloat(true, 1)
	f2 := newFloat(true, 2)

	assert.Equal(t, f1.Add(f2).Float64(), float64(3))
	assert.Equal(t, f1.Sub(f2).Float64(), float64(-1))
	assert.Equal(t, f1.Mul(f2).Float64(), float64(2))
	assert.Equal(t, f1.Quo(f2).Float64(), float64(0.5))

	assert.Equal(t, f1.Quo(f0).Float64(), math.MaxFloat64)

	assert.Equal(t, f1.Add(nil).Float64(), float64(1))
	assert.Equal(t, f1.Sub(nil).Float64(), float64(1))
	assert.Equal(t, f1.Mul(nil).Float64(), float64(1))
	assert.Equal(t, f1.Quo(nil).Float64(), float64(1))
}
