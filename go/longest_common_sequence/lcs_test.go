package dp

import "testing"

var ElementinArray = []TestCase{
	{"", "", 0},
	{"a", "abc", 1},
	{"AGGTABTABTABTAB", "GXTXAYBTABTABTAB", 13},
	{"AGGTABGHSRCBYJSVDWFVDVSBCBVDWFDWVV", "GXTXAYBRGDVCBDVCCXVXCWQRVCBDJXCVQSQQ", 14},
}

func Benchmark_LCS(t *testing.B) {
	runTestCases(Lcs, ElementinArray)
}
