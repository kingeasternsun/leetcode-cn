package halvearray

import "container/heap"

type MaxHeap struct {
	nums []float64
}

func (m *MaxHeap) Len() int           { return len(m.nums) }
func (m *MaxHeap) Less(i, j int) bool { return m.nums[j] < m.nums[i] } // Because we need MaxHeap
func (m *MaxHeap) Swap(i, j int)      { m.nums[i], m.nums[j] = m.nums[j], m.nums[i] }
func (m *MaxHeap) Pop() any {
	n := m.nums[m.Len()-1]
	m.nums = m.nums[:m.Len()-1]
	return n
}

func (m *MaxHeap) Push(n any) {
	m.nums = append(m.nums, n.(float64))
}

func halveArray(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	half_sum := float64(sum) / 2.0
	left_sum := float64(sum)

	ret := 0
	maxHeap := &MaxHeap{}
	for _, n := range nums {
		maxHeap.nums = append(maxHeap.nums, float64(n))
	}
	heap.Init(maxHeap)

	for {
		m := (heap.Pop(maxHeap)).(float64)
		ret += 1
		left_sum -= m / 2.0
		if left_sum <= half_sum {
			break
		}
		heap.Push(maxHeap, m/2.0)
	}

	return ret
}

func halveArrayWithFix(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	half_sum := float64(sum) / 2.0
	left_sum := float64(sum)

	ret := 0
	maxHeap := &MaxHeap{}
	for _, n := range nums {
		maxHeap.nums = append(maxHeap.nums, float64(n))
	}
	heap.Init(maxHeap)

	for {
		m := maxHeap.nums[0]
		ret += 1
		left_sum -= m / 2.0
		if left_sum <= half_sum {
			break
		}
		maxHeap.nums[0] = m / 2.0
		heap.Fix(maxHeap, 0)
	}

	return ret
}
