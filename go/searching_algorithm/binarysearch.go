package search

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr)
	pos := -1
	for low < high {
		mid := (low + high) >> 1
		if target > arr[mid] {
			low = mid + 1
		} else {
			high = mid
			if target == arr[mid] {
				pos = mid
			}
		}
	}
	return pos
}
