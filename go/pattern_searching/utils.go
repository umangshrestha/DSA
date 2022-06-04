package psearch

import (
	"log"
	"reflect"
	"runtime"
)

type TestCase struct {
	data     string
	pattern  string
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

type FunctionSignature func(string, string) []int

func runTestCases(function FunctionSignature, testCases []TestCase) {
	functionName := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	for i, testCase := range testCases {
		if output := function(testCase.data, testCase.pattern); !isEqual(output, testCase.expected) {
			log.Fatalf("%d: %s(\"%s\", \"%s\") expected:%v observed: %v\n",
				i,
				functionName,
				testCase.data,
				testCase.pattern,
				testCase.expected,
				output)
		}
	}
}
