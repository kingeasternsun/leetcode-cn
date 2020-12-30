package main

func main() {

}

//TreeNode  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*

- 出栈一个节点，加入输出
- 一路向左，输出当前节点，并把当前节点右节点入栈
*/

func preorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	res := make([]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		head := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		//一路向左边遍历
		for head != nil {

			res = append(res, head.Val)

			if head.Right != nil {
				queue = append(queue, head.Right)
			}

			head = head.Left

		}

	}

	return res
}
