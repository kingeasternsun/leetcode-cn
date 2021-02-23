/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2020-07-01 12:44:30
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-23 16:01:03
 * @FilePath: \34\34.go
 */
package leetcode

import "sort"

func searchRange(nums []int, target int) []int {

	left := sort.SearchInts(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	right := sort.SearchInts(nums, target+1)
	return []int{left, right - 1}

}

func searchRange_old(nums []int, target int) []int {

	left := binSearch(nums, target, true)
	if left == -1 {
		return []int{-1, -1}
	}

	right := binSearch(nums, target, false)

	return []int{left, right}
}

func binSearch(nums []int, target int, getLeft bool) int {

	sLen := len(nums)
	right := len(nums) - 1
	left := 0

	for left < sLen && right >= 0 && left <= right {

		//只有一个处理
		if left == right {
			if nums[left] == target {
				return left
			}
			return -1
		}

		// fmt.Println(left,right,nums)
		// 对区间内有两个的特殊处理
		if left == right-1 {
			if nums[left] != target && nums[right] != target {
				return -1
			}

			if nums[left] == target && nums[right] == target {
				if getLeft {
					return left
				}
				return right
			}

			if nums[left] == target {
				return left
			}

			if nums[right] == target {
				return right
			}
		}

		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {

			if getLeft {
				right = mid
			} else {
				left = mid
			}

		}

	}

	return -1

}
