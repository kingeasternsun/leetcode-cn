package reversebetween

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if left == right {
		return head
	}

	pre := &ListNode{Next: head}
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	left_node := pre.Next
	pre.Next = nil

	right_pre := left_node
	for i := 0; i < right-left-1; i++ {
		right_pre = right_pre.Next
	}

	right_node := right_pre.Next
	right_pre.Next = nil

	tmpHead, tmpTail := reverse(left_node)
	pre.Next = tmpHead
	tmpTail.Next = right_node

	return pre.Next

}

func reverse(head *ListNode) (*ListNode, *ListNode) {

	if head.Next == nil {
		return head, head
	}

	tmpHead, tmpTail := reverse(head.Next)
	head.Next = nil
	tmpTail.Next = head

	return tmpHead, head
}
