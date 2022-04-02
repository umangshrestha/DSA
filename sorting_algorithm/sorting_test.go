package sort

import "testing"

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
	TestMergeSort(OneOff)
}

func Benchmark_BubbleSort(t *testing.B) {
	TestBubbleSort(OneOff)
}

func Benchmark_InsertionSort(t *testing.B) {
	TestBubbleSort(OneOff)
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
	TestMergeSort(VaringSize)
}

func Benchmark_BubbleSort_varyingSize(t *testing.B) {
	TestBubbleSort(VaringSize)
}

func Benchmark_InsertionSort_varyingSize(t *testing.B) {
	TestBubbleSort(VaringSize)
}

var DescendingOrder = []TestCase{
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 0}, []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}},
}

func Benchmark_MergeSort_DescendingOrder(t *testing.B) {
	TestMergeSort(DescendingOrder)
}

func Benchmark_BubbleSort_DescendingOrder(t *testing.B) {
	TestBubbleSort(DescendingOrder)
}

func Benchmark_InsertionSort_DescendingOrder(t *testing.B) {
	TestBubbleSort(DescendingOrder)
}

var AscendingOrder = []TestCase{
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}},
}

func Benchmark_MergeSort_AscendingOrder(t *testing.B) {
	TestMergeSort(AscendingOrder)
}

func Benchmark_BubbleSort_AscendingOrder(t *testing.B) {
	TestBubbleSort(AscendingOrder)
}

func Benchmark_InsertionSort_AscendingOrder(t *testing.B) {
	TestBubbleSort(AscendingOrder)
}

var EmptyArray = []TestCase{
	{[]int{}, []int{}},
}

func Benchmark_MergeSort_EmptyArray(t *testing.B) {
	TestMergeSort(EmptyArray)
}

func Benchmark_BubbleSort_EmptyArray(t *testing.B) {
	TestBubbleSort(EmptyArray)
}

func Benchmark_InsertionSort_EmptyArray(t *testing.B) {
	TestBubbleSort(EmptyArray)
}

var DuplicateElements = []TestCase{
	{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
}

func Benchmark_MergeSort_DuplicateElements(t *testing.B) {
	TestMergeSort(EmptyArray)
}

func Benchmark_BubbleSort_DuplicateElements(t *testing.B) {
	TestBubbleSort(EmptyArray)
}

func Benchmark_InsertionSort_DuplicateElements(t *testing.B) {
	TestBubbleSort(EmptyArray)
}
