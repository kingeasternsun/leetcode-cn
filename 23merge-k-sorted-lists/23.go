package main

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
type ListNode struct {
	Val  int
	Next *ListNode
}

//12ms 5.3MB
func mergeKLists(lists []*ListNode) *ListNode {

	count := len(lists)
	if count == 0 {
		return nil
	}
	if count == 1 {
		return lists[0]
	}

	left := mergeKLists(lists[:count/2])
	right := mergeKLists(lists[count/2:])
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	var head, pre, tmp *ListNode

	for left != nil && right != nil {

		if left.Val < right.Val {
			tmp = left
			left = left.Next
		} else {
			tmp = right
			right = right.Next
		}

		if head == nil {
			head = tmp
		} else {
			pre.Next = tmp
		}

		pre = tmp
	}
	if left != nil {
		pre.Next = left
	} else if right != nil {
		pre.Next = right
	}

	return head

}
