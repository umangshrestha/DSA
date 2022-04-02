package psearch

import "log"

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

func TestNativeSearch(testCases []TestCase) {
	for i, testCase := range testCases {
		if output := NativeSearch(testCase.data, testCase.pattern); !isEqual(output, testCase.expected) {
			log.Fatalf("%d: NativeSearch(\"%s\", \"%s\") expected:%v observed: %v\n",
				i,
				testCase.data,
				testCase.pattern,
				testCase.expected,
				output)
		}
	}
}
