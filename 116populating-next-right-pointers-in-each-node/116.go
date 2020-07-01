package main

//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//bfs 8ms 6.1MB
func connect1(root *Node) *Node {

	if root == nil {
		return nil
	}

	cur := []*Node{root}

	for len(cur) > 0 {

		newCur := make([]*Node, 0)
		if len(cur) > 1 {

			i := 0
			for i < len(cur)-1 {
				cur[i].Next = cur[i+1]
				i++
			}

		}

		for i := range cur {
			if cur[i].Left != nil {
				newCur = append(newCur, cur[i].Left)
			}
			if cur[i].Right != nil {
				newCur = append(newCur, cur[i].Right)
			}
		}
		cur = newCur
	}

	return root
}

//8ms 6.1MB
func connect(root *Node) *Node {

	if root == nil {
		return root
	}

	leftMostNode := root
	for leftMostNode.Left != nil {

		tmpLeftNode := leftMostNode.Left

		for leftMostNode != nil {

			leftMostNode.Left.Next = leftMostNode.Right

			if leftMostNode.Next != nil {
				leftMostNode.Right.Next = leftMostNode.Next.Left
			}

			leftMostNode = leftMostNode.Next
		}

		leftMostNode = tmpLeftNode

	}

	return root
}
