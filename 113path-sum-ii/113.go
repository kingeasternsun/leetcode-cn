package main

//113. 路径总和 II https://leetcode-cn.com/problems/path-sum-ii/submissions/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {

	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	curArray := make([]int, 0)
	help(root, sum, 0, 0, curArray, &result)
	return result
}

/*
curArry 已经访问的节点列表
curNum  当前curArry的长度
*/
func help(root *TreeNode, sum int, curSum int, curNum int, curArry []int, result *[][]int) {

	curArry = append(curArry, root.Val)
	curSum += root.Val
	curNum++

	if root.Left == nil && root.Right == nil {
		if curSum == sum {
			tmp := make([]int, curNum)
			copy(tmp, curArry)
			*result = append(*result, tmp)
			return
		}

		return
	}

	// var has bool
	if root.Left != nil {
		help(root.Left, sum, curSum, curNum, curArry, result)

	}

	if root.Right != nil {
		help(root.Right, sum, curSum, curNum, curArry, result)
	}

	return
}
