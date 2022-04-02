package psearch

import "testing"

var OneWord = []TestCase{
	{"Apple", "A", []int{0}},
	{"Apple", "Ap", []int{0}},
	{"Apple", "App", []int{0}},
	{"Apple", "Appl", []int{0}},
	{"Apple", "Apple", []int{0}},
	{"Apple", "pple", []int{1}},
	{"Apple", "ple", []int{2}},
	{"Apple", "le", []int{3}},
	{"Apple", "e", []int{4}},
	{"Apple", "p", []int{1, 2}},
}

func Benchmark_NativeSearch(t *testing.B) {
	TestNativeSearch(OneWord)
}

func Benchmark_BoyerMoore(t *testing.B) {
	TestBoyerMoore(OneWord)
}

var EmptyData = []TestCase{
	{"", "", []int{}},
	{"Apple", "", []int{}},
	{"", "Apple", []int{}},
}

func Benchmark_NativeSearch_EmptyData(t *testing.B) {
	TestNativeSearch(OneWord)
}
func Benchmark_BoyerMoore_EmptyData(t *testing.B) {
	TestBoyerMoore(EmptyData)
}
