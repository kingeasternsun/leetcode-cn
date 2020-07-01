package main

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {

	count := len(nums)
	if count < k {
		return 0
	}
	rand.Seed(time.Now().Unix())

	return findKHelp(nums, 0, count-1, count-k)
}

func findKHelp(nums []int, left, end, k int) int {

	parID := randomPartion(nums, left, end)
	if parID == k {
		return nums[parID]
	}

	if parID < k {
		return findKHelp(nums, parID+1, end, k)
	}

	return findKHelp(nums, left, parID-1, k)

}

// 以left 进行分割
/*
，(..., i]  内的都比 <=x ， (i, j]内的都是 > x
*/
func partion(nums []int, left, end int) int {

	i := left - 1
	x := nums[end]
	for j := left; j < end; j++ {

		//这个交换 是算法最精髓的地方
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}

	}

	// 这里 i+1 的位置上就肯定是比左边的都大，比右边的都小
	nums[i+1], nums[end] = nums[end], nums[i+1]

	return i + 1

}

func randomPartion(nums []int, left, end int) int {
	r := rand.Int()%(end-left+1) + left
	nums[r], nums[end] = nums[end], nums[r]
	return partion(nums, left, end)
}
