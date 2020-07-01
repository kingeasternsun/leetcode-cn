package main

/*
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)

	s := make([]*TreeNode, 0) //队列
	s = append(s, root)
	sLen := 1 //当前层的长度
	for {
		if sLen == 0 {
			return result
		}

		tmpResult := make([]int, sLen) //当前层遍历结果
		tmpLen := 0                    //下一层的长度
		for i := 0; i < sLen; i++ {    //遍历当前层
			tmpResult[i] = s[i].Val
			if s[i].Left != nil {
				s = append(s, s[i].Left)
				tmpLen++
			}
			if s[i].Right != nil {
				s = append(s, s[i].Right)
				tmpLen++

			}
		}

		result = append(result, tmpResult)
		s = s[sLen:]
		sLen = tmpLen

	}

}
