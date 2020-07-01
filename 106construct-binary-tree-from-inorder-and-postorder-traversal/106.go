package main

func main() {

}

//https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
//106. 从中序与后序遍历序列构造二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	pLen := len(postorder)
	iLen := len(inorder)
	if pLen == 0 {
		return nil
	}

	if pLen != iLen {
		return nil
	}

	root, result := helpBuild(postorder, pLen, inorder)
	if result < 0 {
		return nil
	}

	return root

}

func helpBuild(postorder []int, pLen int, inorder []int) (*TreeNode, int) {

	if postorder == nil || pLen == 0 {
		return nil, 0
	}

	if pLen == 1 {
		return &TreeNode{Val: postorder[0]}, 0
	}

	root := &TreeNode{Val: postorder[pLen-1]}

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

	root.Left, result = helpBuild(postorder[:mid], mid, inorder[:mid]) //左子树
	if result < 0 {
		return nil, -1
	}
	root.Right, result = helpBuild(postorder[mid:pLen-1], pLen-mid-1, inorder[mid+1:]) //右子树
	if result < 0 {
		return nil, -1
	}
	return root, 0
}
