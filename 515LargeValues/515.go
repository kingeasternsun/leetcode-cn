package main

import "math"

func main() {

}

//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {

	//BFS

	if root == nil {
		return nil
	}
	res := make([]int, 0)
	cur := make([]*TreeNode, 0)
	cur = append(cur, root)
	for len(cur) > 0 {
		newCur := make([]*TreeNode, 0)

		maxValue := math.MinInt32
		for _, c := range cur {
			if c.Val > maxValue {
				maxValue = c.Val
			}

			if c.Left != nil {
				newCur = append(newCur, c.Left)
			}
			if c.Right != nil {
				newCur = append(newCur, c.Right)
			}

		}
		res = append(res, maxValue)
		cur = newCur

	}

	return res
}
