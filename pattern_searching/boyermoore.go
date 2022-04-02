package psearch

import (
	"fmt"
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
	for s+p <= d {
		i := p - 1
		for i >= 0 && pattern[i] == data[s+i] {
			i--
		}
		log.Println(pattern, i, s+p <= d)
		if i == -1 {
			out = append(out, s)
			fmt.Println(pattern, out)

			if s+p <= d {
				log.Println("0=>", data[s+p], s, s+p)
				s++
			} else {
				log.Println("1=>", data[s+p], s, s+1)
				s += i - badChar[data[s+p]]

			}
		} else if badChar[data[s+i]] < 1 {
			log.Println("2=>", data[s+i], s+i, s+i-badChar[data[s+i]])
			s += i - badChar[data[s+i]]
		} else {
			log.Println("3=>", data[s+i], s, s+1)
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
