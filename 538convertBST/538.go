package main

func main() {

}

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {

	var curSum = 0
	sum(root, &curSum)
	return root
}

func sum(root *TreeNode, curSum *int) {

	if root == nil {
		return
	}

	sum(root.Right, curSum)
	*curSum = *curSum + root.Val
	root.Val = *curSum
	sum(root.Left, curSum)
	// return root.Val + sum(root.Left)
}
