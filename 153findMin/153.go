/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-01-24 16:17:56
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-04-08 22:50:46
 * @FilePath: /153findMin/153.go
 */

package main

//https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
/*
 - 如果mid的数值小于right的数值，说明mid在低数值的数据范围内，mid的左边可能有更小的。所以在 [left，mid] 里面找
 - 如果mid的数值大于right的数值，说明mid在高数值的数据范围内，更小的在右边。所以在 (mid,right] 里面找
*/
//153. 寻找旋转排序数组中的最小值 0ms 2.2MB
func findMin1(nums []int) int {
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

//0ms
func findMin2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	//如果最后一个比第一个要大，说明是递增的
	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}

	beg := 0
	end := len(nums) - 1
	for beg < end {

		if nums[beg] < nums[end] {
			return nums[beg]
		}

		mid := (end-beg)/2 + beg
		// 在左边的递增数组上,这里最关键，要从mid  和 end上的值进行比较，因为 mid 有可能就是 beg
		// 特别容易错的地方就在于 ben = end-1 的情况时候
		if nums[mid] > nums[end] {
			beg = mid + 1
			continue
		}

		end = mid

	}

	return nums[beg]

}

//4ms
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	//如果最后一个比第一个要大，说明是递增的
	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}

	beg := 0
	end := len(nums) - 1
	for beg < end {

		if nums[beg] < nums[end] {
			return nums[beg]
		}

		mid := (end-beg)/2 + beg
		if nums[mid] > nums[end] {
			beg = mid + 1
			continue
		}

		if mid == 0 || nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		end = mid - 1

	}

	return nums[beg]

}
