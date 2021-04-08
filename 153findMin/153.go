package main

//https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
/*
 - 如果mid的数值小于right的数值，说明mid在低数值的数据范围内，mid的左边可能有更小的。所以在 [left，mid] 里面找
 - 如果mid的数值大于right的数值，说明mid在高数值的数据范围内，更小的在右边。所以在 (mid,right] 里面找
*/
//153. 寻找旋转排序数组中的最小值 0ms 2.2MB
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

		right = mid // [beg mid]
	}

	return 0

}
