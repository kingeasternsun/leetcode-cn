package main

type ListNode struct {
	    Val int
	    Next *ListNode
}
	
func removeElements(head *ListNode, val int) *ListNode {

	cur:=head
	var pre *ListNode

	for cur!=nil{

		if cur.Val!=val{
			pre = cur
			cur = cur.Next

			continue
		}

		head,pre = removeNode(head,pre,cur,val)
		if pre ==nil{
			return head
		}

		cur = pre.Next

	}

	return head
   
}


func removeNode(head,pre,cur *ListNode,val int)(newHead, newPre*ListNode){

	end:=cur
	for end!=nil && end.Val==val{
		end = end.Next
	}

	
	newPre = end

	//链表最开始的是val数值
	if pre == nil{
		newHead = end
	}else{
		newHead = head
		pre.Next = end

	}

	return 

}