package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || q == nil || p == nil {
		return nil
	}

	if p.Val > q.Val {
		p, q = q, p
	}

	cur := root
	for cur != nil {
		if p.Val <= cur.Val && cur.Val <= q.Val {
			return cur
		}

		if cur.Val < p.Val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	return cur
}
