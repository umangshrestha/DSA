package psearch

const Size = 2 << 8 // handing only UTF-8
// boyermoore
func badHeruristic(pattern string) []int {
	badChar := make([]int, Size)
	for i := 0; i < Size; i++ {
		badChar[i] = -1
	}
	for i, val := range pattern {
		badChar[val] = i
	}
	return badChar
}

func BoyerMooreSearch(data, pattern string) []int {
	p := len(pattern)
	d := len(data)

	out := []int{}
	badChar := badHeruristic(pattern)

	for s := 0; p > 0 && s+p <= d; s++ {
		i := p - 1
		for i >= 0 && pattern[i] == data[s+i] {
			i--
		}
		if i == -1 {
			out = append(out, s)
			if s+p < d {
				s += p - badChar[data[s+p]] - 1
			}
		} else if badChar[data[s+i]] < 1 {
			s += i - badChar[data[s+i]] - 1
		}
	}
	return out
}
