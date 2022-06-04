package sort

import (
	"log"
	"reflect"
	"runtime"
)

type TestCase struct {
	arr      []int
	expected []int
}

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

func swap(a *int, b *int) {
	if *a > *b {
		*a = *a + *b
		*b = *a - *b
		*a = *a - *b
	}
}

type FunctionSignature func([]int)

func runTestCases(function FunctionSignature, testCases []TestCase) {
	functionName := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	for i, testCase := range testCases {
		output := append([]int{}, testCase.arr...)
		if function(output); !isEqual(output, testCase.expected) {
			log.Fatalf("%d: %s(%v) expected:%v observed: %v\n",
				i,
				functionName,
				testCase.arr,
				testCase.expected,
				output)
		}
	}
}
