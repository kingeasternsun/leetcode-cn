package main

func main() {

}

/*
121 cur可以认为理解为到第i天卖出所能获得的最大利润，也就是DP[I] 和 (DP[I-1])之间的关系
  if dp[i-1]<0 , dp[i] = prices[i]-prices[i-1]
  else  dp[i] = dp[i-1]+prices[i]-prices[i-1]

  或者 i天卖出或不卖出也不买入
  dp[i] = max(0,dp[i-1]+prices[i]-prices[i-1])
122 i-1买进，i卖出，i再买入，i+1卖出，可以等价为i-1买进，i+1卖出，所以dp可以为
	if prices[i]-proces[i-1] > 0{
 		dp[i] = dp[i-1]+prices[i]-proces[i-1]
	}else{
		dp[i] = dp[i-1]
	}

309  cool可以理解为上上一个sell，所以cool=sell没问题，不过cool=sell； 一定要在sell被重新更新之前

*/
func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// dp[i][j][0]
//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/
func maxProfit1(prices []int) int {

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
			dp[i][1] = -prices[i]
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

			if i == 1 { // 第二天拥有股票
				dp[i][1] = max(dp[i-1][1], dp[0][0]-prices[i])
			} else {
				dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
			}

		}
		// }

	}

	return dp[pLen-1][0]
}

func maxProfit(prices []int) int {

	pLen := len(prices)
	if pLen < 2 {
		return 0
	}

	var sell, buy, preSell int
	buy = -prices[0]

	for i := 1; i < pLen; i++ {
		buy = max(buy, preSell-prices[i]) //上一天buy，今天什么都没有做；前天sell，今天buy，前天sell就是preSell
		preSell = sell                    //这里sell保存的是i-1的sell，到i+1天就是preSEll
		sell = max(sell, buy+prices[i])   //上一天sell，今天什么都没有做；昨天buy，今天sell
	}

	return sell

}
