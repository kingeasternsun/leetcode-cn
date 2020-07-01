package main

type Neibor struct {
	N int
	W int
}

func minReorder(n int, connections [][]int) int {

	g := make([][]Neibor, n)

	for i := range connections {
		//因为道路原先就是从 connections[i][0] 到 connections[i][1]，如果刚好和bfs的访问路径一致，也就是方向从城市0出来的方向，就需要反转
		g[connections[i][0]] = append(g[connections[i][0]], Neibor{N: connections[i][1], W: 1})

		g[connections[i][1]] = append(g[connections[i][1]], Neibor{N: connections[i][0], W: 0})
	}

	visited := make([]bool, n)
	// bfs
	wight := 0
	cities := []int{0}
	visited[0] = true

	for len(cities) > 0 {

		newCities := make([]int, 0)

		for i := range cities {

			for _, neibor := range g[cities[i]] {
				if !visited[neibor.N] {
					newCities = append(newCities, neibor.N)
					visited[neibor.N] = true
					wight += neibor.W
				}
			}

		}

		cities = newCities

	}

	return wight

}
