package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {

	count := len(nums)
	if count == 0 {
		return nil
	}

	result := make([]string, 0)
	pre := 0

	for i := 1; i < count; i++ {
		if nums[i] == nums[i-1]+1 {
			continue
		}

		result = append(result, makeStr(nums, pre, i-1))
		pre = i

	}

	result = append(result, makeStr(nums, pre, count-1))
	return result

}

func makeStr(nums []int, i, j int) string {
	if i == j {
		return strconv.Itoa(nums[i])
	}

	return fmt.Sprintf("%v->%v", nums[i], nums[j])
}
