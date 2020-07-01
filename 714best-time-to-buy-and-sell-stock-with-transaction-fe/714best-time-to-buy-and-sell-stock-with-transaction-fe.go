package main

func main() {

}
func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// dp[i][j][0]

func maxProfit(prices []int, fee int) int {

	pLen := len(prices)
	if pLen <= 1 {
		return 0
	}

	dp := make([][2]int, pLen)
	for i := 0; i < pLen; i++ {
		dp[i] = [2]int{}
	}
	for i := 0; i < pLen; i++ {

		// for k := 1; k < maxK; k++ {

		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i] - fee
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
		}
		// }

	}

	return dp[pLen-1][0]
}
