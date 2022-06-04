package taylor

import (
	"math"
	"testing"
)

func Benchmark_Sin(t *testing.B) {
	runTestCases(Sin, math.Sin)
}

func Benchmark_Cos(t *testing.B) {
	runTestCases(Cos, math.Cos)
}

func Benchmark_Tan(t *testing.B) {
	runTestCases(Tan, math.Tan)
}
