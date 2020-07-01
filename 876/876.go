package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 876. 链表的中间结点 https://leetcode-cn.com/problems/middle-of-the-linked-list/
func middleNode(head *ListNode) *ListNode {

	if head == nil {
		return nil

	}

	one := head
	two := head

	for two != nil && two.Next != nil { //这个判断是关键
		one = one.Next
		two = two.Next
		two = two.Next

	}

	return one

}
