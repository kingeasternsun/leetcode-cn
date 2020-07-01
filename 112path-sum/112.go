package main

func main() {

}

//112. 路径总和 https://leetcode-cn.com/problems/path-sum/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return true
		}

		return false
	}

	var has bool
	if root.Left != nil {
		has = hasPathSum(root.Left, sum-root.Val)
		if has {
			return true
		}
	}

	if root.Right != nil {
		has = hasPathSum(root.Right, sum-root.Val)
		if has {
			return true
		}
	}

	return false
}

func hasPathSum1(root *TreeNode, sum int) bool {

	if root == nil {
		return false
	}

	return help(root, sum, 0)
}

func help(root *TreeNode, sum int, curSum int) bool {

	if root.Left == nil && root.Right == nil {
		if curSum+root.Val == sum {
			return true
		}

		return false
	}

	var has bool
	if root.Left != nil {
		has = help(root.Left, sum, curSum+root.Val)
		if has {
			return true
		}
	}

	if root.Right != nil {
		has = help(root.Right, sum, curSum+root.Val)
		if has {
			return true
		}
	}

	return false
}
