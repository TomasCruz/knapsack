package knapsack

import "errors"

type knapsackItems struct {
	benefits, weights []float64
}

func constructKnapsackItems(benefits, weights []float64) (items knapsackItems, err error) {
	if benefits == nil || weights == nil || len(benefits) != len(weights) {
		err = errors.New("items not valid")
		return
	}

	items.benefits = benefits
	items.weights = weights
	return
}
