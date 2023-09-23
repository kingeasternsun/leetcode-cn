package schedulercourses

import (
	"container/heap"
	"sort"
)

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	sum := 0
	cnt := 0
	myHeap := IntHeap{}

	heap.Push(&myHeap, 0)

	for _, cs := range courses {
		// 可以新增一个课程
		if sum+cs[0] <= cs[1] {
			heap.Push(&myHeap, cs[0])
			cnt += 1
			sum += cs[0]
		} else {
			// 当前课程数量无法增加，就要想办法减少sum: 用这个cs替换之前的一个课程
			if myHeap[0] > cs[0] {

				sum -= myHeap[0]
				sum += cs[0]

				myHeap[0] = cs[0]
				heap.Fix(&myHeap, 0)
			}

		}

	}

	return cnt

}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	queueMap := make([][]int, 8)
	for i := 0; i < 8; i++ {
		queueMap[i] = make([]int, 8)
	}

	for _, q := range queens {
		queueMap[q[0]][q[1]] = 1
	}

	ret := make([][]int, 0)

	dirs := [][]int{{0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}, {-1, 0}, {1, 0}}
	for _, dir := range dirs {
		pos := append([]int{}, king...)
		for pos[0] >= 0 && pos[0] < 8 && pos[1] >= 0 && pos[1] < 8 {

			if queueMap[pos[0]][pos[1]] == 1 {
				ret = append(ret, pos)
				break
			}

			pos[0] += dir[0]
			pos[1] += dir[1]
		}

	}

	return ret

}

// 1026

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	ret, _, _ := minMaxAncestorDiff(root)
	return ret
}

func minMaxAncestorDiff(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, -1, -1
	}

	leftV, leftMin, leftMax := minMaxAncestorDiff(root.Left)
	rightV, rightMin, rightMax := minMaxAncestorDiff(root.Right)

	cur := max(leftV, rightV)
	if leftMin >= 0 {
		cur = max(cur, abs(root.Val-leftMin))
	}

	if leftMax >= 0 {
		cur = max(cur, abs(root.Val-leftMax))
	}

	if rightMin >= 0 {
		cur = max(cur, abs(root.Val-rightMin))
	}

	if rightMax >= 0 {
		cur = max(cur, abs(root.Val-rightMax))
	}

	curMin := root.Val
	curMax := root.Val

	if leftMin >= 0 {
		curMin = min(curMin, leftMin)
	}
	if rightMin >= 0 {
		curMin = min(curMin, rightMin)
	}

	if leftMax >= 0 {
		curMax = max(curMax, leftMax)
	}

	if rightMax >= 0 {
		curMax = max(curMax, rightMax)
	}

	return cur, curMin, curMax

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}
