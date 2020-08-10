package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {

	sum := 0
	dfs(root, distance, &sum)
	return sum
}

func dfs(root *TreeNode, distance int, sum *int) (dis []int) {

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{0}
	}

	left := dfs(root.Left, distance, sum)
	right := dfs(root.Right, distance, sum)

	for _, dis1 := range left {
		for _, dis2 := range right {
			if dis1+dis2+2 <= distance {
				*sum = *sum + 1
			}
		}
	}

	for _, dis1 := range left {

		if dis1+1 < distance {
			dis = append(dis, dis1+1)
		}
	}

	for _, dis2 := range right {

		if dis2+1 < distance {
			dis = append(dis, dis2+1)
		}
	}

	return

}
