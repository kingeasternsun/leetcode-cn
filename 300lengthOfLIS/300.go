package leetcode

import "sort"

//300. 最长上升子序列  https://leetcode-cn.com/problems/longest-increasing-subsequence/

/*
  n*n 的做法， dp【i】记录以i为结尾的最大上升字串
*/
func lengthOfLIS1(nums []int) int {

	count := len(nums)
	if count <= 1 {
		return count
	}

	dp := make([]int, count)
	for i := 0; i < count; i++ {
		dp[i] = 1
	}

	maxLen := 1
	for i := 1; i < count; i++ {
		// d[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			} else if nums[i] == nums[j] && dp[j] > dp[i] {
				dp[i] = dp[j]
			}

		}

		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen

}

//https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/zui-chang-shang-sheng-zi-xu-lie-dong-tai-gui-hua-2/
/*
n*logn
dp【i】记录的是 长度为i+1的上升子串的最小的最后一个数字，使用贪心算法，在每次迭代中，维护这个数字最小
并且保持dp的值是单调递增的，可用反证法证明
*/
func lengthOfLIS(nums []int) int {

	count := len(nums)
	if count <= 1 {
		return count
	}
	dp := make([]int, len(nums))
	//dp初始化
	maxLen := 1     //当前已经知道的最长递增数字串 为1
	dp[0] = nums[0] //当前长度为1的递增数字串的最后一个数字是nums[0]

	//从第二个数字开始
	for _, v := range nums[1:] {

		//因为是严格递增序列 在dp[:maxLen]里面查询 子序列中最后一个数字大于等于v 且子串长度最短的dp进行替换，保证 dp[i]中记录的是最小的数字
		// tmpID := biSearch(dp[0:maxLen], v)
		tmpID := sort.SearchInts(dp[0:maxLen], v)
		if tmpID == maxLen { //v 比 dp里面的都大，就可以新开一个子序列
			dp[maxLen] = v
			maxLen++
		} else {
			dp[tmpID] = v
		}

	}

	return maxLen

}

//在dp中查找大于等于v的位置的最小索引
func biSearch(dp []int, v int) int {

	left := 0
	right := len(dp) //这里巧妙的在于 right 用 len(dp)而不是len(dp)-1，因为每次for循环里面是直接取平均值的，所以没有问题，而且可以处理v比dp里面都大的情况

	for left < right {

		mid := (left + right) / 2
		if dp[mid] >= v {
			right = mid
		} else {
			left = mid + 1
		}

	}
	return left

}
