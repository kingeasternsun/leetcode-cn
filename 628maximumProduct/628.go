package leetcode

import "sort"

func maximumProduct(nums []int) int {

	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	sort.Ints(nums)

	rightMost := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	//如果大于等于0 或 都小于等于0 ，那么就肯定选最大的三个返回
	if nums[0] >= 0 && nums[len(nums)-1] <= 0 {
		return rightMost
	}

	//有正数，也有负数
	/*
		1.如果只有一个负数，那么肯定还是返回最大的三个数
		2.如果两个或两个以上的负数，那么就从左边选两个最小的负数，再选右边最大的正式
	*/

	if nums[0]*nums[1]*nums[len(nums)-1] < rightMost {
		return rightMost
	}
	return nums[0] * nums[1] * nums[len(nums)-1]
}
