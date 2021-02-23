/*
 * @Description:704
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-23 15:20:25
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-23 15:33:00
 * @FilePath: \704search\704.go
 */
package leetcode

import "sort"

func search1(nums []int, target int) int {

	res := sort.SearchInts(nums, target)
	if res == len(nums) {
		return -1
	}

	if nums[res] != target {
		return -1
	}

	return res

}

func search(nums []int, target int) int {

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return -1

}
