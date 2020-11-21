package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	sortNode := head //sortNode

	for sortNode.Next != nil {

		// 特殊情况1 ，已经比前面的要大，直接往后移动
		if sortNode.Val <= sortNode.Next.Val {
			sortNode = sortNode.Next
			continue
		}

		cur := sortNode.Next //要插入的节点
		sortNode.Next = sortNode.Next.Next

		// 特殊情况2 比head还要小，直接放到最前面
		if cur.Val <= head.Val {
			cur.Next = head
			head = cur
			continue
		}

		//因为前面两个特殊情况已经提前处理过了，所以 这里插入的肯定在head后，sortNode前面
		for tmp := head; tmp != sortNode; tmp = tmp.Next {
			//找到要插入的位置
			if cur.Val <= tmp.Next.Val {

				cur.Next = tmp.Next
				tmp.Next = cur

				break
			}
		}

	}

	return head
}
