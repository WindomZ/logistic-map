package logisticmap

import (
	"math"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestNewLogisticImproved(t *testing.T) {
	const x = 0.2
	const COUNT = 30000

	l := NewLogisticImproved(true, x, COUNT)

	result := make([]float64, COUNT+1)
	result[0] = x
	for i := 1; i <= COUNT; i++ {
		result[i] = l.Next().Value()
	}
	avgs := make([]int, 10)
	for _, r := range result {
		avgs[int(math.Floor(r*10.0))]++
	}

	assert.Equal(t, []int{1291, 1385, 1669, 2013, 2384, 3442, 4106, 5673, 4627, 3411}, avgs)
}

func TestLogisticImprovedMap(t *testing.T) {
	const x = 0.2
	const COUNT = 30000

	result := LogisticImprovedMap(x, 1000, COUNT)

	avgs := make([]int, 10)
	for _, r := range result {
		avgs[int(math.Floor(r*10.0))]++
	}

	assert.Equal(t, []int{1333, 1471, 1654, 1884, 2401, 3526, 4126, 5655, 4519, 3432}, avgs)
}

func TestLogisticImprovedMap2(t *testing.T) {
	const x = 0.2
	const COUNT = 30000

	result := LogisticImprovedMap(x, 0, COUNT)

	avgs := make([]int, 10)
	for _, r := range result {
		avgs[int(math.Floor(r*10.0))]++
	}

	assert.Equal(t, []int{0, 4349, 2336, 2161, 2783, 2213, 2667, 3586, 2907, 6999}, avgs)
}
