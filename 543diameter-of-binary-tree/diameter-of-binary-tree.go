package main

func main() {

}

//https://leetcode-cn.com/problems/diameter-of-binary-tree/
// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}

	d, _ := helper(root)
	return d
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/*

返回当前节点的最大直径，以及当前节点的层高

*/
func helper(root *TreeNode) (diameter int, height int) {
	if root == nil {
		return -1, -1
	}

	var (
		leftD, leftH   int
		rightD, rightH int
	)

	leftD, leftH = helper(root.Left)
	rightD, rightH = helper(root.Right)
	height = max(leftH, rightH) + 1
	diameter = leftH + rightH + 2
	diameter = max(max(diameter, leftD), rightD)

	return
}
