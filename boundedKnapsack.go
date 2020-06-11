package knapsack

import (
	"github.com/pkg/errors"
)

type boundedKnapsackItems struct {
	knapsackItems []Item
	maxValues     []map[float64]float64
	itemCopies    []int
}

func constructBoundedKnapsackItemsF(benefits, weights []float64, itemCopies []int) (
	evaluator knapsackEvaluator, err error) {

	var boundedItems boundedKnapsackItems
	if boundedItems, err = initItemCopiesAndMap(len(benefits), itemCopies); err != nil {
		return
	}

	if boundedItems.knapsackItems, err = constructKnapsackItems(benefits, weights); err != nil {
		return
	}

	evaluator = boundedItems
	return
}

func constructBoundedKnapsackItems(items []Item, itemCopies []int) (
	evaluator knapsackEvaluator, err error) {

	var boundedItems boundedKnapsackItems
	if boundedItems, err = initItemCopiesAndMap(len(items), itemCopies); err != nil {
		return
	}

	if err = verifyKnapsackItems(items); err != nil {
		return
	}

	boundedItems.knapsackItems = items

	evaluator = boundedItems
	return
}

func initItemCopiesAndMap(itemCount int, itemCopies []int) (
	boundedItems boundedKnapsackItems, err error) {

	if itemCopies == nil || len(itemCopies) != itemCount {
		err = errors.New("Cardinalities of itemCopies and items must be equal")
		return
	}

	boundedItems.itemCopies = itemCopies

	boundedItems.maxValues = make([]map[float64]float64, itemCount)
	for i := 0; i < itemCount; i++ {
		boundedItems.maxValues[i] = make(map[float64]float64)
	}

	return
}

func (boundedItems boundedKnapsackItems) maxKnapsackValue(capacity float64) (maxKnapsackValue float64) {
	return boundedItems.maxItemValuePerCapacity(len(boundedItems.knapsackItems)-1, capacity)
}

// maxItemValuePerCapacity returns max value per capacity, using items with indices up to itemIndex.
func (boundedItems boundedKnapsackItems) maxItemValuePerCapacity(itemIndex int, capacity float64) (maxKnapsackValue float64) {
	if itemIndex == -1 {
		return 0
	}

	var ok bool
	if maxKnapsackValue, ok = boundedItems.maxValues[itemIndex][capacity]; ok {
		return
	}

	for i := 0; i <= itemIndex; i++ {
		maxCurrCapacityKnapsackValue := boundedItems.maxItemValuePerCapacity(i-1, capacity)

		for numberCopies := 1; numberCopies <= boundedItems.itemCopies[i]; numberCopies++ {
			itemCopiesWeight := float64(numberCopies) * boundedItems.knapsackItems[i].weight
			if itemCopiesWeight > capacity {
				break
			}

			itemValueAdded := boundedItems.maxItemValuePerCapacity(i-1, capacity-itemCopiesWeight) +
				float64(numberCopies)*boundedItems.knapsackItems[i].benefit

			if itemValueAdded > maxCurrCapacityKnapsackValue {
				maxCurrCapacityKnapsackValue = itemValueAdded
			}
		}

		boundedItems.maxValues[i][capacity] = maxCurrCapacityKnapsackValue
	}

	maxKnapsackValue = boundedItems.maxValues[itemIndex][capacity]
	return
}
