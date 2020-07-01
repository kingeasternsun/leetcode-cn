package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {

	if s == nil && t == nil {
		return true
	}

	if s == nil {
		return false
	}

	if t == nil {
		return true
	}

	if s.Val != s.Val {
		return isSubtree(s.Left, t) || isSubtree(s.Right, t)
	}

	if isIdentical(s, t) {
		return true
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)

}

func isIdentical(s *TreeNode, t *TreeNode) bool {

	if s == t {
		return true
	}

	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return isIdentical(s.Left, t.Left) && isIdentical(s.Right, t.Right)
}
