package sort

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

func swap(a *int, b *int) {
	if *a > *b {
		*a = *a + *b
		*b = *a - *b
		*a = *a - *b
	}
}
