package knapsack

import (
	"testing"

	"gotest.tools/assert"
)

func TestMaxUnboundedKnapsackValue(t *testing.T) {
	testData := unboundedKnapsackData()

	for _, data := range testData {
		testUnboundedFunctionF(t, data)
		testUnboundedFunction(t, data)
	}
}

func testUnboundedFunctionF(t *testing.T, testData testKnapsack) {
	maxKnapsackValue, err := MaxUnboundedKnapsackValueF(
		testData.benefits,
		testData.weights,
		testData.capacity)

	assert.Assert(t, err == nil)
	assert.Assert(t, maxKnapsackValue == testData.expected, "wrong result: %f, should be %f", maxKnapsackValue, testData.expected)
}

func testUnboundedFunction(t *testing.T, testData testKnapsack) {
	items, err := constructKnapsackItems(testData.benefits, testData.weights)
	if err != nil {
		return
	}

	maxKnapsackValue, err := MaxUnboundedKnapsackValue(items, testData.capacity)

	assert.Assert(t, err == nil)
	assert.Assert(t, maxKnapsackValue == testData.expected, "wrong result: %f, should be %f", maxKnapsackValue, testData.expected)
}
