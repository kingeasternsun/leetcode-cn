package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 二叉树的前序遍历
- 一路向左，将左节点放入结果，将遇到的右边子节点入栈
- 到达最最边到叶子节点后，出栈，在继续一路向东
*/
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	sLen := 0

	cur := root
	for {
		for cur != nil {
			result = append(result, cur.Val)
			if cur.Right != nil {
				stack = append(stack, cur.Right)
				sLen++
			}
			cur = cur.Left
		}

		if sLen == 0 {
			return result
		}

		// 出栈
		cur = stack[sLen-1]
		stack = stack[:sLen-1]
		sLen--

	}
}
