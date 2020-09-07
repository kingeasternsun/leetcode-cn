package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {

	if head == nil {
		return nil
	}

	num := 0
	cur := head
	for {
		if cur == nil {
			break
		}

		num++
		cur = cur.Next
	}

	return helper(head, num)

}

func helper(head *ListNode, num int) *TreeNode {

	// fmt.Println("num", num)
	if head == nil {
		return nil
	}

	if num == 0 { //这里一定要判定 num为0的情况，因为中间的子链在操作时，head可能不为nil
		// fmt.Println("num is zero", num, head.Val)
		return nil
	}

	if num == 1 {
		// fmt.Println("leaf", head.Val)
		return &TreeNode{Val: head.Val}
	}

	mid := num / 2
	i := 0
	cur := head
	for i < mid {
		cur = cur.Next
		i++
	}

	// fmt.Println("mid value", cur.Val, num)
	curNode := &TreeNode{
		Val: cur.Val,
	}

	left := helper(head, mid)
	right := helper(cur.Next, num-1-mid)
	curNode.Left = left
	curNode.Right = right

	return curNode

}
