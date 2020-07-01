package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	head, _ = help(head)

	return head
}

func help(head *ListNode) (newHead, newEnd *ListNode) {

	if head == nil {
		return nil, nil
	}

	if head.Next == nil {
		return head, head
	}

	newHead, newEnd = help(head.Next)
	newEnd.Next = head
	head.Next = nil

	return newHead, head
}
