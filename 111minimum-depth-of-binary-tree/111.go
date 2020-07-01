package main

//https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
//111. 二叉树的最小深度

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先
func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	qLen := 1
	step := 0

	for {
		if qLen == 0 {
			return step
		}

		step++
		newQlen := 0

		for i := 0; i < qLen; i++ {
			if queue[i].Left == nil && queue[i].Right == nil {
				return step
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
				newQlen++
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
				newQlen++
			}
		}

		queue = queue[qLen:]
		qLen = newQlen

	}

	return qLen

}
