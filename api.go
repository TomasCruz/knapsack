package knapsack

// Item holds single item info, i.e. benefit and weight
type Item struct {
	benefit float64
	weight  float64
}

// MaxUnboundedKnapsackValueF calculates maximum value of knapsack with unlimited number of copies of items used.
// Knapsack weight is limited by capacity.. Items are given with their benefits and weights.
func MaxUnboundedKnapsackValueF(benefits, weights []float64, capacity float64) (maxKnapsackValue float64, err error) {
	var evaluator knapsackEvaluator
	if evaluator, err = constructUnboundedKnapsackItemsF(benefits, weights); err != nil {
		return
	}

	maxKnapsackValue = evaluator.maxKnapsackValue(capacity)
	return
}

// MaxUnboundedKnapsackValue calculates maximum value of knapsack with unlimited number of copies of items used.
// Knapsack weight is limited by capacity.
func MaxUnboundedKnapsackValue(items []Item, capacity float64) (maxKnapsackValue float64, err error) {
	var evaluator knapsackEvaluator
	if evaluator, err = constructUnboundedKnapsackItems(items); err != nil {
		return
	}

	maxKnapsackValue = evaluator.maxKnapsackValue(capacity)
	return
}

// MaxBoundedKnapsackValueF calculates maximum value of knapsack with number of copies of items limited to itemCopies.
// Knapsack weight is limited by capacity. Items are given with their benefits and weights.
func MaxBoundedKnapsackValueF(benefits, weights []float64, itemCopies []int, capacity float64) (maxKnapsackValue float64, err error) {
	var evaluator knapsackEvaluator
	if evaluator, err = constructBoundedKnapsackItemsF(benefits, weights, itemCopies); err != nil {
		return
	}

	maxKnapsackValue = evaluator.maxKnapsackValue(capacity)
	return
}

// MaxBoundedKnapsackValue calculates maximum value of knapsack with number of copies of items limited to itemCopies.
// Knapsack weight is limited by capacity.
func MaxBoundedKnapsackValue(items []Item, itemCopies []int, capacity float64) (maxKnapsackValue float64, err error) {
	var evaluator knapsackEvaluator
	if evaluator, err = constructBoundedKnapsackItems(items, itemCopies); err != nil {
		return
	}

	maxKnapsackValue = evaluator.maxKnapsackValue(capacity)
	return
}

// MaxZeroOneKnapsackValueF calculates maximum value of knapsack with number of copies of items limited to 1.
// Knapsack weight is limited by capacity. Items are given with their benefits and weights.
func MaxZeroOneKnapsackValueF(benefits, weights []float64, capacity float64) (maxKnapsackValue float64, err error) {
	itemCopies := make([]int, len(benefits))
	for i := 0; i < len(itemCopies); i++ {
		itemCopies[i] = 1
	}

	var evaluator knapsackEvaluator
	if evaluator, err = constructBoundedKnapsackItemsF(benefits, weights, itemCopies); err != nil {
		return
	}

	maxKnapsackValue = evaluator.maxKnapsackValue(capacity)
	return
}

// MaxZeroOneKnapsackValue calculates maximum value of knapsack with number of copies of items limited to 1.
// Knapsack weight is limited by capacity.
func MaxZeroOneKnapsackValue(items []Item, capacity float64) (maxKnapsackValue float64, err error) {
	itemCopies := make([]int, len(items))
	for i := 0; i < len(itemCopies); i++ {
		itemCopies[i] = 1
	}

	var evaluator knapsackEvaluator
	if evaluator, err = constructBoundedKnapsackItems(items, itemCopies); err != nil {
		return
	}

	maxKnapsackValue = evaluator.maxKnapsackValue(capacity)
	return
}
