package psearch

// Knuth Moris Pratt
func KMPSearch(data, pattern string) []int {
	out := []int{}
	if len(pattern) > 0 {
		lps := computeLPSArray(pattern)
		j := 0
		for i := 0; i < len(data); i++ {
			for j > 0 && pattern[j] != data[i] {
				j = lps[j-1]
			}
			if pattern[j] == data[i] {
				j++
			}
			if j == len(pattern) {
				out = append(out, i-j+1)
				j = lps[j-1]
			}
		}
	}
	return out
}

func computeLPSArray(pattern string) []int {

	lps := make([]int, len(pattern))

	for i := 1; i < len(pattern); i++ {
		j := lps[i-1]

		for j > 0 && pattern[i] != pattern[j] {
			j = lps[j-1]
		}

		if pattern[i] == pattern[j] {
			lps[i] = j + 1
		} else {
			lps[i] = j
		}
	}
	return lps
}
