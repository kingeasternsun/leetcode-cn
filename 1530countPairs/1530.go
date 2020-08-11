package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {

	sum := 0
	dfs2(root, distance, &sum)
	return sum
}

/*
方法一：记录每个节点到对应的叶子节点的长度 28ms
*/
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

/*
相比方法一：记录每个节点到对应的叶子节点的长度，然后将相同长度的计数,还没有方法一快 40ms
*/
func dfs2(root *TreeNode, distance int, sum *int) (dis map[int]int) {
	dis = make(map[int]int, 0)

	if root == nil {
		return dis
	}

	if root.Left == nil && root.Right == nil {
		return map[int]int{0: 1}
	}

	left := dfs2(root.Left, distance, sum)
	right := dfs2(root.Right, distance, sum)

	// fmt.Printf("left %+v\n", left)
	// fmt.Printf("right %+v\n", right)

	for dis1, dis1c := range left {
		for dis2, dis2c := range right {
			if dis1+dis2+2 <= distance {
				*sum = *sum + dis1c*dis2c
			}
		}
	}

	for dis1, dis1c := range left {

		if dis1+1 < distance {
			dis[dis1+1] = dis[dis1+1] + dis1c
		}
	}

	for dis2, dis2c := range right {

		if dis2+1 < distance {
			// dis = append(dis, dis2+1)
			// dis[dis2+1]++
			dis[dis2+1] = dis[dis2+1] + dis2c
		}
	}

	return dis

}
