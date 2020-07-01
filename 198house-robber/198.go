package main

/*
198. 打家劫舍 https://leetcode-cn.com/problems/house-robber/
*/

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
func rob(nums []int) int {

	pre0 := 0
	pre1 := 0
	count := len(nums)
	if count == 0 {
		return 0
	}
	// maxV := 0
	for i := 0; i < count; i++ {
		tmp0 := max(pre0, pre1) //当前不rob
		pre1 = pre0 + nums[i]
		pre0 = tmp0

	}

	return max(pre0, pre1)

}
