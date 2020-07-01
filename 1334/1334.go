package main

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {

	if len(edges) == 0 {
		return 0
	}

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = math.MaxInt32
		}
	}

	for i := range edges {
		graph[edges[i][0]][edges[i][1]] = edges[i][2]
		graph[edges[i][1]][edges[i][0]] = edges[i][2]
	}

	for k := 0; k < n; k++ {
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if graph[u][k]+graph[v][k] < graph[u][v] {
					graph[u][v] = graph[u][k] + graph[v][k]
				}
			}
		}
	}

	var minNear = math.MaxInt32
	var result = -1
	for u := 0; u < n; u++ {
		near := 0
		for v := 0; v < n; v++ {

			//注意 u!=v
			if u != v && graph[u][v] <= distanceThreshold {
				near++
			}
		}

		//注意 near <= minNear ,= 是关键，因为 题目中如果相同 就要 取最大的
		if near <= minNear {
			minNear = near
			result = u
		}
	}

	return result

}
