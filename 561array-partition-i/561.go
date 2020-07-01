package main

import "sort"

func arrayPairSum(nums []int) int {

	count := len(nums)
	if count == 0 {
		return 0
	}

	if count&1 == 1 {
		return 0
	}

	sort.Ints(nums)
	sum := 0
	for i := 0; i < count; i = i + 2 {
		sum += nums[i]
	}

	return sum
}
