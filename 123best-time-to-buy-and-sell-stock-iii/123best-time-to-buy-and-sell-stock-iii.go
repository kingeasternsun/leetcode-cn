package leetcode

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

//dp[i][k][0] 表示在第i天，已经执行了k次买入动作，今天是要卖掉
//dp[i][k][1] 表示在第i天，已经执行了k次买入动作，今天仍然持有
func maxProfit(prices []int) int {

	const maxK = 3
	pLen := len(prices)
	if pLen <= 1 {
		return 0
	}

	dp := make([][maxK][2]int, pLen)
	for i := 0; i < pLen; i++ {
		dp[i] = [maxK][2]int{}
	}
	// 对于每一天，如果k是0，也就是说没有执行过买入动作，获取的收益肯定就是0，也就是 dp[i][0][0] dp[i][0][1]都是0没问题

	for i := 0; i < pLen; i++ {

		for k := 1; k < maxK; k++ {

			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
			} else {
				//第i天，已经执行了k次买入，当前没有股票；那么可能是在i-1天已经执行了k次买入，手上没有股票 或 是在i-1天已经执行k次买入手上有股票，然后在i天卖出了
				dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
				//第i天，已经执行了k次买入，  当前有股票；那么可能是在i-1天已经执行了k次买入，  手上有股票 或 是在i-1天已经执行k次买入手上没有股票，然后在i天买入了
				dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
			}
		}

	}

	return dp[pLen-1][maxK-1][0]
}
