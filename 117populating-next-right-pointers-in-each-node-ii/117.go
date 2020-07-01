package main

//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//bfs 4ms 5.9MB
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

//4ms 5.9MB
/*
 just record the preNode
*/
func connect2(root *Node) *Node {

	if root == nil {
		return root
	}

	leftMostNode := root
	for leftMostNode != nil {

		var tmpLeftNode, preNode *Node

		for leftMostNode != nil {

			if leftMostNode.Left != nil && leftMostNode.Right != nil {
				leftMostNode.Left.Next = leftMostNode.Right
				if preNode != nil {
					preNode.Next = leftMostNode.Left
				}
				preNode = leftMostNode.Right

				if tmpLeftNode == nil {
					tmpLeftNode = leftMostNode.Left
				}

			} else if leftMostNode.Left != nil {
				if preNode != nil {
					preNode.Next = leftMostNode.Left
				}
				preNode = leftMostNode.Left
				if tmpLeftNode == nil {
					tmpLeftNode = leftMostNode.Left
				}
			} else if leftMostNode.Right != nil {
				if preNode != nil {
					preNode.Next = leftMostNode.Right
				}
				preNode = leftMostNode.Right
				if tmpLeftNode == nil {
					tmpLeftNode = leftMostNode.Right
				}
			} else {
				//nothing
			}

			leftMostNode = leftMostNode.Next
		}

		leftMostNode = tmpLeftNode

	}

	return root
}

//4ms 5.9MB
func connect(root *Node) *Node {

	if root == nil {
		return root
	}

	leftMostNode := root
	for leftMostNode != nil {

		tmpNode := leftMostNode
		leftMostNode = nil
		var preNode *Node

		for tmpNode != nil {

			var leftNode, rightNode *Node
			if tmpNode.Left == nil && tmpNode.Right == nil {
				tmpNode = tmpNode.Next
				continue
			}

			if tmpNode.Left != nil && tmpNode.Right != nil {
				tmpNode.Left.Next = tmpNode.Right
				leftNode = tmpNode.Left
				rightNode = tmpNode.Right
			} else if tmpNode.Left != nil {
				leftNode = tmpNode.Left
				rightNode = tmpNode.Left
			} else {
				leftNode = tmpNode.Right
				rightNode = tmpNode.Right
			}

			if preNode != nil {
				preNode.Next = leftNode
			}
			preNode = rightNode

			if leftMostNode == nil {
				leftMostNode = leftNode
			}

			tmpNode = tmpNode.Next
		}

	}

	return root
}
