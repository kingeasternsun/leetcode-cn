package leetcode

//DFS 递归做法
func makeConnectedDFS(n int, connections [][]int) int {

	/*

		1. 对计算机进行dfs或bfs遍历，看用到多少边 used
		2. 计算m个孤立网络
		3. 看剩余的边数是否 大于m-1
	*/

	g := make(map[int][]int, 0)
	for _, c := range connections {

		g[c[0]] = append(g[c[0]], c[1])
		g[c[1]] = append(g[c[1]], c[0])
	}

	usedEdge := 0
	visited := make([]bool, n)

	groupNum := 0

	var dfs func(i int)
	dfs = func(i int) {
		if visited[i] {
			return
		}

		visited[i] = true

		//访问i的相邻主机
		for _, j := range g[i] {
			if visited[j] {
				continue
			}

			usedEdge++
			// visited[j] = true
			dfs(j)

		}

	}

	for i := 0; i < n; i++ {

		if visited[i] {
			continue
		}

		groupNum++
		dfs(i)

	}

	if len(connections)-usedEdge < (groupNum - 1) {
		return -1
	}

	return groupNum - 1
}

//用bfs ，迭代做法
func makeConnected(n int, connections [][]int) int {

	/*

		1. 对计算机进行dfs或bfs遍历，看用到多少边 used
		2. 计算m个孤立网络
		3. 看剩余的边数是否 大于m-1
	*/

	g := make(map[int][]int, 0)
	for _, c := range connections {

		g[c[0]] = append(g[c[0]], c[1])
		g[c[1]] = append(g[c[1]], c[0])
	}

	//bfs遍历
	usedEdge := 0
	visited := make([]bool, n)

	groupNum := 0

	for i := 0; i < n; i++ {

		if visited[i] {
			continue
		}

		groupNum++
		usedEdge += bfs(g, i, visited)

	}

	if len(connections)-usedEdge < (groupNum - 1) {
		return -1
	}

	return groupNum - 1
}

func bfs(g map[int][]int, root int, visited []bool) (usedEdge int) {

	cur := []int{root}
	visited[root] = true

	for len(cur) > 0 { //cur中都是已经被标注过访问的

		next := make([]int, 0)

		for _, i := range cur {

			for _, j := range g[i] {

				if visited[j] {
					continue
				}

				visited[j] = true
				usedEdge++
				next = append(next, j)
			}

		}

		cur = next
	}

	return
}
