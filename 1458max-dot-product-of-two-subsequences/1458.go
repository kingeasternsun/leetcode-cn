package main

import "math"

/*

由于可能出现负值，所以会比较麻烦些，
 1. dp默认初始化的值 math.MinInt32
 2. 每次迭代的时候，  	dp[i&1][0] = math.MinInt32  要设置为最小数值
 3. 对角线的 ，要判断 dp[preID][j-1]是否小于0
*/
func maxDotProduct(nums1 []int, nums2 []int) int {

	count1 := len(nums1)
	count2 := len(nums2)
	if count1 == 0 || count2 == 0 {
		return 0
	}

	dp := make([][]int, 2)
	dp[0] = make([]int, count2+1)
	dp[1] = make([]int, count2+1)

	for j := 0; j <= count2; j++ {
		dp[1][j] = math.MinInt32
	}

	for i := 0; i < count1; i++ {
		preID := (i + 1) & 1
		dp[i&1][0] = math.MinInt32

		for j := 1; j <= count2; j++ {

			preValue := dp[preID][j-1]
			if preValue < 0 {
				preValue = 0
			}

			// if i == 0 {
			// 	dp[i&1][j] = max(dp[i&1][j-1], dp[preID][j-1]+nums1[i]*nums2[j-1])

			// } else {
			dp[i&1][j] = max(max(dp[preID][j], dp[i&1][j-1]), preValue+nums1[i]*nums2[j-1])

		}

	}

	return dp[(count1-1)&1][count2]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
