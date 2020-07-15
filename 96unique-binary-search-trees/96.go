package main

func numTrees(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[0] = 1

	for i := 2; i <= n; i++ {

		for j := 0; j < i; j++ {

			dp[i] = dp[i] + dp[j-0]*dp[i-1-j]
		}

	}

	return dp[n]

}
