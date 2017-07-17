package logisticmap

import "math"

type LogisticImproved struct {
	Logistic
	g        *Float
	four     *Float
	ten2Neg6 *Float
	hPi      *Float
}

func NewLogisticImproved(exact bool, x float64, g int) *LogisticImproved {
	return &LogisticImproved{
		Logistic: *NewLogistic(exact, 3.569945673, x),
		g:        NewExactFloat(float64(g)),
		four:     NewExactFloat(4),
		ten2Neg6: NewExactFloat(0.000001),
		hPi:      NewExactFloat(math.Pi / 2),
	}
}

// Next the mathematical formula: [u + (4 - u) * cos((0.000001 + x) * pi / 2)] * x * (1 - x / g)
func (l *LogisticImproved) Next() ILogistic {
	l.n++
	u := l.u.Add(l.four.Sub(l.u).Mul(newFloat(l.exact, math.Cos(l.ten2Neg6.Add(l.x).Mul(l.hPi).Float64()))))
	l.x = u.Mul(l.x.Mul(l.one.Sub(l.x.Quo(l.g))))
	return l
}

func (l LogisticImproved) Value() float64 {
	return l.x.Quo(l.g).Float64()
}

func LogisticImprovedMap(x float64, g int, n int64) []float64 {
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
