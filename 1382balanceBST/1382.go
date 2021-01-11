package leetcode

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	var list []*TreeNode
	inorder(root, &list)
	return toBST(list)

}

//中序遍历变为数组
func inorder(root *TreeNode, list *[]*TreeNode) {

	if root == nil {
		return
	}

	inorder(root.Left, list)
	*list = append(*list, root)
	inorder(root.Right, list)

}

// 数组变为BST
func toBST(list []*TreeNode) (root *TreeNode) {
	if len(list) == 0 {
		return nil
	}
	mid := len(list) / 2
	left := toBST(list[:mid])
	right := toBST(list[mid+1:])
	root = list[mid]
	root.Left = left
	root.Right = right
	return

}
