package detectcycle

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	a := head
	b := head

	for b != nil && b.Next != nil {
		a = a.Next
		b = b.Next.Next
		if a == b {
			h := head
			for a != h {
				a = a.Next
				h = h.Next
			}
			return h
		}

	}
	return nil

}
