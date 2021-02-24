/*
 * @Description: 832
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-24 09:07:07
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-24 09:24:00
 * @FilePath: \832flipAndInvertImage\832.go
 */
package leetcode

func flipAndInvertImage(A [][]int) [][]int {
	if len(A) == 0 {
		return A
	}
	cols := len(A[0])
	for i := 0; i < len(A); i++ {
		for beg, end := 0, cols-1; beg <= end; beg, end = beg+1, end-1 {

			if beg != end {
				A[i][beg], A[i][end] = A[i][end], A[i][beg]
				A[i][beg] ^= 1
			}
			A[i][end] ^= 1
		}
	}

	return A
}
