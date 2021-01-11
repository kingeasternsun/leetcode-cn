package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//执行中序遍历，判断是否比上一个大
func isValidBST(root *TreeNode) bool {

	var pre *TreeNode

	return help(root, &pre)
}

func help(cur *TreeNode, pre **TreeNode) bool {

	if cur == nil {
		return true
	}

	if !help(cur.Left, pre) {
		return false
	}

	if *pre != nil && cur.Val <= (*pre).Val {
		return false
	}

	*pre = cur

	if !help(cur.Right, pre) {
		return false
	}

	return true
}
