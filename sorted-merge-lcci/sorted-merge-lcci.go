package main

/*
面试题 10.01. 合并排序的数组
*/

func merge(A []int, m int, B []int, n int) {

	if m == 0 {
		copy(A, B)
		return
	}

	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		j := m - 1
		for j >= 0 {
			if B[i] < A[j] {
				A[j+1] = A[j]
				j--
			} else {
				A[j+1] = B[i]
				m++
				break
			}
		}

		//B[i]比m中所有数都小
		if j < 0 {
			A[0] = B[i]
			m++
		}

	}

}
