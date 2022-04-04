package sort

func QuickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		sort(arr, p+1, high)
		sort(arr, low, p-1)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	j := low
	for i := low; i < high; i++ {
		if arr[i] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}
	arr[j], arr[high] = arr[high], arr[j]
	return j
}
