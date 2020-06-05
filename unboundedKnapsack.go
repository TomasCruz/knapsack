package knapsack

type unboundedKnapsackItems struct {
	knapsackItems
	maxCapacityValues map[float64]float64
}

func constructUnboundedKnapsackItems(benefits, weights []float64) (items knapsackEvaluator, err error) {
	unboundedItems := unboundedKnapsackItems{}
	if unboundedItems.knapsackItems, err = constructKnapsackItems(benefits, weights); err != nil {
		return
	}

	unboundedItems.maxCapacityValues = make(map[float64]float64)
	unboundedItems.maxCapacityValues[0] = 0

	items = unboundedItems
	return
}

func (items unboundedKnapsackItems) maxKnapsackValue(capacity float64) (maxKnapsackValue float64) {
	// if maximum value for this capacity is already stored, return it
	var ok bool
	if maxKnapsackValue, ok = items.maxCapacityValues[capacity]; ok {
		return
	}

	var currMaxVal float64
	for i := 0; i < len(items.benefits); i++ {
		// ignore item that cannot fit knapsack capacity
		if items.weights[i] > capacity {
			continue
		}

		currMaxVal = items.benefits[i] + items.maxKnapsackValue(capacity-items.weights[i])
		if currMaxVal > maxKnapsackValue {
			maxKnapsackValue = currMaxVal
		}
	}

	// store maximum value for this capacity before returning it
	items.maxCapacityValues[capacity] = maxKnapsackValue
	return
}
