package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

/* 94. 二叉树的中序遍历
- 1. 一路向左入栈
- 2. 到达最左边后，出栈保存，把当前节点的右子节点作为当前节点，回到第一步

*/
func inorderTraversal1(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	sLen := 0
	result := make([]int, 0)

	cur := root
	for {
		for cur != nil {
			stack = append(stack, cur)
			sLen++
			cur = cur.Left
		}

		if sLen == 0 {
			return result
		}

		cur = stack[sLen-1]
		result = append(result, cur.Val)

		//出栈
		stack = stack[:sLen-1]
		sLen--

		// 这个节点的左子树的节点，当前节点均已经遍历，所以遍历当前节点的右边节点
		cur = cur.Right
	}

}

/* 94. 二叉树的中序遍历
- 1. 一路向左入栈
- 2. 到达最左边后，出栈保存，把当前节点的右子节点作为当前节点，回到第一步
*/
func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	cur := root

	for {

		// - 1. 一路向左入栈
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		if len(stack) == 0 {
			return result
		}

		// - 2. 到达最左边后，出栈保存，把当前节点的右子节点作为当前节点，回到第一步
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right

	}

	return result
}
