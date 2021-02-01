package leetcode

//https://github.com/kingeasternsun/leetcode-cn
func findRoot(i int, parent []int) int {

	for parent[i] != i {
		i = parent[i]
	}
	return i
}

func union(i int, j int, parent []int) {

	i = findRoot(i, parent)
	j = findRoot(j, parent)
	parent[j] = i

}

//讲每个方块分为4部分 进行合并
func regionsBySlashes(grid []string) int {

	rows := len(grid)
	if rows == 0 {
		return 0
	}

	cols := len(grid[0])
	if cols == 0 {
		return 0
	}

	parent := make([]int, 4*rows*cols)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	//向左，向上合并
	for rowID := 0; rowID < rows; rowID++ {

		for colID := 0; colID < cols; colID++ {
			squreID := rowID*cols + colID //方块的id

			//分为4块，顺时针编号，0号肯定和上面的2号能够合并， 3号肯定和左边的1号能够合并
			// 0 号和上面的2合并
			if rowID > 0 {
				union(squreID*4, (squreID-cols)*4+2, parent)
			}

			//3 号和左边的1合并
			if colID > 0 {
				union(squreID*4+3, (squreID-1)*4+1, parent)
			}

			// 当前方块中的 1和2可以合并，3和0可以合并
			if grid[rowID][colID] == '/' {
				union(squreID*4+1, squreID*4+2, parent)
				union(squreID*4, squreID*4+3, parent)

			} else if grid[rowID][colID] == '\\' { // 当前方块中的 0和1可以合并，2和3可以合并
				union(squreID*4, squreID*4+1, parent)
				union(squreID*4+2, squreID*4+3, parent)
			} else {
				union(squreID*4+1, squreID*4+2, parent)
				union(squreID*4+2, squreID*4+3, parent)
				union(squreID*4, squreID*4+3, parent)
			}

		}

	}

	groupSet := make(map[int]struct{}, 0)
	for i := 0; i < 4*rows*cols; i++ {
		groupSet[findRoot(i, parent)] = struct{}{}
	}

	return len(groupSet)
}
