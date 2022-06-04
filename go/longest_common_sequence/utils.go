package dp

import (
	"log"
	"reflect"
	"runtime"
)

type TestCase struct {
	str1     string
	str2     string
	expected int
}

type FunctionSignature func(string, string) int

func runTestCases(function FunctionSignature, testCases []TestCase) {
	functionName := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	for i, testCase := range testCases {
		if output := function(testCase.str1, testCase.str2); output != testCase.expected {
			log.Fatalf("%d: %s(%s, %s) expected:%d observed: %d\n",
				i,
				functionName,
				testCase.str1,
				testCase.str2,
				testCase.expected,
				output)
		}
	}
}
