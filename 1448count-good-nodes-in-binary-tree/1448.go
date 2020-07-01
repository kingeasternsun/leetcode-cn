package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 // learn from huahualeecode 递归太强了
 1. 这个题目的节点 可以是树中的节点，不一定是叶子节点
*/

func goodNodes(root *TreeNode) int {
	return help(root, math.MinInt32)
}

//maxVal 表示从root到当前节点路径上（不包括当前节点）的最大值
func help(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}

	num := help(root.Left, max(maxVal, root.Val)) + help(root.Right, max(maxVal, root.Val))
	if root.Val >= maxVal { //如果当前节点的数值 大于路径上的最大值，则满足
		num++
	}
	return num
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
