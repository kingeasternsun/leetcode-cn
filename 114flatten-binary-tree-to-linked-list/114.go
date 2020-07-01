package main

func main() {

}

//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 114. 二叉树展开为链表 https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
func flatten(root *TreeNode) {

	if root == nil {
		return
	}

	help(root, nil)

}

func help(root *TreeNode, pre *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Right != nil { //这里必须判断 ，因为如果不判断，就会在程序允许中把pre置为nil
		pre = help(root.Right, pre)
	}

	if root.Left != nil { //这里必须判断 ，因为如果不判断，就会在程序允许中把pre置为nil
		pre = help(root.Left, pre)
	}

	// fmt.Printf("%v %v\n", root, pre)

	root.Right = pre
	root.Left = nil
	pre = root

	// fmt.Printf("%v %v\n", root.Val, root.Right)

	return root

}
