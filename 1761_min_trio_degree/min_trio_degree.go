package mintriodegree

// 1284ms 12.15mb
func minTrioDegree(n int, edges [][]int) int {
	newEdges := make([]map[int]struct{}, n+1)
	for i := 0; i < n+1; i++ {
		newEdges[i] = make(map[int]struct{})
	}

	// 记录每个节点相邻的节点
	for _, edge := range edges {
		// if edge[0] > edge[1] {
		// 	edge[0], edge[1] = edge[1], edge[0]
		// }
		newEdges[edge[0]][edge[1]] = struct{}{}
		newEdges[edge[1]][edge[0]] = struct{}{}
	}

	ret := -1
	for node, neighbour := range newEdges {
		for u := range neighbour {
			for v := range neighbour {
				//  两个去重：1. 只考虑比node大的相邻节点; 2. 两个相邻节点只考虑 u < v
				if u == v || u < node || v < node || u > v {
					continue
				}

				// 如果u，v相连，说明构成三元组
				if _, ok := newEdges[u][v]; ok {
					tmp := len(neighbour) - 2 + len(newEdges[u]) - 2 + len(newEdges[v]) - 2
					if ret == -1 || tmp < ret {
						ret = tmp
					}

				}

			}

		}
	}

	return ret
}
