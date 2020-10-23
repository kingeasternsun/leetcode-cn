package main

import "fmt"

func main() {
	c := &ListNode{Val: 1}
	b := &ListNode{Val: 2, Next: c}
	a := &ListNode{Val: 1, Next: b}
	fmt.Println(isPalindrome(a))
	fmt.Println(isPalindrome(b))
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1. 先找到中间节点
2. 然后利用递归,递归到中间节点，右侧节点向右移动后返回
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// var end int
	mid, cnt := findMiddle(head)
	if cnt&1 == 1 {
		mid = mid.Next
	}
	end := cnt/2 - 1 //如果偶数，例如2个，左边只需要下标0就可了，如果奇数，例如3，左边也只需要到下标0就可以
	_, result := help(head, mid, 0, end)

	return result
}

//找到中间节点,如果是偶数 返回右边那个
func findMiddle(head *ListNode) (*ListNode, int) {
	one := head
	two := head.Next
	cnt := 1

	for two != nil {
		one = one.Next
		two = two.Next
		cnt = cnt + 1
		if two != nil {
			two = two.Next
			cnt++
		}
	}

	return one, cnt
}

//比较左右是否相同，然后right向右边移动,pos 表示left是当前第几个点
func help(left, right *ListNode, pos, end int) (*ListNode, bool) {
	if pos == end {
		//到达中间位置，直接比较,然后右边节点向右移动
		return right.Next, left.Val == right.Val
	}

	right, result := help(left.Next, right, pos+1, end)
	if !result {
		return nil, result
	}

	return right.Next, left.Val == right.Val
}
