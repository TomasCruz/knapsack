package knapsack

import (
	"testing"

	"gotest.tools/assert"
)

func TestMaxBoundedKnapsackValue(t *testing.T) {
	testData := boundedKnapsackData()

	for _, data := range testData {
		testBoundedFunctionF(t, data)
		testBoundedFunction(t, data)
	}
}

func testBoundedFunctionF(t *testing.T, testData testKnapsack) {
	maxKnapsackValue, err := MaxBoundedKnapsackValueF(
		testData.benefits,
		testData.weights,
		testData.itemCopies,
		testData.capacity)

	assert.Assert(t, err == nil)
	assert.Assert(t, maxKnapsackValue == testData.expected, "wrong result: %f, should be %f", maxKnapsackValue, testData.expected)
}

func testBoundedFunction(t *testing.T, testData testKnapsack) {
	items, err := constructKnapsackItems(testData.benefits, testData.weights)
	if err != nil {
		return
	}

	maxKnapsackValue, err := MaxBoundedKnapsackValue(items, testData.itemCopies, testData.capacity)

	assert.Assert(t, err == nil)
	assert.Assert(t, maxKnapsackValue == testData.expected, "wrong result: %f, should be %f", maxKnapsackValue, testData.expected)
}
