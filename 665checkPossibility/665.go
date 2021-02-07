/*
 * @Description: 665
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-07 09:18:33
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-07 11:46:28
 * @FilePath: \665checkPossibility\665.go
 */
package leetcode

import "sort"

/*
1300ms 6.5MB
*/
func checkPossibility1(nums []int) bool {

	if len(nums) <= 2 {
		return true
	}

	//计算最长的递增子序列长度
	dp := make([]int, len(nums)) //以当前元素为结尾的自增序列长度
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		tmp := 1
		for j := 0; j < i; j++ {
			if nums[i] >= nums[j] && dp[j]+1 > tmp {
				tmp = dp[j] + 1
			}
		}
		dp[i] = tmp
	}

	cnt := 1
	for i := range dp {
		if dp[i] > cnt {
			cnt = dp[i]
		}
	}
	return cnt >= (len(nums) - 1)
}

/*
 二分 贪心法 转为求最长非连续递增子序列问题 44ms 6.7MB
[1,2,4]  new 3 ,fount 3, get  [1,2,3]  //单调递增
[1,2,4]  new 3 ,found 4, get  [1,2,3]  //递增

[1,2,4]  new 2, found 2, get  [1,2,4]  //单调自增
[1,2,4]  new 2, found 3, get  [1,2,2]  //递增
*/
func checkPossibility2(nums []int) bool {

	if len(nums) <= 2 {
		return true
	}

	//计算最长的递增子序列长度
	increNums := []int{nums[0]} //维护当前可以构建的递增序列
	for i := 1; i < len(nums); i++ {

		//因为是可以有重复，所以这里查询的时候要+1
		pos := sort.SearchInts(increNums, nums[i]+1)
		if pos == len(increNums) {
			increNums = append(increNums, nums[i])
		} else {

			increNums[pos] = nums[i]
		}

	}

	return len(increNums) >= (len(nums) - 1)

}

/*
参考 题解 https://leetcode-cn.com/problems/non-decreasing-array/solution/3-zhang-dong-tu-bang-zhu-ni-li-jie-zhe-d-06gi/
[1,4,2]-> [1,2,3]  nums[2]<nums[i-1] && nums[i]>=nums[i-2] 就让 nums[i-1] = nums[i]
[3,4,2]-> [3,4,4]  nums[2]<nums[i-1] && nums[i]<nums[i-2]  就让 nums[i-1] = nums[i]
本质也是贪心，维护最小的递增序列
*/
func checkPossibility(nums []int) bool {

	if len(nums) <= 2 {
		return true
	}

	changeCnt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[i-1] {
			continue
		}

		changeCnt++
		//如果一开始就出现了乱序
		if i == 1 {
			//把 nums[i-1] 降低数值到 nums[i] 。贪心思路
			nums[i-1] = nums[i]
			continue
		}

		if changeCnt >= 2 {
			return false
		}
		//维护 nums[i-2] nums[i-1] nums[i]递增
		if nums[i] >= nums[i-2] {
			//贪心思路 ,把 nums[i-1] 降低数值到 nums[i],这样递增序列的最大值最小
			nums[i-1] = nums[i]
		} else {
			nums[i] = nums[i-1]
		}
	}

	return true

}
