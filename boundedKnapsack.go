package knapsack

type boundedKnapsackItems struct {
	knapsackItems
	maxValues         []map[float64]float64
	possibleCapacites map[float64]struct{}
	itemCopies        int
}

func constructBoundedKnapsackItems(benefits, weights []float64, itemCopies int, capacity float64) (
	items knapsackEvaluator, err error) {

	boundedItems := boundedKnapsackItems{}
	boundedItems.itemCopies = itemCopies
	if boundedItems.knapsackItems, err = constructKnapsackItems(benefits, weights); err != nil {
		return
	}

	boundedItems.possibleCapacites = make(map[float64]struct{})
	boundedItems.possibleCapacites[capacity] = struct{}{}

	itemCount := len(boundedItems.benefits)
	boundedItems.maxValues = make([]map[float64]float64, itemCount)

	for i := 0; i < itemCount; i++ {
		boundedItems.maxValues[i] = make(map[float64]float64)
	}

	items = boundedItems
	return
}

func (items boundedKnapsackItems) maxKnapsackValue(capacity float64) (maxKnapsackValue float64) {
	return items.maxItemValuePerCapacity(len(items.benefits)-1, capacity)
}

// maxItemValuePerCapacity returns max value per capacity, using items with indices up to itemIndex.
func (items boundedKnapsackItems) maxItemValuePerCapacity(itemIndex int, capacity float64) (maxKnapsackValue float64) {
	if itemIndex == -1 {
		return 0
	}

	var ok bool
	if maxKnapsackValue, ok = items.maxValues[itemIndex][capacity]; ok {
		return
	}

	for i := 0; i <= itemIndex; i++ {
		for w := range items.possibleCapacites {
			maxCurrCapacityKnapsackValue := items.maxItemValuePerCapacity(i-1, w)

			for numberCopies := 1; numberCopies <= items.itemCopies; numberCopies++ {
				itemCopiesWeight := float64(numberCopies) * items.weights[i]
				if itemCopiesWeight > w {
					break
				}

				items.possibleCapacites[w-itemCopiesWeight] = struct{}{}
				itemValueAdded := items.maxItemValuePerCapacity(i-1, w-itemCopiesWeight) +
					float64(numberCopies)*items.benefits[i]

				if itemValueAdded > maxCurrCapacityKnapsackValue {
					maxCurrCapacityKnapsackValue = itemValueAdded
				}
			}

			items.maxValues[i][w] = maxCurrCapacityKnapsackValue
		}
	}

	maxKnapsackValue = items.maxValues[itemIndex][capacity]
	return
}
