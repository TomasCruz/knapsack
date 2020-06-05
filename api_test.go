package knapsack

import (
	"testing"

	"gotest.tools/assert"
)

type testKnapsack struct {
	benefits, weights  []float64
	capacity, expected float64
	itemCopies         int
}

func TestMaxUnboundedKnapsackValue(t *testing.T) {
	testData := []testKnapsack{
		testKnapsack{
			benefits: []float64{}, weights: []float64{}, capacity: 15.36, expected: 0},

		testKnapsack{
			benefits: []float64{12, 2, 41},
			weights:  []float64{7, 3, 20},
			capacity: 58,
			expected: 108,
		},

		testKnapsack{
			benefits: []float64{40, 30, 50, 25},
			weights:  []float64{1, 2, 5, 3},
			capacity: 2,
			expected: 80,
		},

		testKnapsack{
			benefits: []float64{40, 60, 50},
			weights:  []float64{12, 20, 15},
			capacity: 45,
			expected: 150,
		},
	}

	for _, data := range testData {
		testUnboundedFunction(t, data)
	}
}

func testUnboundedFunction(t *testing.T, testData testKnapsack) {
	maxKnapsackValue, err := MaxUnboundedKnapsackValue(
		testData.benefits,
		testData.weights,
		testData.capacity)

	assert.Assert(t, err == nil)
	assert.Assert(t, maxKnapsackValue == testData.expected, "wrong result: %f, should be %f", maxKnapsackValue, testData.expected)
}

func TestMaxZeroOneKnapsackValue(t *testing.T) {
	testData := []testKnapsack{
		testKnapsack{
			benefits:   []float64{8, 4, 0, 5, 3},
			weights:    []float64{1, 2, 3, 2, 2},
			capacity:   4,
			itemCopies: 1,
			expected:   13,
		},

		testKnapsack{
			benefits:   []float64{92, 57, 49, 68, 60, 43, 67, 84, 87, 72},
			weights:    []float64{23, 31, 29, 44, 53, 38, 63, 85, 89, 82},
			capacity:   165,
			itemCopies: 1,
			expected:   309,
		},

		testKnapsack{
			benefits:   []float64{24, 13, 23, 15, 16},
			weights:    []float64{12, 7, 11, 8, 9},
			capacity:   26,
			itemCopies: 1,
			expected:   51,
		},
	}

	for _, data := range testData {
		testBoundedFunction(t, data)
	}
}

func testBoundedFunction(t *testing.T, testData testKnapsack) {
	maxKnapsackValue, err := MaxBoundedKnapsackValue(
		testData.benefits,
		testData.weights,
		testData.itemCopies,
		testData.capacity)

	assert.Assert(t, err == nil)
	assert.Assert(t, maxKnapsackValue == testData.expected, "wrong result: %f, should be %f", maxKnapsackValue, testData.expected)
}

func TestMaxBoundedKnapsackValue(t *testing.T) {
	testData := []testKnapsack{
		testKnapsack{
			benefits: []float64{}, weights: []float64{}, capacity: 15.36, itemCopies: 5, expected: 0},

		testKnapsack{
			benefits:   []float64{12, 2, 41},
			weights:    []float64{7, 3, 20},
			capacity:   58,
			itemCopies: 2,
			expected:   108,
		},

		testKnapsack{
			benefits:   []float64{40, 30, 50, 25},
			weights:    []float64{1, 2, 5, 3},
			capacity:   2,
			itemCopies: 2,
			expected:   80,
		},

		testKnapsack{
			benefits:   []float64{40, 60, 50},
			weights:    []float64{12, 20, 15},
			capacity:   45,
			itemCopies: 3,
			expected:   150,
		},

		testKnapsack{
			benefits:   []float64{40, 60, 50},
			weights:    []float64{12, 20, 15},
			capacity:   45,
			itemCopies: 2,
			expected:   140,
		},
	}

	for _, data := range testData {
		testBoundedFunction(t, data)
	}
}
