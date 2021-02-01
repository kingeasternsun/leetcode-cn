package leetcode

/*
判断每个节点的度，度为1的就是起点或终点
*/
func restoreArray(adjacentPairs [][]int) []int {
	g := make(map[int][]int, 0)
	for _, ad := range adjacentPairs {
		g[ad[0]] = append(g[ad[0]], ad[1])
		g[ad[1]] = append(g[ad[1]], ad[0])
	}

	var root int
	for k, v := range g {
		if len(v) == 1 {
			root = k
			break
		}
	}

	//记录哪个已经访问过了
	visited := make(map[int]struct{}, 0)
	visited[root] = struct{}{}

	res := []int{root}
	for {
		end := res[len(res)-1]
		newAddNum := 0
		for _, ad := range g[end] {
			if _, ok := visited[ad]; !ok { //相邻的没有访问过
				visited[ad] = struct{}{}
				newAddNum++
				res = append(res, ad)
			}
		}
		if newAddNum == 0 {
			break
		}
	}

	return res

}
