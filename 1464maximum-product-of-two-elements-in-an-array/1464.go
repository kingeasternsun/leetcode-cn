package main

import (
	"math"
	"sort"
)

func maxProduct1(nums []int) int {

	max := math.MinInt32

	count := len(nums)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if (nums[i]-1)*(nums[j]-1) > max {
				max = (nums[i] - 1) * (nums[j] - 1)
			}
		}
	}

	return max
}
func maxProduct2(nums []int) int {

	for i := range nums {
		nums[i] = nums[i] - 1
	}

	sort.Ints(nums)

	max := nums[0] * nums[1]
	if nums[len(nums)-1]*nums[len(nums)-2] > max {
		max = nums[len(nums)-1] * nums[len(nums)-2]
	}

	return max
}

// 4ms 3MB
func maxProduct(nums []int) int {

	count := len(nums)
	if count < 2 {
		return 0
	}

	if count == 2 {
		return (nums[0] - 1) * (nums[1] - 1)
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < count-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	for i := 0; i < 2; i++ {
		for j := count - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}

	if (nums[0]-1)*(nums[1]-1) > (nums[count-1]-1)*(nums[count-2]-1) {
		return (nums[0] - 1) * (nums[1] - 1)
	}

	return (nums[count-1] - 1) * (nums[count-2] - 1)
}
