package minpathsum

func minPathSum(grid [][]int) int {
	dp := make([][]int, 2)
	dp[0] = make([]int, len(grid[0]))
	dp[1] = make([]int, len(grid[0]))

	for rowID := 0; rowID < len(grid); rowID++ {
		dpRowID := rowID & 1
		dpPreRowID := (rowID + 1) & 1
		for colID := 0; colID < len(grid[0]); colID++ {
			if rowID == 0 && colID == 0 {
				dp[dpRowID][colID] = grid[0][0]
			} else if rowID == 0 {
				dp[dpRowID][colID] = dp[dpRowID][colID-1] + grid[rowID][colID]
			} else if colID == 0 {
				dp[dpRowID][colID] = dp[dpPreRowID][colID] + grid[rowID][colID]
			} else {
				if dp[dpRowID][colID-1] < dp[dpPreRowID][colID] {
					dp[dpRowID][colID] = dp[dpRowID][colID-1] + grid[rowID][colID]
				} else {
					dp[dpRowID][colID] = dp[dpPreRowID][colID] + grid[rowID][colID]
				}
			}
		}
	}

	return dp[(len(grid)-1)&1][len(grid[0])-1]

}
