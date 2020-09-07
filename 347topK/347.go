package main

import "container/heap"

func main() {

}

type Item struct {
	Num int
	Cnt int
}

// An IntHeap is a min-heap of ints.
type IntHeap []Item

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].Cnt < h[j].Cnt }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Item))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v] = m[v] + 1
	}

	h := &IntHeap{}
	heap.Init(h)
	for num, cnt := range m {

		if h.Len() < k {
			heap.Push(h, Item{Num: num, Cnt: cnt})
		} else {
			if cnt > (*h)[0].Cnt {
				heap.Pop(h)
				heap.Push(h, Item{Num: num, Cnt: cnt})
			}
		}
	}

	var result []int
	for _, v := range *h {
		result = append(result, v.Num)
	}

	return result
}
