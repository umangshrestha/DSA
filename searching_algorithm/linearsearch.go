package search

import "log"

func LinearSearch(arr []int, target int) int {
	for pos, value := range arr {
		if target == value {
			return pos
		}
	}
	return -1
}

func TestLinearSearch(testCases []TestCase) {
	for i, testCase := range testCases {
		if output := LinearSearch(testCase.arr, testCase.target); output != testCase.expected {
			log.Fatalf("%d: LinearSearch(%v, %d) expected:%d observed: %d\n",
				i,
				testCase.arr,
				testCase.target,
				testCase.expected,
				output)
		}
	}
}
