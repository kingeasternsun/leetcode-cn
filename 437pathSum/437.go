package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {

	if root == nil {
		return 0
	}

	return rootSum(root, sum) +
		pathSum(root.Left, sum) + pathSum(root.Right, sum)

}

//rootSum 表示以当前root为起点，路径和为sum的个数
func rootSum(root *TreeNode, sum int) int {

	if root == nil {
		return 0
	}

	n := 0
	if sum == root.Val { //注意这里不是终止，因为后面路径和可以为0
		n = 1
	}

	return n + rootSum(root.Left, sum-root.Val) + rootSum(root.Right, sum-root.Val)

}
