package sort

import (
	"log"
)

func IntertionSort(arr []int) {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		for i := j - 1; i > 0 && key < arr[i]; i-- {
			arr[i+1] = arr[i]
		}
		arr[j] = key
	}
}

func TestInsertionSort(testCases []TestCase) {
	for i, testCase := range testCases {
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
