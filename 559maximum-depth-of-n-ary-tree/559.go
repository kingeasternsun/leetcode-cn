package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {

	if root == nil {
		return 0
	}
	cur := []*Node{root}

	depth := 0
	for len(cur) > 0 {
		depth++
		newLevel := make([]*Node, 0)
		for _, tmpNode := range cur {
			// for j := range tmpNode.Children {
			newLevel = append(newLevel, tmpNode.Children...)
			// }
		}
		cur = newLevel

	}

	return depth

}
