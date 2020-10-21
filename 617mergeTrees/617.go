package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	left := mergeTrees(t1.Left, t2.Left)
	right := mergeTrees(t1.Right, t2.Right)
	t1.Val += t2.Val
	t1.Left = left
	t1.Right = right

	return t1
}
