/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-11 23:26:03
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-12 00:13:53
 * @FilePath: /703KthLargest/703.go
 */
package leetcode

import "container/heap"

//求最大的K个，使用最小堆
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	Nums IntHeap
	K    int
}

func Constructor(k int, nums []int) KthLargest {
	if len(nums) <= k {
		h := IntHeap(nums)
		heap.Init(&h)
		return KthLargest{Nums: h, K: k}
	}

	h := IntHeap(nums[:k])
	heap.Init(&h)
	for _, v := range nums[k:] {

		if v > h[0] {
			heap.Pop(&h)
			heap.Push(&h, v)
		}

	}
	return KthLargest{Nums: h, K: k}

}

func (this *KthLargest) Add(val int) int {
	if len(this.Nums) == 0 || len(this.Nums) < this.K {
		heap.Push(&this.Nums, val)
	} else if val > this.Nums[0] {
		heap.Pop(&this.Nums)
		heap.Push(&this.Nums, val)
	}
	return this.Nums[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
