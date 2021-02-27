/*
 * @Description: 21
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-25 21:41:16
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-25 21:48:42
 * @FilePath: /21mergeTwoLists/21.go
 */
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return head.Next

}
