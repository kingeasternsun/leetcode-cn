package leetcode

//二分法 加DFS  76 ms	7.8 MB
func minimumEffortPathDFS(heights [][]int) int {
	if len(heights) == 0 && len(heights[0]) == 0 {
		return 0
	}

	beg := 0
	end := 1000000
	for beg < end {
		mid := (end-beg)/2 + beg
		g := NewGraph(heights)
		if g.DFS(0, 0, mid) { //如果可以，就缩小高度再尝试
			end = mid
		} else {
			beg = mid + 1
		}

	}

	return beg

}

//184 ms	7.3 MB
func minimumEffortPathBFS(heights [][]int) int {
	if len(heights) == 0 && len(heights[0]) == 0 {
		return 0
	}

	beg := 0
	end := 1000000
	for beg < end {
		mid := (end-beg)/2 + beg
		g := NewGraph(heights)
		if g.BFS(mid) { //如果可以，就缩小高度再尝试
			end = mid
		} else {
			beg = mid + 1
		}

	}

	return beg

}

var dirs = []int{0, 1, 0, -1, 0}

func abs(x int) int {
	if x > 0 {
		return x
	}

	return -x
}

type Graph struct {
	heights [][]int
	visted  []bool
	rows    int
	cols    int
}

func NewGraph(heights [][]int) *Graph {

	return &Graph{
		heights: heights,
		rows:    len(heights),
		cols:    len(heights[0]),
		visted:  make([]bool, len(heights)*len(heights[0])),
	}
}

//判断h是否可以
func (g *Graph) DFS(x, y int, h int) bool {

	g.visted[x*g.cols+y] = true
	//可以到达最下角
	if x == g.rows-1 && y == g.cols-1 {
		return true
	}

	//对于4个方向
	for i := 0; i < 4; i++ {
		nx := x + dirs[i]
		ny := y + dirs[i+1]
		if 0 <= nx && nx < g.rows && 0 <= ny && ny < g.cols {
			if g.visted[nx*g.cols+ny] {
				continue
			}

			//体力消耗不大于h，就可以跳入nx,ny
			if abs(g.heights[x][y]-g.heights[nx][ny]) <= h {
				if g.DFS(nx, ny, h) {
					return true
				}
			}
		}

	}

	return false
}

func (g *Graph) BFS(h int) bool {

	cur := []int{0}
	g.visted[0] = true
	for len(cur) > 0 {

		newPoints := make([]int, 0)
		for _, v := range cur {
			x := v / g.cols
			y := v % g.cols

			if x == g.rows-1 && y == g.cols-1 {
				return true
			}

			//对于4个方向
			for i := 0; i < 4; i++ {
				nx := x + dirs[i]
				ny := y + dirs[i+1]
				if 0 <= nx && nx < g.rows && 0 <= ny && ny < g.cols {
					if g.visted[nx*g.cols+ny] {
						continue
					}

					//体力消耗不大于h，就可以跳入nx,ny
					if abs(g.heights[x][y]-g.heights[nx][ny]) <= h {
						g.visted[nx*g.cols+ny] = true
						newPoints = append(newPoints, nx*g.cols+ny)
					}
				}

			}

		}

		cur = newPoints

	}
	return false
}
