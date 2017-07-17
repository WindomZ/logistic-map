package logisticmap

import (
	"math"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestNewLogistic(t *testing.T) {
	const u = 3.9999
	const x = 0.2
	const COUNT = 30000

	l := NewLogistic(true, u, x)

	result := make([]float64, COUNT+1)
	result[0] = x
	for i := 1; i <= COUNT; i++ {
		result[i] = l.Next().Value()
	}
	avgs := make([]int, 10)
	for _, r := range result {
		avgs[int(math.Floor(r*10.0))]++
	}

	assert.Equal(t, avgs, []int{5901, 2752, 2263, 2074, 1895, 1941, 2044, 2163, 2725, 6243})
}

func TestLogisticMap(t *testing.T) {
	const u = 3.9999
	const x = 0.2
	const COUNT = 30000

	result := LogisticMap(u, x, COUNT)

	avgs := make([]int, 10)
	for _, r := range result {
		avgs[int(math.Floor(r*10.0))]++
	}

	assert.Equal(t, avgs, []int{5901, 2752, 2263, 2074, 1895, 1941, 2044, 2163, 2725, 6243})
}

func TestLogistic_Next(t *testing.T) {
	const u = 3.9999
	const x = 0.2
	const COUNT = 30000

	l := NewLogistic(false, u, x)
	result1 := make([]float64, COUNT+1)
	result1[0] = x
	for i := 1; i <= COUNT; i++ {
		result1[i] = l.Next().Value()
	}

	result2 := LogisticMap(u, x, COUNT)

	assert.Equal(t, result1, result2)
}
