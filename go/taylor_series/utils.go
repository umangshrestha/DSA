package taylor

import (
	"log"
	"math"
	"reflect"
	"runtime"
)

type FunctionSignature func(float64) float64

const FloatPercision = 1e-9
const (
	start = -3.0
	end   = 3.0
	step  = 0.1
)

func runTestCases(implementedFunction, refrenceFunction FunctionSignature) {
	functionName := runtime.FuncForPC(reflect.ValueOf(implementedFunction).Pointer()).Name()
	count := 0

	for i := start; i < end; i += step {
		input := i * math.Pi

		expected := refrenceFunction(input)
		output := implementedFunction(input)

		if math.Abs(output-expected) > FloatPercision {
			log.Fatalf("%d: %s(%v) expected:%v observed: %v\n",
				count,
				functionName,
				input,
				expected,
				output)
		}

		count++
	}
}
