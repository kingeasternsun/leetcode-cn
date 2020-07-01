package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	h := maxDepth(root.Left)
	hl := maxDepth(root.Right)
	if hl > h {
		h = hl
	}

	return h + 1

}
