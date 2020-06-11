library for resolving knapsack problems with float64 weights

# Item
helper struct, Item holds single item info, i.e. benefit and weight

# unbounded-knapsack
functions for calculating maximum value of knapsack with unlimited number of copies of items used. Knapsack weight is limited by capacity.

func MaxUnboundedKnapsackValueF(benefits, weights []float64, capacity float64) (maxKnapsackValue float64, err error)
func MaxUnboundedKnapsackValue(items []Item, capacity float64) (maxKnapsackValue float64, err error)

# bounded-knapsack
functions for calculating maximum value of knapsack with number of copies of items limited to itemCopies. Knapsack weight is limited by capacity.

func MaxBoundedKnapsackValueF(benefits, weights []float64, itemCopies []int, capacity float64) (maxKnapsackValue float64, err error)
func MaxBoundedKnapsackValue(items []Item, itemCopies []int, capacity float64) (maxKnapsackValue float64, err error)

## one-zero knapsack
0-1 knapsack problem is a special case of bounded knapsack problem, having itemCopies one

func MaxZeroOneKnapsackValueF(benefits, weights []float64, capacity float64) (maxKnapsackValue float64, err error)
func MaxZeroOneKnapsackValue(items []Item, capacity float64) (maxKnapsackValue float64, err error)
