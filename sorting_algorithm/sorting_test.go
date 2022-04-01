package sort

import (
	"log"
	"testing"
)

func isEqual(output, expected []int) bool {
	if len(output) != len(expected) {
		return false
	}

	for j := 0; j < len(expected); j++ {
		if output[j] != expected[j] {
			return false
		}
	}
	return true

}
func BenchmarkMergeSort(t *testing.B) {
	for i, testCase := range TestCases {
		output := append([]int{}, testCase.arr...)
		if MergeSort(output); !isEqual(output, testCase.expected) {
			log.Fatalf("%d: MergeSort(%v) expected:%v observed: %v\n",
				i,
				testCase.arr,
				testCase.expected,
				output)
		}
	}
}

func BenchmarkBubbleSort(t *testing.B) {
	for i, testCase := range TestCases {
		output := append([]int{}, testCase.arr...)
		if BubbleSort(output); !isEqual(output, testCase.expected) {
			log.Fatalf("%d: BubbleSort(%v) expected:%v observed: %v\n",
				i,
				testCase.arr,
				testCase.expected,
				output)
		}
	}
}

func BenchmarkSelectionSort(t *testing.B) {
	for i, testCase := range TestCases {
		output := append([]int{}, testCase.arr...)
		if SelectionSort(output); !isEqual(output, testCase.expected) {
			log.Fatalf("%d: SelectionSort(%v) expected:%v observed: %v\n",
				i,
				testCase.arr,
				testCase.expected,
				output)
		}
	}
}

func BenchmarkInsertionSort(t *testing.B) {
	for i, testCase := range TestCases {
		output := append([]int{}, testCase.arr...)
		if IntertionSort(output); !isEqual(output, testCase.expected) {
			log.Fatalf("%d: IntertionSort(%v) expected:%v observed: %v\n",
				i,
				testCase.arr,
				testCase.expected,
				output)
		}
	}
}
