package knapsack

type unboundedKnapsackItems struct {
	knapsackItems     []Item
	maxCapacityValues map[float64]float64
}

func constructUnboundedKnapsackItemsF(benefits, weights []float64) (evaluator knapsackEvaluator, err error) {
	unboundedItems := unboundedKnapsackItems{}
	if unboundedItems.knapsackItems, err = constructKnapsackItems(benefits, weights); err != nil {
		return
	}

	unboundedItems.initializeMaxCapacityValues()

	evaluator = unboundedItems
	return
}

func constructUnboundedKnapsackItems(items []Item) (evaluator knapsackEvaluator, err error) {
	unboundedItems := unboundedKnapsackItems{}

	if err = verifyKnapsackItems(items); err != nil {
		return
	}

	unboundedItems.knapsackItems = items
	unboundedItems.initializeMaxCapacityValues()

	evaluator = unboundedItems
	return
}

func (unboundedItems *unboundedKnapsackItems) initializeMaxCapacityValues() {
	unboundedItems.maxCapacityValues = make(map[float64]float64)
	unboundedItems.maxCapacityValues[0] = 0
}

func (unboundedItems unboundedKnapsackItems) maxKnapsackValue(capacity float64) (maxKnapsackValue float64) {
	// if maximum value for this capacity is already stored, return it
	var ok bool
	if maxKnapsackValue, ok = unboundedItems.maxCapacityValues[capacity]; ok {
		return
	}

	var currMaxVal float64
	for i := 0; i < len(unboundedItems.knapsackItems); i++ {
		// ignore item that cannot fit knapsack capacity
		if unboundedItems.knapsackItems[i].weight > capacity {
			continue
		}

		currMaxVal = unboundedItems.maxKnapsackValue(capacity-unboundedItems.knapsackItems[i].weight) +
			unboundedItems.knapsackItems[i].benefit

		if currMaxVal > maxKnapsackValue {
			maxKnapsackValue = currMaxVal
		}
	}

	// store maximum value for this capacity before returning it
	unboundedItems.maxCapacityValues[capacity] = maxKnapsackValue
	return
}
