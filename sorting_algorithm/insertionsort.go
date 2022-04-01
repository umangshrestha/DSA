package sort

func IntertionSort(arr []int) {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		for i := j - 1; i > 0 && key < arr[i]; i-- {
			arr[i+1] = arr[i]
		}
		arr[j] = key
	}
}
