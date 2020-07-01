package main

// https://leetcode-cn.com/problems/rotate-list/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	nLen := 0
	cur := head
	for cur != nil {
		nLen++
		if nLen == k+1 {
			break
		}
		cur = cur.Next
	}

	// cur == nil mearns  k is larger than nLen
	if cur == nil { // there nLen is the list length
		k = k % nLen
		if k == 0 {
			return head
		}
		cur = head
		for k > 0 {
			cur = cur.Next
			k--
		}
	}

	//finally ,get the kth node from begin, use two pointer to move to end
	pre := head
	for cur.Next != nil {
		cur = cur.Next
		pre = pre.Next
	}

	newHead := pre.Next
	pre.Next = nil
	cur.Next = head
	return newHead

}
