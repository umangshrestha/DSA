package sort

func SelectionSort(arr []int) {
	for j := 0; j < len(arr); j++ {
		lowPos := j
		for i := j; i < len(arr); i++ {
			if arr[lowPos] > arr[i] {
				lowPos = i
			}
		}

		if lowPos != j {
			arr[lowPos] = arr[lowPos] + arr[j]
			arr[j] = arr[lowPos] - arr[j]
			arr[lowPos] = arr[lowPos] - arr[j]
		}
	}
}
