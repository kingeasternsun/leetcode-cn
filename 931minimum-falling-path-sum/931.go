package main

import "math"

/*
 dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) 注意边界值的处理就可以了
*/
func minFallingPathSum(A [][]int) int {
	rows := len(A)
	if rows == 0 {
		return 0
	}
	cols := len(A[0])
	if cols == 0 {
		return 0
	}

	minV := math.MaxInt32

	if rows == 1 {
		for i := range A[0] {
			if A[0][i] < minV {
				minV = A[0][i]
			}
		}
		return minV
	}

	dp := make([][]int, 2)
	dp[0] = make([]int, cols)
	copy(dp[0], A[0])
	dp[1] = make([]int, cols)

	for i := 1; i < rows; i++ {

		preID := (i + 1) & 1

		for j := 0; j < cols; j++ {

			dp[i&1][j] = dp[preID][j]

			if j > 0 && dp[preID][j-1] < dp[i&1][j] {
				dp[i&1][j] = dp[preID][j-1]
			}

			if j < cols-1 && dp[preID][j+1] < dp[i&1][j] {
				dp[i&1][j] = dp[preID][j+1]
			}

			dp[i&1][j] += A[i][j]

			if i == rows-1 && dp[i&1][j] < minV {
				minV = dp[i&1][j]
			}

		}
	}

	return minV

}
