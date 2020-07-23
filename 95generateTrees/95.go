package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return help(1, n)
}

func help(start, end int) []*TreeNode {

	if start > end {
		return []*TreeNode{nil}
	}

	if start == end {
		return []*TreeNode{&TreeNode{Val: start}}
	}

	result := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {

		//以 i 作为当前根节点
		left := help(start, i-1)
		right := help(i+1, end)

		for j := range left {

			for k := range right {

				root := &TreeNode{
					Val:   i,
					Left:  left[j],
					Right: right[k],
				}

				result = append(result, root)

			}
		}

	}

	return result

}
