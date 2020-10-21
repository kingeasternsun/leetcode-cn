package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {

	var maxCount, recentCount, recent int
	result := make([]int, 0)
	find(root, &maxCount, &result, &recent, &recentCount)

	return result
}

func find(root *TreeNode, maxCount *int, result *[]int, recent *int, recentCount *int) {
	if root == nil {
		return
	}

	find(root.Left, maxCount, result, recent, recentCount)

	if len(*result) == 0 {
		*result = append(*result, root.Val)
		*maxCount = 1
		*recent = root.Val
		*recentCount = 1
	} else {

		if root.Val == *recent {
			*recentCount = *recentCount + 1
		} else {
			*recent = root.Val
			*recentCount = 1
		}

		//和当前的众数的个数一样
		if *recentCount == *maxCount {
			*result = append(*result, root.Val)
		} else if *recentCount > *maxCount { //比当前众数的个数还要多
			*result = make([]int, 0)
			*result = append(*result, root.Val)
			*maxCount = *recentCount
		}

	}

	find(root.Right, maxCount, result, recent, recentCount)

}
