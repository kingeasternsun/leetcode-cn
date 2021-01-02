package leetcode

//https://github.com/kingeasternsun/leetcode-cn

import "math"

//dp
func maxResult_timeout(nums []int, k int) int {

	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)) //表示跳在第i个位置的最大得分
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {

		beg := i - k
		if beg < 0 {
			beg = 0
		}

		dp[i] = maxOfSlice(dp[beg:i]) + nums[i]

	}
	return dp[len(nums)-1]

}

func maxOfSlice(dp []int) int {

	res := math.MinInt32
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	return res
}

func maxResult(nums []int, k int) int {

	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)) //表示跳在第i个位置的最大得分
	dp[0] = nums[0]
	monoQueue := []int{nums[0]} //基于dp的滑窗

	for i := 1; i < len(nums); i++ {

		// dp[i] = maxOfSlice(dp[beg:i]) + nums[i]
		//类似 239题目，利用单调队列
		tmpV := monoQueue[0] + nums[i]
		dp[i] = tmpV

		//处理滑窗
		for len(monoQueue) > 0 {

			//tmpV比滑窗最后边的要大，就把最右边的剔除
			if tmpV > monoQueue[len(monoQueue)-1] {
				monoQueue = monoQueue[:len(monoQueue)-1]
			} else {
				break
			}

		}

		monoQueue = append(monoQueue, tmpV)
		// 如果当前sliding的大小超过了k，就需要从dp中剔除左边的一个,时刻让sliding窗口的大小最大为k,
		// 所以每次只需要剔除滑窗左边的一个(也就是i-k位置上的那个)就可以了
		if i >= k {
			//如果dp[i-k]刚好是最大值 ,monoQueue也需要移除左边的最大值
			if dp[i-k] == monoQueue[0] {
				monoQueue = monoQueue[1:]
			}
		}

	}
	return dp[len(nums)-1]

}
