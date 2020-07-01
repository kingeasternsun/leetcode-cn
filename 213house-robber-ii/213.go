package main

/*
213. 打家劫舍 II https://leetcode-cn.com/problems/house-robber-ii/
相似 198. 打家劫舍 https://leetcode-cn.com/problems/house-robber/
既然构成环，那么首尾相接的地方至少有一个不rob，那么我们就把第一个去掉按照198进行计算，最后一个去掉按照198计算，求最大值
*/

func rob(nums []int) int {

	count := len(nums)
	if count == 0 {
		return 0
	}

	if count == 1 {
		return nums[0]
	}

	return max(rob1(nums[1:]), rob1(nums[:count-1]))
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
func rob1(nums []int) int {

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
