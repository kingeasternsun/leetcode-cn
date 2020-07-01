package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从前序与中序遍历序列构造二叉树
//https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
func buildTree(preorder []int, inorder []int) *TreeNode {
	pLen := len(preorder)
	iLen := len(inorder)
	if pLen == 0 {
		return nil
	}

	if pLen != iLen {
		return nil
	}

	root, result := helpBuild(preorder, pLen, inorder)
	if result < 0 {
		return nil
	}

	return root

}

func helpBuild(preorder []int, pLen int, inorder []int) (*TreeNode, int) {

	if preorder == nil || pLen == 0 {
		return nil, 0
	}

	if pLen == 1 {
		return &TreeNode{Val: preorder[0]}, 0
	}

	root := &TreeNode{Val: preorder[0]}

	mid := 0
	for ; mid < pLen; mid++ {
		if inorder[mid] == root.Val {
			break
		}
	}

	if mid == pLen {
		return nil, -1
	}

	var result int

	root.Left, result = helpBuild(preorder[1:mid+1], mid, inorder[:mid]) //左子树
	if result < 0 {
		return nil, -1
	}
	root.Right, result = helpBuild(preorder[mid+1:], pLen-mid-1, inorder[mid+1:]) //右子树
	if result < 0 {
		return nil, -1
	}
	return root, 0
}
