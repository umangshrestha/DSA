package ssearch

type TestCase struct {
	data     string
	pattern  string
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
