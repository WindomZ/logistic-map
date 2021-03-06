package logisticmap

// ILogistic interface of logistic map algorithm.
type ILogistic interface {
	Next() ILogistic
	Value() float64
}

// Logistic A logistic map algorithm node instance.
type Logistic struct {
	exact bool
	n     int64
	u     *Float
	x     *Float
	one   *Float
}

// NewLogistic allocates and returns a new Logistic.
func NewLogistic(exact bool, u float64, x float64) *Logistic {
	return &Logistic{
		exact: true,
		u:     newFloat(exact, u),
		x:     newFloat(exact, x),
		one:   NewExactFloat(1),
	}
}

// Next calculate the next x through the mathematical formula:
// u * x * (1 - x)
func (l *Logistic) Next() ILogistic {
	l.n++
	l.x = l.u.Mul(l.x.Mul(l.one.Sub(l.x)))
	return l
}

// Value returns the float64 result of Float.
func (l Logistic) Value() float64 {
	return l.x.Float64()
}

// LogisticMap returns (n+1) size of float64 array,
// with logistic map algorithm parameter u and x, calculates n iterates.
func LogisticMap(u float64, x float64, n int64) []float64 {
	l := NewLogistic(true, u, x)
	result := make([]float64, n+1)
	result[0] = x

	var f float64
	for l.n < n {
		f = l.Next().Value()
		result[l.n] = f
	}

	return result
}
