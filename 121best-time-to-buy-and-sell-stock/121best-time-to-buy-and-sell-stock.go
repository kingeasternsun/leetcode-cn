package main

func main() {

}

func maxProfit(prices []int) int {

	pLen := len(prices)
	if pLen <= 1 {
		return 0
	}

	//后一个数值 到前一个数值的插值，转为最大连续子数组的问题
	t := make([]int, pLen-1)

	for i := 1; i < pLen; i++ {
		t[i-1] = prices[i] - prices[i-1]
	}

	dp := make([]int, pLen-1)
	dp[0] = t[0]
	maxP := dp[0]

	if pLen == 2 {
		if dp[0] > 0 {
			return dp[0]
		}

		return 0
	}

	for i := 1; i < pLen-1; i++ {
		if dp[i-1] < 0 {
			dp[i] = t[i]
		} else {
			dp[i] = dp[i-1] + t[i]
		}

		if dp[i] > maxP {
			maxP = dp[i]
		}
	}

	if maxP > 0 {
		return maxP
	}

	return 0

}
