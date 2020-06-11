package knapsack

func zeroOneKnapsackData() (testData []testKnapsack) {
	testData = []testKnapsack{
		testKnapsack{
			benefits:   []float64{8, 4, 0, 5, 3},
			weights:    []float64{1, 2, 3, 2, 2},
			capacity:   4,
			itemCopies: []int{1, 1, 1, 1, 1},
			expected:   13,
		},

		testKnapsack{
			benefits:   []float64{92, 57, 49, 68, 60, 43, 67, 84, 87, 72},
			weights:    []float64{23, 31, 29, 44, 53, 38, 63, 85, 89, 82},
			capacity:   165,
			itemCopies: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected:   309,
		},

		testKnapsack{
			benefits:   []float64{24, 13, 23, 15, 16},
			weights:    []float64{12, 7, 11, 8, 9},
			capacity:   26,
			itemCopies: []int{1, 1, 1, 1, 1},
			expected:   51,
		},

		testKnapsack{
			benefits:   []float64{24.71, 13.23, 15.23},
			weights:    []float64{12.25, 0.7, 11.0009},
			capacity:   100,
			itemCopies: []int{1, 1, 1},
			expected:   53.17,
		},

		testKnapsack{
			benefits:   []float64{24.71, 13.23, 15.23},
			weights:    []float64{12.25, 0.7, 11.0009},
			capacity:   14.61,
			itemCopies: []int{1, 1, 1},
			expected:   37.94,
		},
	}

	return
}
