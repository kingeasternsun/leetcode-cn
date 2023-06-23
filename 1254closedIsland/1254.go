package main

func closedIsland(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	num := 0
	for i := range grid {
		for j := range grid[i] {
			// if is water or visited
			if grid[i][j] > 0 {
				continue
			}
			num += bfs(Point{i, j}, grid)
		}
	}
	return num
}

type Point struct {
	x, y int
}

var dirs = []int{0, 1, 0, -1, 0}

func bfs(p Point, grid [][]int) int {

	boundry := false
	points := []Point{p}
	for len(points) > 0 {

		newPoints := []Point{}
		for _, p := range points {
			// reach boundry
			if p.x == 0 || p.y == 0 || p.x == len(grid)-1 || p.y == len(grid[0])-1 {
				boundry = true
			}

			grid[p.x][p.y] = 2

			for i := 0; i < 4; i++ {
				newP := Point{x: p.x + dirs[i], y: p.y + dirs[i+1]}
				if newP.x < 0 || newP.y < 0 || newP.x > len(grid)-1 || newP.y > len(grid[0])-1 {
					continue
				}

				// water or visited
				if grid[newP.x][newP.y] > 0 {
					continue
				}

				newPoints = append(newPoints, newP)

			}

		}

		points = newPoints
	}

	if boundry {
		return 0
	}
	return 1
}
