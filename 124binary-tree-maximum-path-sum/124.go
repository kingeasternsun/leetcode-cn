// /124. 二叉树中的最大路径和 https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/

package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	_, curSum := help(root)
	return curSum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 返回当前节点对应子树上，以当前节点为起点的子路径的最大和， 以及当前节点为根节点的子树的上的最大路径和
func help(cur *TreeNode) (int, int) {

	if cur == nil {
		return math.MinInt32, math.MinInt32 //必须返回最小值
	}

	leftSingleSum, leftSum := help(cur.Left)
	rightSingleSum, rightSum := help(cur.Right)

	curSum := cur.Val + max(leftSingleSum, 0) + max(rightSingleSum, 0) //经过cur节点的最大路径和

	curSingleSum := cur.Val + max(max(leftSingleSum, rightSingleSum), 0) //以当前节点为起点的子路径的最大和

	return curSingleSum, max(max(leftSum, rightSum), curSum)

}
