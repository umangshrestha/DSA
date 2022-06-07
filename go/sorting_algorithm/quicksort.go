package sort

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	i++
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func _quicksort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		_quicksort(arr, low, p-1)
		_quicksort(arr, p+1, high)
	}
}
func QuickSort(arr []int) {
	_quicksort(arr, 0, len(arr)-1)
}
