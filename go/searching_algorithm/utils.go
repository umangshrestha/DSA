package search

import (
	"log"
	"reflect"
	"runtime"
)

type TestCase struct {
	arr      []int
	target   int
	expected int
}

type FunctionSignature func([]int, int) int

func runTestCases(function FunctionSignature, testCases []TestCase) {
	functionName := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	for i, testCase := range testCases {
		if output := function(testCase.arr, testCase.target); output != testCase.expected {
			log.Fatalf("%d: %s(%v, %d) expected:%d observed: %d\n",
				i,
				functionName,
				testCase.arr,
				testCase.target,
				testCase.expected,
				output)
		}
	}
}
