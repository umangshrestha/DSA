package sort

import "testing"

var OneOff = []TestCase{
	{[]int{1, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{2, 0, 1, 3, 4, 5, 6, 7, 8, 9, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{3, 0, 1, 2, 4, 5, 6, 7, 8, 9, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{4, 0, 1, 2, 3, 5, 6, 7, 8, 9, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{5, 0, 1, 2, 3, 4, 6, 7, 8, 9, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{6, 0, 1, 2, 3, 4, 5, 7, 8, 9, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{7, 0, 1, 2, 3, 4, 5, 6, 8, 9, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{8, 0, 1, 2, 3, 4, 5, 6, 7, 9, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{0, 1, 3, 4, 5, 6, 7, 8, 9, 10, 2}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{0, 1, 2, 4, 5, 6, 7, 8, 9, 10, 3}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{0, 1, 2, 3, 5, 6, 7, 8, 9, 10, 4}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{0, 1, 2, 3, 4, 6, 7, 8, 9, 10, 5}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{0, 1, 2, 3, 4, 5, 7, 8, 9, 10, 6}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{0, 1, 2, 3, 4, 5, 6, 8, 9, 10, 7}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 8}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
}

func Benchmark_MergeSort(t *testing.B) {
	runTestCases(MergeSort, OneOff)
}

func Benchmark_BubbleSort(t *testing.B) {
	runTestCases(BubbleSort, OneOff)
}

func Benchmark_InsertionSort(t *testing.B) {
	runTestCases(IntertionSort, OneOff)
}

var VaringSize = []TestCase{
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
	{[]int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	{[]int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
	{[]int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
	{[]int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 11, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 12, 11, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 13, 12, 11, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 14, 13, 12, 11, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 15, 14, 13, 12, 11, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
}

func Benchmark_MergeSort_varyingSize(t *testing.B) {
	runTestCases(MergeSort, VaringSize)
}

func Benchmark_BubbleSort_varyingSize(t *testing.B) {
	runTestCases(BubbleSort, VaringSize)
}

func Benchmark_InsertionSort_varyingSize(t *testing.B) {
	runTestCases(IntertionSort, VaringSize)
}

var DescendingOrder = []TestCase{
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 0}, []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}},
}

func Benchmark_MergeSort_DescendingOrder(t *testing.B) {
	runTestCases(MergeSort, DescendingOrder)
}

func Benchmark_BubbleSort_DescendingOrder(t *testing.B) {
	runTestCases(BubbleSort, DescendingOrder)
}

func Benchmark_InsertionSort_DescendingOrder(t *testing.B) {
	runTestCases(IntertionSort, DescendingOrder)
}

var AscendingOrder = []TestCase{
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}},
}

func Benchmark_MergeSort_AscendingOrder(t *testing.B) {
	runTestCases(MergeSort, AscendingOrder)
}

func Benchmark_BubbleSort_AscendingOrder(t *testing.B) {
	runTestCases(BubbleSort, AscendingOrder)
}

func Benchmark_InsertionSort_AscendingOrder(t *testing.B) {
	runTestCases(IntertionSort, AscendingOrder)
}

var EmptyArray = []TestCase{
	{[]int{}, []int{}},
}

func Benchmark_MergeSort_EmptyArray(t *testing.B) {
	runTestCases(MergeSort, EmptyArray)
}

func Benchmark_BubbleSort_EmptyArray(t *testing.B) {
	runTestCases(BubbleSort, EmptyArray)
}

func Benchmark_InsertionSort_EmptyArray(t *testing.B) {
	runTestCases(BubbleSort, EmptyArray)
}

var DuplicateElements = []TestCase{
	{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
}
