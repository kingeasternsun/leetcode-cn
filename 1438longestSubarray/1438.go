/*
 * @Description:1438
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-22 09:12:51
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-22 09:49:28
 * @FilePath: \1438longestSubarray\1438.go

 1438. 绝对差不超过限制的最长连续子数组
给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。

如果不存在满足条件的子数组，则返回 0 。

https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
*/

package leetcode

func longestSubarray(nums []int, limit int) int {
	if len(nums) == 0 {
		return 0
	}

	moreThanNums := []int{nums[0]} //
	lessThanNums := []int{nums[0]}

	beg := 0
	end := 0
	res := 1

	for end < len(nums) && beg <= end {
		if (moreThanNums[0] - lessThanNums[0]) <= limit {
			if end-beg+1 > res {
				res = end - beg + 1
			}

			end++
			if end == len(nums) {
				break
			}

			for len(moreThanNums) > 0 && moreThanNums[len(moreThanNums)-1] < nums[end] {
				moreThanNums = moreThanNums[0 : len(moreThanNums)-1]
			}
			moreThanNums = append(moreThanNums, nums[end])

			for len(lessThanNums) > 0 && lessThanNums[len(lessThanNums)-1] > nums[end] {
				lessThanNums = lessThanNums[0 : len(lessThanNums)-1]
			}
			lessThanNums = append(lessThanNums, nums[end])

		} else {

			if nums[beg] >= moreThanNums[0] {
				moreThanNums = moreThanNums[1:]
			}

			if nums[beg] <= lessThanNums[0] {
				lessThanNums = lessThanNums[1:]
			}

			beg++
		}

	}

	return res

}
