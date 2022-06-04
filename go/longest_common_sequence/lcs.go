// longest common sequence
package dp

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Lcs(str1, str2 string) int {
	l1 := len(str1)
	l2 := len(str2)
	// allocating memory fir 2 dimentional matrix
	a := make([][]int, l1+1)
	for i := range a {
		a[i] = make([]int, l2+1)
	}

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			if str1[i] == str2[j] {
				a[i][j] = a[i+1][j+1] + 1
			} else {
				a[i][j] = max(a[i][j+1], a[i+1][j])
			}
		}
	}

	return a[0][0]
}
