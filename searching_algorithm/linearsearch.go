package search

func LinearSearch(arr []int, target int) int {
	for pos, value := range arr {
		if target == value {
			return pos
		}
	}
	return -1
}
