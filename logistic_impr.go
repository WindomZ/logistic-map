package logisticmap

import "math"

// LogisticImproved A improved logistic map algorithm node instance.
type LogisticImproved struct {
	Logistic
	g        *Float
	four     *Float
	ten2Neg6 *Float
	hPi      *Float
}

// NewLogisticImproved allocates and returns a new LogisticImproved.
func NewLogisticImproved(exact bool, x float64, g int) *LogisticImproved {
	return &LogisticImproved{
		Logistic: *NewLogistic(exact, 3.569945673, x),
		g:        NewExactFloat(float64(g)),
		four:     NewExactFloat(4),
		ten2Neg6: NewExactFloat(0.000001),
		hPi:      NewExactFloat(math.Pi / 2),
	}
}

// Next calculate the next x through the mathematical formula:
// [u + (4 - u) * cos((0.000001 + x) * pi / 2)] * x * (1 - x / g)
func (l *LogisticImproved) Next() ILogistic {
	l.n++
	u := l.u.Add(l.four.Sub(l.u).Mul(newFloat(l.exact, math.Cos(l.ten2Neg6.Add(l.x).Mul(l.hPi).Float64()))))
	l.x = u.Mul(l.x.Mul(l.one.Sub(l.x.Quo(l.g))))
	return l
}

// Value returns the float64 result of Float.
func (l LogisticImproved) Value() float64 {
	return l.x.Quo(l.g).Float64()
}

// LogisticImprovedMap returns (n+1) size of float64 array,
// with improved logistic map algorithm parameter x, calculates n iterates.
// g is a amplification factor that default minimum value is 1.
func LogisticImprovedMap(x float64, g int, n int64) []float64 {
	if g <= 0 {
		g = 1
	}
	l := NewLogisticImproved(true, x, g)
	result := make([]float64, n+1)
	result[0] = x

	var f float64
	for l.n < n {
		f = l.Next().Value()
		result[l.n] = f
	}

	return result
}
