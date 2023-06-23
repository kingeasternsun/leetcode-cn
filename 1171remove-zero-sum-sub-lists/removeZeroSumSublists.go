package main

import (
	"fmt"
	"time"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	ok := false

	for {
		fmt.Println("first", head.Val)
		head, ok = removeZeroSumSublistsOnce(head)
		if !ok {
			return head
		}

		time.Sleep(time.Second)
	}

}

func removeZeroSumSublistsOnce(head *ListNode) (*ListNode, bool) {
	if head == nil {
		return nil, false
	}

	sumMap := make(map[int]*ListNode, 0)

	preHead := &ListNode{Val: 0, Next: head}
	sumMap[0] = preHead

	cur := head
	curSum := 0
	for cur != nil {
		curSum += cur.Val
		fmt.Println(curSum)

		if preNode, ok := sumMap[curSum]; ok {
			preNode.Next = cur.Next
			return preHead.Next, ok
		} else {
			sumMap[curSum] = cur
			cur = cur.Next
		}
	}

	return preHead.Next, false

}

func main() {

	// [1,2,3,-3,-2]
	// n5 := &ListNode{Val: -2}
	// n4 := &ListNode{Val: -3, Next: n5}
	// n3 := &ListNode{Val: 3, Next: n4}
	// n2 := &ListNode{Val: 2, Next: n3}
	// n1 := &ListNode{Val: 1, Next: n2}

	n5 := &ListNode{Val: 1}
	n4 := &ListNode{Val: 3, Next: n5}
	n3 := &ListNode{Val: -3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}

	removeZeroSumSublists(n1)

}
