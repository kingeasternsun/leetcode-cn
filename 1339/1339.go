package main

import "fmt"

// 1339. 分裂二叉树的最大乘积

func main() {
	fmt.Println(10 ^ 9 + 7)
	fmt.Println(1e9 + 7)
	fmt.Println(1000000007)
	fmt.Println(6043763521071 % 1000000007)
	// fmt.Println(6043763521071 % (1e9 + 7))
}

//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var sproduct int

	// s := 0

	gSum := sum(root, 0, &sproduct)
	sum(root, gSum, &sproduct)

	var kMod int = 1000000007

	return (sproduct % kMod)
}

func max(args ...int) int {

	maxVal := args[0]
	for i := range args {
		if args[i] > maxVal {
			maxVal = args[i]
		}
	}

	return maxVal

}

func sum(root *TreeNode, gSum int, p *int) int {

	if root == nil {
		return 0
	}

	l := sum(root.Left, gSum, p)
	r := sum(root.Right, gSum, p)
	// fmt.Println(l, r)

	if gSum > 0 {
		*p = max(*p, (gSum-l)*l, (gSum-r)*r)
	}

	return (l + r + (root.Val))

}
