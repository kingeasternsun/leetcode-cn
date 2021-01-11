package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 其实提醒可以理解为，对bst进行一个排序，然后进行累加。 我们可以执行中序排列，先遍历右树，当前节点，左树
func bstToGst(root *TreeNode) *TreeNode {
	curSum := 0
	help(root, &curSum)

	return root

}

//curSum 表示 遍历到cur之前已经得到的累加和
func help(cur *TreeNode, curSum *int) {

	if cur == nil {
		return
	}

	help(cur.Right, curSum)
	*curSum += cur.Val
	cur.Val = *curSum
	help(cur.Left, curSum)

}
