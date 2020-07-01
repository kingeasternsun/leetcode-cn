package main

import "math"

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
func minPathSum(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}

	cols := len(grid[0])
	if cols == 0 {
		return 0
	}

	dp := make([][]int, 2)
	dp[0] = make([]int, cols)
	dp[1] = make([]int, cols)

	// copy(dp[(rows-1)&1], grid[rows-1])
	for j := cols - 1; j >= 0; j-- {
		// dp[(rows-1)&1][j] = dp[(rows-1)&1][j] + dp[(rows-1)&1][j+1]
		dp[rows&1][j] = math.MaxInt32 // mock a new row
	}

	dp[rows&1][cols-1] = 0

	// from bottom to up,from right to left
	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			if col == cols-1 {
				dp[row&1][col] = grid[row][col] + dp[(row+1)&1][col]
				continue
			}

			dp[row&1][col] = grid[row][col] + min(dp[(row+1)&1][col], dp[(row)&1][col+1])
		}
	}

	return dp[0][0]
}
