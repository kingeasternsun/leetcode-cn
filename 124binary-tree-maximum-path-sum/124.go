// /124. 二叉树中的最大路径和 https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/

package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	_, sum := help(root)
	return sum
}

func max(i, j, k int) int {
	if j > i {
		i = j
	}
	if k > i {
		i = k
	}
	return i
}
func help(root *TreeNode) (maxLen, sum int) {
	if root == nil {
		return math.MinInt32, math.MinInt32 //这里一定要返回最小负数，而不是0
	}

	leftLen, leftSum := help(root.Left)
	rightLen, rightSum := help(root.Right)

	maxLen = root.Val + max(leftLen, rightLen, 0)
	sum = root.Val
	if leftLen > 0 {
		sum = sum + leftLen
	}
	if rightLen > 0 {
		sum = sum + rightLen
	}

	sum = max(sum, leftSum, rightSum)
	return
}
