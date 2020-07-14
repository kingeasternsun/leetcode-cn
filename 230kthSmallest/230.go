package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {

	if root == nil {
		return 0
	}

	cur := root
	stack := make([]*TreeNode, 0)
	for {

		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		if len(stack) == 0 {
			return 0
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k = k - 1
		if k == 0 {
			return cur.Val
		}

		cur = cur.Right

	}

}
