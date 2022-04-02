package search

import "testing"

var ElementinArray = []TestCase{
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 0},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 1},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 2},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 3},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, 4},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 5},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 6},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 7},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 8},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 9},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 10},
}

func Benchmark_BinarySearch(t *testing.B) {
	runTestCases(BinarySearch, ElementinArray)
}

func Benchmark_LinearSearch(t *testing.B) {
	runTestCases(LinearSearch, ElementinArray)
}

var VaringSize = []TestCase{
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 0},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 1, 1},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 2, 2},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 3, 3},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 4, 4},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 5, 5},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 6, 6},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, 7, 7},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, 8, 8},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, 9, 9},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 10, 10},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 11, 11},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 12, 12},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 0},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 11, 11},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 10, 10},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 9, 9},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 8, 8},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 7, 7},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 6, 6},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, 5, 5},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, 4, 4},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, 3, 3},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 2, 2},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 1, 1},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 0, 0},
}

func Benchmark_BinarySearch_VaringSize(t *testing.B) {
	runTestCases(BinarySearch, VaringSize)
}

func Benchmark_LinearSearch_VaringSize(t *testing.B) {
	runTestCases(LinearSearch, VaringSize)
}

var DuplicateElements = []TestCase{
	{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0, 0},
	{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 9, 0},
	{[]int{1, 1, 1, 1, 1, 1, 9, 9, 9, 9, 9}, 9, 6},
}

func Benchmark_BinarySearch_DuplicateElements(t *testing.B) {
	runTestCases(BinarySearch, DuplicateElements)
}

func Benchmark_LinearSearch_DuplicateElements(t *testing.B) {
	runTestCases(LinearSearch, DuplicateElements)
}

var EmptyArray = []TestCase{
	{[]int{}, -1, -1},
	{[]int{}, 0, -1},
	{[]int{}, 1, -1},
	{[]int{}, 2, -1},
}

func Benchmark_BinarySearch_EmptyArray(t *testing.B) {
	runTestCases(BinarySearch, EmptyArray)
}

func Benchmark_LinearSearch_EmptyArray(t *testing.B) {
	runTestCases(LinearSearch, EmptyArray)
}

var NotInArray = []TestCase{
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1, -1},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, -1},
	{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 11, -1},
	{[]int{1, 1, 1, 1, 1, 1, 9, 9, 9, 9, 9}, 0, -1},
	{[]int{1, 1, 1, 1, 1, 1, 9, 9, 9, 9, 9}, 10, -1},
}

func Benchmark_BinarySearch_NotInArray(t *testing.B) {
	runTestCases(BinarySearch, NotInArray)
}

func Benchmark_LinearSearch_NotInArray(t *testing.B) {
	runTestCases(LinearSearch, NotInArray)
}
