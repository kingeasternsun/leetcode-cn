package numenclaves

func numEnclaves(grid [][]int) int {
	cnt := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 0 || grid[row][col] == -1 {
				continue
			}

			tmpCnt, ok := bfs(grid, row, col)
			if !ok {
				cnt += tmpCnt
			}
		}
	}
	return cnt
}

func bfs(grid [][]int, x, y int) (int, bool) {
	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	ret := false
	cnt := 0
	queue := make([][2]int, 0, 1)

	queue = append(queue, [2]int{x, y})
	grid[x][y] = -1
	cnt++

	for len(queue) > 0 {
		newQueue := make([][2]int, 0)
		for _, q := range queue {
			for _, dir := range dirs {
				newX := q[0] + dir[0]
				newY := q[1] + dir[1]
				if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[0]) {
					ret = true
					continue
				}

				if grid[newX][newY] != 1 {
					continue
				}
				newQueue = append(newQueue, [2]int{newX, newY})
				grid[newX][newY] = -1
				cnt++
			}
		}

		queue = newQueue

	}
	return cnt, ret
}
