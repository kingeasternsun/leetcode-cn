package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
142. 环形链表 II https://leetcode-cn.com/problems/linked-list-cycle-ii/
1. 假设链表中，环形链表交叉点之前有m个节点，环中有n个节点，假设n大于m
2. 快慢指针开始移动，当慢指针移动到环形链表交叉点的时候，快指针离这个交叉点之后m个节点上。
3. 那么快指针离慢指针还差n-m节点，
4. 慢指针再移动n-m，快指针移动2*（n-m）个，两个就相遇了。
5. 此时慢指针离这个交叉点之后n-m个节点上
6. 因为环中有n个节点， 此时慢指针离这个交叉点之后n-m个节点上，那么慢指针再走m个节点就可以到达交叉点了
7. 由于环形交叉点之前有m个节点，所以另外一个节点再从头开始和慢指针一起移动，就会再次相遇到交叉点

https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/ 这位也讲的很清楚
*/
//8ms 3.5MB
func detectCycle(head *ListNode) *ListNode {

	one := head
	two := head
	var mid *ListNode

	for two != nil && two.Next != nil {

		one = one.Next

		two = two.Next
		two = two.Next

		if one == two {
			mid = one //此时慢指针离这个交叉点之后n-m个节点上，那么慢指针再走m个节点就可以到达交叉点了
			break
		}

	}

	if mid == nil {
		return mid
	}

	one = head

	for one != nil {
		if one == mid {
			return one
		}

		one = one.Next
		mid = mid.Next
	}

	return nil

}
