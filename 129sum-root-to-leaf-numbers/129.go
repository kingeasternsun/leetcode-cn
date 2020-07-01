package main

func main() {

}

/*
 129. 求根到叶子节点数字之和 https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/

 dfs 解决，
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers1(root *TreeNode) int {

	if root == nil {
		return 0
	}

	// curSum := 0
	sum := 0
	help1(root, root.Val, &sum)
	return sum
}

// curNum表示从根节点到当前节点cur所组成路径的数字
func help1(cur *TreeNode, curNum int, sum *int) {

	// if cur == nil {
	// 	*sum = curNum
	// 	return
	// }

	// curNum = curNum*10 + cur.Val
	if cur.Left == nil && cur.Right == nil { //如果当前时叶子节点
		*sum = *sum + curNum
		return
	}

	if cur.Left != nil {
		help1(cur.Left, curNum*10+cur.Left.Val, sum)
	}

	if cur.Right != nil {
		help1(cur.Right, curNum*10+cur.Right.Val, sum)
	}

}

/**************两外一种dfs思路*********************/
func sumNumbers(root *TreeNode) int {

	if root == nil {
		return 0
	}

	// curSum := 0
	sum := 0
	help(root, 0, &sum)
	return sum
}

// curNum表示从根节点到当前节点cur 的父节点 所组成路径的数字,不包含当前节点
func help(cur *TreeNode, curNum int, sum *int) {

	// if cur == nil {
	// 	*sum = curNum
	// 	return
	// }

	curNum = curNum*10 + cur.Val
	if cur.Left == nil && cur.Right == nil { //如果当前时叶子节点
		*sum = *sum + curNum
		return
	}

	if cur.Left != nil {
		help(cur.Left, curNum, sum)
	}

	if cur.Right != nil {
		help(cur.Right, curNum, sum)
	}

}
