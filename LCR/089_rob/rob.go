package rob

func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		dp[i] = dp[i-1]
		if i >= 2 && nums[i]+dp[i-2] > dp[i] {
			dp[i] = nums[i] + dp[i-2]
		}
	}
	return dp[len(nums)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
