package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	preAHead := &ListNode{} //用于记录大于等于x的节点
	preA := preAHead
	preBHead := &ListNode{} //用于记录小于x的节点
	preB := preBHead

	for ; head != nil; head = head.Next {
		if head.Val >= x {
			preB.Next = head
			preB = preB.Next
		} else {
			preA.Next = head
			preA = preA.Next
		}
	}

	//拼接起来
	preA.Next = preBHead.Next
	preB.Next = nil //重要 记得要置为nil

	return preAHead.Next
}
