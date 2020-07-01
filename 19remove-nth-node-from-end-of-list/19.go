package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1. 双指针移动
2. 优先判断k的长度刚好是链表长度的情况
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	if n == 0 {
		return head
	}

	// var pre *ListNode
	beg := head
	end := head

	//end先向右移动k
	k := 0
	for end != nil && k < n {
		end = end.Next
		k++
	}

	//说明链接长度刚好是 k , 删除第一个就可以了
	if end == nil {
		return head.Next
	}

	//同时往后移动 注意这里的终止条件是 end.Next
	for end.Next != nil {

		beg = beg.Next
		end = end.Next
	}

	beg.Next = beg.Next.Next

	return head
}
