/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-25 21:53:41
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-25 21:59:53
 * @FilePath: /83deleteDuplicates/83.go
 */
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	i := head
	j := head.Next

	for j != nil {
		if i.Val == j.Val {
			j = j.Next
			continue
		}

		i.Next = j
		i, j = j, j.Next
	}
	i.Next = nil

	return head

}
