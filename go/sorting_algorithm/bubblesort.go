package sort

func BubbleSort(arr []int) {
	for j := 0; j < len(arr); j++ {
		isSorted := true

		for i := 0; i < len(arr)-j-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
			isSorted = false
		}

		if isSorted {
			break
		}
	}
}
