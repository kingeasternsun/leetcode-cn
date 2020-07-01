package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1. DFS
2. 记录每个数字出现的个数，最后计算 出现个数是奇数的数字的个数，大于1个就不满足条件
3. 记得回退的时候 减去出现次数
148ms 19.6MB
*/
func pseudoPalindromicPaths(root *TreeNode) int {

	if root == nil {
		return 0
	}

	num := 0

	count := make([]int, 10)

	help(root, count, &num)
	return num
}

func help(root *TreeNode, count []int, num *int) {

	count[root.Val] = count[root.Val] + 1
	if root.Left == nil && root.Right == nil {
		*num = *num + check(count)
		count[root.Val] = count[root.Val] - 1
		return
	}

	if root.Left != nil {
		help(root.Left, count, num)
	}

	if root.Right != nil {
		help(root.Right, count, num)
	}

	count[root.Val] = count[root.Val] - 1
}

func check(count []int) int {

	oddNum := 0

	for i := range count {
		if count[i]&1 == 1 {
			oddNum++
		}
	}

	if oddNum >= 2 {
		return 0
	}

	return 1
}
