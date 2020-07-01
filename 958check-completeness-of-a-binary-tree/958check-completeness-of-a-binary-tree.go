package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
//958. 二叉树的完全性检验
https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/

- 执行深度优先遍历（DFS），将每次遍历的结果存入result
- 判断result
- 如果是完全性的检验，那么result的值从左到右肯定是保持一样或减少，而且最多比树高少一层
- 中间如果左子树为空，右边子树存在，肯定就是false
*/

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	result := make([]int, 0)
	// result = append(result,)
	if helper(root, 0, &result) == false {
		return false
	}

	if len(result) == 0 {
		return true
	}

	a := result[0]
	maxH := result[0]
	for i := range result {
		//保证左边比右边相等或更高
		if result[i] > a {
			return false
		}

		//当高度如果比整个树高少了两层
		if result[i] <= maxH-2 {
			return false
		}

		a = result[i]
	}
	return true
}

func helper(root *TreeNode, height int, result *[]int) (isComplete bool) {

	if root == nil {
		*result = append(*result, height)
		return true
	}

	if root.Left == nil && root.Right != nil {
		return false
	}

	leftC := helper(root.Left, height+1, result)
	if !leftC {
		return false
	}

	rightC := helper(root.Right, height+1, result)
	if !rightC {
		return false
	}

	return true
}
