package ssearch

import (
	"log"
)

const Size = 2 << 8 // handing only UTF-8

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

func BoyerMoore(data, pattern string) []int {
	p := len(pattern)
	d := len(data)

	out := []int{}
	badChar := badHeruristic(pattern)
	s := 0
	for s+p < d {
		i := p - 1
		for i >= 0 && pattern[i] == data[s+i] {
			i--
		}
		if i == -1 {
			out = append(out, s)
			if s+p < d {
				s += i - badChar[data[s+p]] + 1
			} else {
				s++
			}
		} else if badChar[data[s+i]] <= i {
			s += i - badChar[data[s+i]]
		} else {
			s++

		}
	}
	return out
}

func TestBoyerMoore(testCases []TestCase) {
	for i, testCase := range testCases {
		if output := BoyerMoore(testCase.data, testCase.pattern); !isEqual(output, testCase.expected) {
			log.Fatalf("%d: BoyerMoore(\"%s\", \"%s\") expected:%v observed: %v\n",
				i,
				testCase.data,
				testCase.pattern,
				testCase.expected,
				output)
		}
	}
}
