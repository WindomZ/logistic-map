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

	assert.Equal(t, avgs, []int{1291, 1385, 1669, 2013, 2384, 3442, 4106, 5673, 4627, 3411})
}

func TestLogisticImprovedMap(t *testing.T) {
	const x = 0.2
	const COUNT = 30000

	result := LogisticImprovedMap(x, 100, COUNT)

	avgs := make([]int, 10)
	for _, r := range result {
		avgs[int(math.Floor(r*10.0))]++
	}

	assert.Equal(t, avgs, []int{964, 1489, 1641, 1318, 1887, 4111, 4653, 6600, 4073, 3265})
}
