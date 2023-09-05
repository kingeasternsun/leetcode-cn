package mincost

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func minCost(costs [][]int) int {
	dp := make([][3]int, len(costs))
	dp[0][0] = costs[0][0]
	dp[0][1] = costs[0][1]
	dp[0][2] = costs[0][2]

	for i := 1; i < len(costs); i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
	}

	return min(min(dp[len(costs)-1][0], dp[len(costs)-1][1]), dp[len(costs)-1][2])
}
