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

func maxProfit(k int, prices []int) int {

	pLen := len(prices)
	if pLen <= 1 {
		return 0
	}

	maxK := k + 1

	//关键点 要限制k的大小
	if k > pLen/2 {
		maxK = pLen/2 + 1
	}

	dp := make([][][2]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][2]int, maxK)
	}
	for i := 0; i < pLen; i++ {

		row := i % 2
		preRow := (i + 1) % 2
		for k := 1; k < maxK; k++ {

			if i == 0 {
				dp[row][k][0] = 0
				dp[row][k][1] = -prices[i]
			} else {
				dp[row][k][0] = max(dp[preRow][k][0], dp[preRow][k][1]+prices[i])
				dp[row][k][1] = max(dp[preRow][k][1], dp[preRow][k-1][0]-prices[i])
			}
		}

	}

	return max(dp[0][maxK-1][0], dp[1][maxK-1][0])
}
