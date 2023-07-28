package hascycle

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	a := head
	b := head

	for b != nil && b.Next != nil {
		// 不能放在这里判断，不然一开始a和b就是相同的
		// if a == b {
		// 	return true
		// }
		a = a.Next
		b = b.Next
		b = b.Next
		if a == b {
			return true
		}
	}

	return false

}

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	a := head
	b := head.Next

	for b != nil && b.Next != nil {
		// 因为一开始a和b不一样，所以这里可以放
		if a == b {
			return true
		}
		a = a.Next
		b = b.Next
		b = b.Next
	}

	return false

}
