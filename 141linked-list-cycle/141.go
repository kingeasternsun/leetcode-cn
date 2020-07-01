package main

func main() {

}

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
141. 环形链表 https://leetcode-cn.com/problems/linked-list-cycle/
快慢指针
*/
func hasCycle(head *ListNode) bool {

	one := head
	two := head

	for two != nil && two.Next != nil {

		one = one.Next

		two = two.Next
		two = two.Next

		if one == two {
			return true
		}

	}
	return false

}
