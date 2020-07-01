package main
//https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
func main() {
	
}


//  * Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}
 
 func levelOrderBottom(root *TreeNode) [][]int {

	if root ==nil{
		return nil
	}

	result:=make([][]int,0)

	q:=make([]*TreeNode,0)
	qLen:=0
	q = append(q,root)
	qLen = 1

	for qLen>0 {

		tmp:=make([]int,qLen)

		newQLen:=0
		for i:=0;i<qLen;i++{
			tmp[i] = q[i].Val

			if q[i].Left!=nil{
				q = append(q,q[i].Left)
				newQLen++
			}

			if q[i].Right!=nil{
				q = append(q,q[i].Right)
				newQLen++
			}
			
		}

		q = q[qLen:]
		qLen = newQLen
		tmpResult:=append([][]int{},tmp)
		result =append(tmpResult,result...)

	}


	return result
}