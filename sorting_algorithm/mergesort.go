package sort

func MergeSort(A []int) {
	if len(A) > 1 {
		mid := len(A) >> 1
		L := make([]int, mid)
		for i := 0; i < mid; i++ {
			L[i] = A[i]
		}
		R := make([]int, len(A)-mid)
		for i := 0; i < len(A)-mid; i++ {
			R[i] = A[mid+i]
		}

		MergeSort(L)
		MergeSort(R)
		Lpos := 0
		Rpos := 0
		Apos := 0

		for Lpos < len(L) && Rpos < len(R) {
			if L[Lpos] < R[Rpos] {
				A[Apos] = L[Lpos]
				Apos++
				Lpos++
			} else {
				A[Apos] = R[Rpos]
				Apos++
				Rpos++
			}
		}

		for Lpos < len(L) {
			A[Apos] = L[Lpos]
			Apos++
			Lpos++
		}

		for Rpos < len(R) {
			A[Apos] = R[Rpos]
			Apos++
			Rpos++
		}
	}
}
