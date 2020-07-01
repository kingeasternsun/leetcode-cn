package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
234. 回文链表
https://leetcode-cn.com/problems/palindrome-linked-list/
1. 先通过快慢指针找到中间节点
2. 从中间节点将链表反向连接起来
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	//获取中间节点
	mid := middleNode(head)
	cur := mid
	next := cur.Next
	for next != nil {
		tmp := next.Next
		next.Next = cur
		cur = next
		next = tmp
	}

	for head != mid {
		if head.Val != cur.Val {
			return false
		}

		head = head.Next
		cur = cur.Next
	}

	return true

}

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
