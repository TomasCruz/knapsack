package knapsack

func unboundedKnapsackData() (testData []testKnapsack) {
	testData = []testKnapsack{
		testKnapsack{
			benefits: []float64{}, weights: []float64{}, capacity: 15.36, expected: 0},

		testKnapsack{
			benefits: []float64{12, 2, 41},
			weights:  []float64{7, 3, 20},
			capacity: 58,
			expected: 108,
		},

		testKnapsack{
			benefits: []float64{40, 30, 50, 25},
			weights:  []float64{1, 2, 5, 3},
			capacity: 2,
			expected: 80,
		},

		testKnapsack{
			benefits: []float64{40, 60, 50},
			weights:  []float64{12, 20, 15},
			capacity: 45,
			expected: 150,
		},
	}

	return
}
