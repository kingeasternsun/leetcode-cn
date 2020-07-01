package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*


1 2 3 4 5 6

2 3 1 4 5 6
3 2 1 4 5 6


1 2 3 4 5

2 1 3 4 5
2 1 4 3 5

*/

/*

[beg end nextbeg]
反转后变为
[end beg nextbeg]

*/
func reverseKGroup(head *ListNode, k int) *ListNode {

	if k == 0 || k == 1 {
		return head
	}

	var newHead *ListNode
	var preEnd *ListNode //用于当前段反转后，前面段连接到当前段

	beg := head
	for beg != nil {

		//得到end节点
		end := beg
		i := 1
		for i < k {
			end = end.Next
			if end == nil {

				if newHead == nil {
					newHead = beg
				}

				return newHead
			}
			i++
		}

		// 下一个开始节点
		newBeg := end.Next

		//反转
		tmp := beg
		for tmp != end {

			curNext := tmp.Next

			//插入到end的后面
			tmp.Next = end.Next
			end.Next = tmp

			tmp = curNext
		}

		// 反转之后，上面段链接到这里
		if preEnd != nil {
			preEnd.Next = end
		}

		if newHead == nil {
			newHead = end
		}

		//反转之后
		preEnd = beg
		beg = newBeg

	}

	return newHead
}
