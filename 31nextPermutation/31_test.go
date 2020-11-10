package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {

	nums := []int{1, 1, 3, 3, 2, 3, 7, 9, 9, 8, 8, 3} //     ->  [1,1,3,3,2,3,8,3,7,8,9,9]
	nextPermutation(nums)                             //我的答案  [1 1 3 3 2 3 8 3 7 8 9 9]
	assert.Equal(t, nums, []int{1, 1, 3, 3, 2, 3, 8, 3, 7, 8, 9, 9}, "they should be equal")

	nums1 := []int{1, 2, 3}
	nextPermutation(nums1)
	assert.Equal(t, nums1, []int{1, 3, 2}, "they should be equal")

	nums2 := []int{3, 2, 1}
	nextPermutation(nums2)
	assert.Equal(t, nums2, []int{1, 2, 3}, "they should be equal")

	nums3 := []int{1, 1, 5}
	nextPermutation(nums3)
	assert.Equal(t, nums3, []int{1, 5, 1}, "they should be equal")

	nums4 := []int{1, 3, 2}
	nextPermutation(nums4)
	assert.Equal(t, nums4, []int{2, 1, 3}, "they should be equal")

	nums5 := []int{5, 4, 7, 5, 3, 2}
	nextPermutation(nums5)
	assert.Equal(t, nums5, []int{5, 5, 2, 3, 4, 7}, "they should be equal")
}
