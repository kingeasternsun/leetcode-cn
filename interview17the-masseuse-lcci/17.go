package main

func main() {

}

/* 跟 股票的问题很相似
dp[i][0] = max(dp[i-1][0], dp[i-1][1]) 	//第i个不预约，那么第i-1个也可以不预约，或者第i-1个预约
dp[i][1] = dp[i-1][0]+nums[i] 					//第i个预约，那么第i-1个只能不预约
*/

type Choose [2]int

// 0ms 2.1MB
func massage1(nums []int) int {
	count := len(nums)
	if count == 0 {
		return 0
	}

	dp := make([]Choose, count+1)

	for i := 1; i <= count; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i-1]
	}

	return max(dp[count][0], dp[count][1])

}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

//0ms 2MB 相比之前没有优化太多
func massage(nums []int) int {
	count := len(nums)
	if count == 0 {
		return 0
	}

	pre0 := 0
	pre1 := 0
	cur0 := 0

	for i := 0; i < count; i++ {
		cur0 = max(pre0, pre1)
		pre1 = pre0 + nums[i]
		pre0 = cur0
	}

	return max(pre0, pre1)

}
