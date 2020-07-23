// 154. 寻找旋转排序数组中的最小值 II https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
package main

// 相比153，考虑重复的情况下，nums[mid]==nums[right] , right --
func findMin(nums []int) int {
	inLen := len(nums)
	if inLen == 0 {
		return -1
	}

	if inLen == 1 {
		return nums[0]
	}

	left := 0
	right := inLen - 1

	var mid int
	// if left<=right{
	// 	return nums[left]
	// }

	for left <= right {
		if left == right {
			return nums[left]
		}

		mid = (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1 // [mid+1 right]
			continue
		}

		if nums[mid] == nums[right] {
			right = right - 1
			continue
		}

		right = mid // [beg mid]
	}

	return 0
}
