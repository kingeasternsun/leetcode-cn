/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-04-07 22:44:04
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-04-07 23:39:19
 * @FilePath: /81search/81.go
 */
package leetcode

import "sort"

func search(nums []int, target int) bool {

	if len(nums) == 0 {
		return false
	}

	if nums[len(nums)-1] > nums[0] {
		return biSearch(nums, target)
	}

	beg := 0
	end := len(nums) - 1
	for beg < end {

		mid := (end-beg)/2 + beg
		if nums[mid] == target || nums[beg] == target || nums[end] == target {
			return true
		}

		//mid 在左边的上升数组
		if nums[mid] > nums[beg] {
			if nums[beg] < target && target < nums[mid] {
				return biSearch(nums[beg:mid], target)
			}

			if target > nums[mid] {
				beg = mid + 1
				continue
			}

		}

		// mid 在右边上升数组
		if nums[mid] < nums[end] {
			if nums[mid] < target && target < nums[end] {

				return biSearch(nums[mid+1:end+1], target) //这里出错
			}

			if target < nums[mid] {
				end = mid - 1
				continue
			}

		}

		//都相同
		beg = beg + 1
		end = end - 1

	}
	return nums[beg] == target
}

func biSearch(nums []int, x int) bool {
	if len(nums) == 0 {
		return false
	}
	res := sort.SearchInts(nums, x)
	if res == len(nums) {
		return false
	}

	if nums[res] != x {
		return false
	}
	return true
}
