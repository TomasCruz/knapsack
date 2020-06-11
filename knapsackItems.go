package knapsack

import (
	"math"

	"github.com/pkg/errors"
)

func constructKnapsackItems(benefits, weights []float64) (items []Item, err error) {
	if benefits == nil || weights == nil || len(benefits) != len(weights) {
		err = errors.New("items not valid")
		return
	}

	items = make([]Item, len(benefits))
	for i := 0; i < len(benefits); i++ {
		items[i] = Item{benefit: benefits[i], weight: weights[i]}
	}

	err = verifyKnapsackItems(items)
	return
}

func verifyKnapsackItems(items []Item) (err error) {
	errInvalidItems := errors.New("items not valid")
	if items == nil {
		err = errInvalidItems
		return
	}

	for _, currItem := range items {
		if !verifySingleKnapsackItem(currItem.benefit, currItem.weight) {
			err = errInvalidItems
			return
		}
	}

	return
}

func verifySingleKnapsackItem(benefit, weight float64) bool {
	return verifyFiniteFloat(benefit) && verifyFiniteFloat(weight)
}

func verifyFiniteFloat(x float64) bool {
	return x >= 0.0 && !math.IsInf(x, 0) && !math.IsNaN(x)
}
