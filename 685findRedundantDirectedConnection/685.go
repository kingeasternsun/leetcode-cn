package main

func main() {

}

// 参考了 https://leetcode-cn.com/problems/redundant-connection-ii/solution/685-rong-yu-lian-jie-iibing-cha-ji-de-ying-yong-xi/
//先查找入度为2，也就是有两个父节点的；如果找不到，就是有环，再从环里面找

//查找根节点
func findRoot(parent []int, child int) int {
	for parent[child] > 0 {
		child = parent[child]
	}
	return child
}

func join(parent []int, u, v int) {

	rootu := findRoot(parent, u)
	rootv := findRoot(parent, v)
	if rootu == rootv {
		return
	}

	parent[rootv] = rootu

}

func same(parent []int, u, v int) bool {

	rootu := findRoot(parent, u)
	rootv := findRoot(parent, v)
	return rootu == rootv

}

func isTreeDeleteEdge(edges [][]int, deleteEdge int) bool {
	parent := make([]int, len(edges)+1)
	// for _, edge := range edges {
	// 	parent[edge[1]] = edge[0]
	// }

	for i, edge := range edges {

		if i == deleteEdge {
			continue
		}

		if same(parent, edge[0], edge[1]) {
			return false
		}

		join(parent, edge[0], edge[1])
	}

	return true
}

func findCycleEdge(edges [][]int) []int {
	parent := make([]int, len(edges)+1)

	for _, edge := range edges {

		if same(parent, edge[0], edge[1]) {
			return edge
		}

		join(parent, edge[0], edge[1])
	}

	return nil
}

func findRedundantDirectedConnection(edges [][]int) []int {

	ingree := make([]int, len(edges)+1)
	var ingreddv int
	for _, edge := range edges {

		ingree[edge[1]] = ingree[edge[1]] + 1
		if ingree[edge[1]] == 2 {
			ingreddv = edge[1] //入度为2的
			break
		}
	}

	if ingreddv > 0 {

		//记录入度为2的边
		dupv := make([]int, 0)
		for i, edge := range edges {

			if edge[1] == ingreddv {
				dupv = append(dupv, i)
			}
		}

		for i := len(dupv) - 1; i >= 0; i-- {

			if isTreeDeleteEdge(edges, dupv[i]) {
				return edges[dupv[i]]
			}

		}

		return nil

	}

	//查找构成环的
	return findCycleEdge(edges)

}
