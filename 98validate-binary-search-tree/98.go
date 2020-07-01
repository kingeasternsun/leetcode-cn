package main



//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	v, _,_ := help(root)

	return v
}

func help(root *TreeNode) (bool,int, int) {

	if root.Left == nil && root.Right == nil {
		return true, root.Val, root.Val
	}

	minV:=root.Val
	maxV:=root.Val
	if root.Left != nil {
		valid,minValue, maxValue := help(root.Left)
		if !valid {
			return false, minV,maxV
		}

		if maxValue >= root.Val {
			return false,  minV,maxV
		}

		minV = minValue
		maxV = root.Val
	}

	if root.Right != nil {
		valid, minValue,maxValue:= help(root.Right)
		if !valid {
			return false, minV,maxV
		}

		if minValue <= root.Val {
			return false,  minV,maxV
		}

		maxV = maxValue
	}

	return true, minV,maxV
}
