package knapsack

import (
	"testing"

	"gotest.tools/assert"
)

func TestMaxZeroOneKnapsackValue(t *testing.T) {
	testData := zeroOneKnapsackData()

	for _, data := range testData {
		testBoundedFunctionF(t, data)
		testBoundedFunction(t, data)
		testZeroOneFunctionF(t, data)
		testZeroOneFunction(t, data)
	}
}

func testZeroOneFunctionF(t *testing.T, testData testKnapsack) {
	maxKnapsackValue, err := MaxZeroOneKnapsackValueF(
		testData.benefits,
		testData.weights,
		testData.capacity)

	assert.Assert(t, err == nil)
	assert.Assert(t, maxKnapsackValue == testData.expected, "wrong result: %f, should be %f", maxKnapsackValue, testData.expected)
}

func testZeroOneFunction(t *testing.T, testData testKnapsack) {
	items, err := constructKnapsackItems(testData.benefits, testData.weights)
	if err != nil {
		return
	}

	maxKnapsackValue, err := MaxZeroOneKnapsackValue(items, testData.capacity)

	assert.Assert(t, err == nil)
	assert.Assert(t, maxKnapsackValue == testData.expected, "wrong result: %f, should be %f", maxKnapsackValue, testData.expected)
}
