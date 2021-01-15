package leetcode

import (
	"container/heap"
	"math"
)

type MaxHeap []int

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MaxHeap) Len() int {
	return len(h)
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {

	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x

}

func minimumDeviation(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}

	lo := math.MaxInt32 //每次迭代的最小值
	for i := range nums {

		//先把奇数变为偶数
		if nums[i]&1 == 1 {
			nums[i] = nums[i] * 2
		}
		if nums[i] < lo {
			lo = nums[i]
		}

	}
	//先把奇数变为偶数后，后面就只需要考虑从偶数变为奇数的情况了

	minHeap := MaxHeap(nums)
	heap.Init(&minHeap)

	res := minHeap[0] - lo
	if res == 0 {
		return res
	}

	for minHeap[0]%2 == 0 {
		top := heap.Pop(&minHeap).(int) / 2
		heap.Push(&minHeap, top)
		lo = min(top, lo)
		res = min(res, minHeap[0]-lo)

	}

	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
