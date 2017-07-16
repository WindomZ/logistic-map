package logisticmap

type Logistic struct {
	exact bool
	u     *Float
	x     *Float
}

func NewLogistic(exact bool, u float64, x float64) *Logistic {
	if exact {
		return &Logistic{
			exact: true,
			u:     NewExactFloat(u),
			x:     NewExactFloat(x),
		}
	}
	return &Logistic{
		u: NewFloat(u),
		x: NewFloat(x),
	}
}

func (l *Logistic) Next() *Logistic {
	l.x = l.u.Mul(l.x.Mul(FloatOne(l.exact).Sub(l.x)))
	return l
}

func (l Logistic) Value() float64 {
	f, _ := l.x.Float64()
	return f
}

func LogisticMap(u float64, x float64, n int) []float64 {
	l := NewLogistic(true, u, x)
	result := make([]float64, n+1)
	result[0] = x

	for i := 1; i <= n; i++ {
		result[i] = l.Next().Value()
	}
	return result
}
