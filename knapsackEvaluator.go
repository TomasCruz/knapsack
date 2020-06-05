package knapsack

type knapsackEvaluator interface {
	maxKnapsackValue(capacity float64) (maxKnapsackValue float64)
}
