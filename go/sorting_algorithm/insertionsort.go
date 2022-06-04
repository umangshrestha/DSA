package sort

func IntertionSort(arr []int) {
	for j := 1; j < len(arr); j++ {
		for i := j; i > 0; i-- {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
	}
}
