/*
 * @Description:783
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-04-13 22:19:23
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-04-13 22:52:21
 * @FilePath: /783minDiffInBST/783.go
 */
/**
 * Definition for a binary tree node.
 */
package leetcod

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {

	pre := -1
	minV := math.MaxInt32
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}

		dfs(n.Left)
		if pre != -1 && n.Val-pre < minV {
			minV = n.Val - pre
		}
		pre = n.Val
		dfs(n.Right)

	}

	dfs(root)
	return minV

}
