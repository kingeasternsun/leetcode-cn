package levelOrder

/**
 * Definition for a Node.
 **/
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	//BFS

	if root == nil {
		return nil
	}
	res := make([][]int, 0)

	cur := []*Node{root}
	res = append(res, []int{root.Val})

	for len(cur) > 0 {

		tmp := make([]*Node, 0)
		tmpV := make([]int, 0)
		for _, v := range cur {

			for _, vc := range v.Children {

				tmp = append(tmp, vc)
				tmpV = append(tmpV, vc.Val)

			}
		}

		//当前层遍历结束后
		if len(tmpV) > 0 {
			res = append(res, tmpV)
		}

		cur = tmp

	}

	return res

}
