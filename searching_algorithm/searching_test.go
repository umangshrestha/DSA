package search

import (
	"log"
	"testing"
)

func BenchmarkBinarySearch(t *testing.B) {
	for i, testCase := range TestCases {
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

func BenchmarkLinearSearch(t *testing.B) {
	for i, testCase := range TestCases {
		if output := BinarySearch(testCase.arr, testCase.target); output != testCase.expected {
			log.Fatalf("%d: LinearSearch(%v, %d) expected:%d observed: %d\n",
				i,
				testCase.arr,
				testCase.target,
				testCase.expected,
				output)
		}
	}
}
