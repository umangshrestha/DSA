package search

import "log"

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr)
	pos := -1
	for low < high {
		mid := (low + high) >> 1
		if target > arr[mid] {
			low = mid + 1
		} else {
			high = mid
			if target == arr[mid] {
				pos = mid
			}
		}
	}
	return pos
}

func TestBinarySearch(testCases []TestCase) {
	for i, testCase := range testCases {
		if output := BinarySearch(testCase.arr, testCase.target); output != testCase.expected {
			log.Fatalf("%d: BinarySearch(%v, %d) expected:%d observed: %d\n",
				i,
				testCase.arr,
				testCase.target,
				testCase.expected,
				output)
		}
	}
}
