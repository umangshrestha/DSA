package psearch

func NativeSearch(data, pattern string) []int {
	p := len(pattern)
	d := len(data)
	out := []int{}
	for i := 0; i < d-p+1; i++ {
		j := 0
		for j = 0; j < p && pattern[j] == data[i+j]; j++ {
		}
		if j == p {
			out = append(out, i)
		}

	}
	return out
}
