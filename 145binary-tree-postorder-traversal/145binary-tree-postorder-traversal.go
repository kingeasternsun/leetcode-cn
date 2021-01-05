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

func postorderTraversal3(root *TreeNode) []int {

	cur := root
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)

	// 用于记录上次的访问（加入到res）的节点，因为是后续遍历，所以如果当前节点cur有Right节点，并且Right节点刚访问过（加入到了res中），那么
	// Right节点所对应的子树肯定都已经遍历过了，而且这个Right节点是最后一个访问的。 既然当前节点cur的Right已经访问过了，那么cur就可以访问了
	var pre *TreeNode
	for {
		//一路向左入栈
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		if len(stack) == 0 {
			return res
		}

		//获取栈顶的点，注意这里先不要出栈
		top := stack[len(stack)-1]
		if top.Right != nil && top.Right != pre { //如果top节点的Right节点还没有访问，那么就要先访问top的右边节点，top还保留栈中
			cur = top.Right

		} else { //top节点的Right节点不存在 或已经访问过了，那么就可以访问当前节点了

			res = append(res, top.Val)
			stack = stack[:len(stack)-1] //这里就可以出栈了

			pre = top //记录pre
		}

	}

	return
}
