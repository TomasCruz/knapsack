package knapsack

func boundedKnapsackData() (testData []testKnapsack) {
	testData = []testKnapsack{
		testKnapsack{
			benefits: []float64{}, weights: []float64{}, capacity: 15.36, itemCopies: []int{}, expected: 0},

		testKnapsack{
			benefits:   []float64{12, 2, 41},
			weights:    []float64{7, 3, 20},
			capacity:   58,
			itemCopies: []int{1, 2, 1},
			expected:   57,
		},

		testKnapsack{
			benefits:   []float64{12, 2, 41},
			weights:    []float64{7, 3, 20},
			capacity:   28,
			itemCopies: []int{1, 7, 1},
			expected:   53,
		},

		testKnapsack{
			benefits:   []float64{40, 30, 50, 25},
			weights:    []float64{1, 2, 5, 3},
			capacity:   2,
			itemCopies: []int{2, 2, 2, 2},
			expected:   80,
		},

		testKnapsack{
			benefits:   []float64{40, 60, 50},
			weights:    []float64{12, 20, 15},
			capacity:   45,
			itemCopies: []int{3, 3, 3},
			expected:   150,
		},

		testKnapsack{
			benefits:   []float64{40, 60, 50},
			weights:    []float64{12, 20, 15},
			capacity:   45,
			itemCopies: []int{2, 2, 2},
			expected:   140,
		},

		testKnapsack{
			benefits:   []float64{12.7, 19.76, 7.43},
			weights:    []float64{1.26, 12.95, 13.23},
			capacity:   26.69,
			itemCopies: []int{2, 2, 2},
			expected:   45.16,
		},

		testKnapsack{
			benefits:   []float64{24.71, 13.23, 15.23},
			weights:    []float64{12.25, 0.7, 11.0009},
			capacity:   100,
			itemCopies: []int{2, 2, 2},
			expected:   106.34,
		},

		testKnapsack{
			benefits:   []float64{24.71, 13.23, 15.23},
			weights:    []float64{12.25, 0.7, 11.0009},
			capacity:   23.96,
			itemCopies: []int{1, 2, 2},
			expected:   56.92,
		},

		testKnapsack{
			benefits:   []float64{24.71, 13.23, 15.23},
			weights:    []float64{12.25, 0.7, 11.0009},
			capacity:   0.5,
			itemCopies: []int{1, 2, 2},
			expected:   0,
		},
	}

	return
}
