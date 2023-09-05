package rob

// 拆分为两个子问题即可：一排房子不包含第一个房子，一排房子不包含最后一个房子
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob1(nums[1:]), rob1(nums[:len(nums)-1]))
}

func rob1(nums []int) int {

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
	if i < j {
		return j
	}
	return i
}
