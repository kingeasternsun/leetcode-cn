package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root.Left, root.Right)

}

func helper(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	left := helper(p.Left, q.Right)
	if !left {
		return false
	}

	right := helper(p.Right, q.Left)
	if !right {
		return false
	}

	return true

}
