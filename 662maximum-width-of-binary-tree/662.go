package main

import "fmt"

func main() {

	root := &TreeNode{
		Val: 1,
	}

	left1 := &TreeNode{
		Val: 2,
	}

	right1 := &TreeNode{
		Val: 3,
	}
	left2 := &TreeNode{
		Val: 4,
	}
	right2 := &TreeNode{
		Val: 5,
	}

	root.Left = left1
	root.Right = right1
	left1.Left = left2
	right1.Right = right2
	fmt.Println(widthOfBinaryTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	ids := make([]uint64, 1) //记录每个节点的标记
	ids = append(ids, 0)

	var maxWidth uint64 = 1

	for len(stack) > 0 {

		if ids[len(ids)-1]-ids[0]+1 > maxWidth {
			maxWidth = ids[len(ids)-1] - ids[0] + 1
		}

		newStack := make([]*TreeNode, 0)
		newids := make([]uint64, 0)
		for i := range stack {

			if stack[i].Left != nil {
				newStack = append(newStack, stack[i].Left)
				newids = append(newids, ids[i]<<1)
			}

			if stack[i].Right != nil {
				newStack = append(newStack, stack[i].Right)
				newids = append(newids, (ids[i]<<1)+1)
			}

		}

		stack = newStack
		ids = newids

	}

	return int(maxWidth)
}
