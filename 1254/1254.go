package main

func main() {

}
func closedIsland(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}

	col := len(grid[0])
	if col == 0 {
		return 0
	}

	visted := make([]bool, row*col)
	num := 0
	for i := range grid {
		for j := range grid[i] {

			if grid[i][j] == 1 { // water
				continue
			}

			num += bfs(i, j, grid, row, col, visted)
		}
	}

	return num
}

type Point struct {
	x, y int
}

var dir = []int{0, 1, 0, -1, 0}

//如果到达边界 就返回0,
func bfs(x, y int, grid [][]int, row, col int, visted []bool) int {

	ret := 1

	if visted[x*col+y] {
		return 0
	}

	q := []Point{Point{x, y}}
	for len(q) > 0 {

		newQ := make([]Point, 0)

		for _, cur := range q {

			// in bondary
			if cur.x == 0 || cur.y == 0 || cur.x == row-1 || cur.y == col-1 {
				ret = 0
			}

			for d := 0; d < 4; d++ {
				newPos := Point{
					x: cur.x + dir[d],
					y: cur.y + dir[d+1],
				}

				if 0 <= newPos.x && newPos.x < row && 0 <= newPos.y && newPos.y < col {
					if grid[newPos.x][newPos.y] == 1 { //water
						continue
					}

					if visted[newPos.x*col+newPos.y] {
						continue
					}

					newQ = append(newQ, newPos)
					visted[newPos.x*col+newPos.y] = true

				}
			}
		}

		q = newQ
	}

	return ret

}
