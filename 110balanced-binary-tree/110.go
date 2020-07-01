package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//110. 平衡二叉树 https://leetcode-cn.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	_, result := helper(root)
	return result
}
func max(i, j int) int {
	if i > j {
		return i
	}

	return j

}
func helper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	var left, right int
	var result bool
	left, result = helper(root.Left)
	if !result {
		return 0, false
	}

	right, result = helper(root.Right)
	if !result {
		return 0, false
	}

	if left-right > 1 || right-left > 1 {
		return 0, false
	}

	return max(left, right) + 1, true
}
