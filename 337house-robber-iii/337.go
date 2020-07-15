package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left, right := rob_help(root)
	return max(left, right)

}

func rob_help(root *TreeNode) (int, int) {

	if root == nil {
		return 0, 0
	}

	left0, left1 := rob_help(root.Left)
	right0, right1 := rob_help(root.Right)
	//抢劫当前节点，那么左右节点就不抢劫
	tmp1 := left0 + right0 + root.Val

	// 当前不抢劫 ,左右子节点可以抢劫 也可以不抢劫
	tmp0 := max(left0, left1) + max(right0, right1)
	return tmp0, tmp1
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
