package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal2(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	var pre, cur *TreeNode
	result := make([]int, 0)
	s := make([]*TreeNode, 0)
	s = append(s, root)
	sLen := 1
	for {

		if sLen == 0 {
			return result
		}

		cur = s[sLen-1]

		//往下走
		if pre == nil || cur == pre.Left || cur == pre.Right {
			if cur.Left != nil {
				s = append(s, cur.Left)
				sLen++
			} else if cur.Right != nil {
				s = append(s, cur.Right)
				sLen++
			}
		} else if cur.Left == pre { //往上走
			if cur.Right != nil {
				s = append(s, cur.Right)
				sLen++
			}
		} else { // pre = cur //走到最左边点
			result = append(result, cur.Val)
			s = s[:sLen-1]
			sLen--
		}

		pre = cur

	}

}

/*  参考
作者：hzhu212
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/yan-se-biao-ji-fa-yi-chong-tong-yong-qie-jian-ming/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
type Node struct {
	Node  *TreeNode
	color bool // true 表示访问 ，false表示没有访问
}

/*
使用颜色标记节点的状态，新节点为白色，已访问的节点为灰色。
如果遇到的节点为白色，则将其标记为灰色，然后将其右子节点、自身、左子节点依次入栈。
如果遇到的节点为灰色，则将节点的值输出。
*/
func postorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*Node, 0)
	stack = append(stack, &Node{root, false})
	skLen := 1

	for {

		if skLen == 0 {
			return result
		}

		cur := stack[skLen-1]
		// stack = stack[:skLen-1]
		// skLen--
		if !cur.color {
			cur.color = true
			if cur.Node.Right != nil {
				stack = append(stack, &Node{cur.Node.Right, false})
				skLen++
			}

			if cur.Node.Left != nil {
				stack = append(stack, &Node{cur.Node.Left, false})
				skLen++
			}
		} else {
			result = append(result, cur.Node.Val)
			stack = stack[:skLen-1]
			skLen--
		}

	}

}
