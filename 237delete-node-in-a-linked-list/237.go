package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {

	if node.Next == nil {
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
	return

}
