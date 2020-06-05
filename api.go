package knapsack

// MaxUnboundedKnapsackValue calculates maximum value of knapsack with unbounded number of copies of items used.
// Unbounded knapsack has capacity as limit of total weight. Items are given with their benefits and weights.
func MaxUnboundedKnapsackValue(benefits, weights []float64, capacity float64) (maxKnapsackValue float64, err error) {
	var items knapsackEvaluator
	if items, err = constructUnboundedKnapsackItems(benefits, weights); err != nil {
		return
	}

	maxKnapsackValue = items.maxKnapsackValue(capacity)
	return
}

// MaxBoundedKnapsackValue calculates maximum value of knapsack with number of copies of items limited to itemCopies.
// Bounded knapsack has capacity as limit of total weight. Items are given with their benefits and weights.
func MaxBoundedKnapsackValue(benefits, weights []float64, itemCopies int, capacity float64) (maxKnapsackValue float64, err error) {
	var items knapsackEvaluator
	if items, err = constructBoundedKnapsackItems(benefits, weights, itemCopies, capacity); err != nil {
		return
	}

	maxKnapsackValue = items.maxKnapsackValue(capacity)
	return
}
