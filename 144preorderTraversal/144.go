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

func preorderTraversal1(root *TreeNode) []int {

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

func preorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	res := make([]int, 0)

	queue := make([]*TreeNode, 0)
	// queue = append(queue, root)

	cur := root

	for {

		//一路向左边遍历，返回当前的val，加入节点的右节点
		for cur != nil {

			res = append(res, cur.Val)

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

			cur = cur.Left

		}

		if len(queue) == 0 {
			return res
		}

		//出栈
		cur = queue[len(queue)-1]
		queue = queue[:len(queue)-1]

	}

	return res
}
