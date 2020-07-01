package main

//https://leetcode-cn.com/problems/binary-tree-right-side-view/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {

	// var result []int

	return bfs(root)
}

//0ms 2.2MB
func bfs(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	// result = append(result, root.Val)
	cur := []*TreeNode{root}

	for len(cur) > 0 {
		result = append(result, cur[0].Val)

		newCur := make([]*TreeNode, 0)

		for i := range cur {
			if cur[i].Right != nil {
				newCur = append(newCur, cur[i].Right)
			}
			if cur[i].Left != nil {
				newCur = append(newCur, cur[i].Left)
			}

		}

		cur = newCur

	}

	return result
}
