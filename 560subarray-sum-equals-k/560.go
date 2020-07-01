package main

func subarraySum(nums []int, k int) int {

	count := len(nums)
	if count == 0 {
		return 0
	}

	kNum := 0

	sumMap := make(map[int]int, 0) // 前i和为key的个数
	preSum := 0
	for i := 0; i < count; i++ {
		preSum += nums[i]

		//这个要在 sumMap[preSum]++ 之前
		kNum += (sumMap[preSum-k]) //从 sum(nums[i:j+1])==k 也即是 preSumj - preSumi == k。 只要看 前j和为 preSum-k 的有多少个

		sumMap[preSum]++

		if preSum == k {
			kNum++
		}

	}

	return kNum

}
