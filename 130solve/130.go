package main

/*
	- 将O的链接起来, 链接的时候，如果一个组在边界，两外一个组就链接到这个组上
	- 如果一个组的根是在边界，也就是dumpPos，那么就不能填充
*/

func findFather(father []int, i int) int {
	for father[i] != i {
		i = father[i]
	}
	return i

}

// i对应的组链接到j对应的组
// 如果提前知道两个组的大小，就将小的组链接到大的组上
// 如果i的根是 dumpPos ,j就链接到i上
func uniZip(father []int, i, j, dumpPos int) int {
	i2 := findFather(father, i)
	j2 := findFather(father, j)
	if i2 == j2 { //本来就在一个组里面了
		return 0
	}

	if i2 == dumpPos {
		i2, j2 = j2, i2
	}

	father[i2] = j2

	//路径压缩，把最底层的节点链接到根节点
	father[i] = j2
	father[j] = j2
	return 1
}

func solve1(board [][]byte) {

	m := len(board)
	if m == 0 {
		return
	}

	n := len(board[0])
	if n == 0 {
		return
	}

	dumpPos := m * n
	father := make([]int, dumpPos+1)
	for i := 0; i < dumpPos+1; i++ {
		father[i] = i
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}

			ID := i*n + j
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				father[ID] = dumpPos
			}
			//向上
			if i > 0 && board[i-1][j] == 'O' {
				// ID,ID-n 两个参数的先后顺序也有讲究
				// ID位置的深度大部分情况下都是 <= ID-n 位置的深度
				// 下面同理
				uniZip(father, ID, ID-n, dumpPos)
			}
			// 向左
			if j > 0 && board[i][j-1] == 'O' {
				uniZip(father, ID, ID-1, dumpPos)
			}
		}

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}

			//如果根节点不是 dumpPos ，表示不链接到边界
			if findFather(father, i*n+j) != dumpPos {
				board[i][j] = 'X'
			}

		}

	}
	return
}

// https://leetcode-cn.com/problems/surrounded-regions/solution/24ms6mb-he-bing-ji-he-jie-jue-130-bei-wei-rao-de-q/

/*
 利用bfs遍历 ，申请内存错误
*/
func solve2(board [][]byte) {

	rows := len(board)
	if rows == 0 {
		return
	}

	cols := len(board[0])
	visited := make([]bool, rows*cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			if board[r][c] == 'O' && (!visited[r*cols+c]) {
				atsize, points := bfs2(board, visited, r*cols+c, rows, cols)
				if !atsize {
					fill(board, points, cols)
				}

			}

		}
	}
}

/*
atside 是否包含边界
points 遍历得到的点
*/
func bfs2(board [][]byte, visited []bool, start int, rows, cols int) (atside bool, points []int) {
	atside = false
	cur := []int{start}
	for len(cur) > 0 {

		new := make([]int, 0)
		for _, point := range cur {
			visited[point] = true

			if !atside {
				points = append(points, point) //记录用于填充
			}

			crow := point / cols
			ccol := point % cols
			if crow == 0 || crow == rows-1 || ccol == 0 || ccol == cols-1 {
				atside = true
			}

			//向左边
			if ccol > 0 && (!visited[crow*cols+(ccol-1)]) && board[crow][ccol-1] == 'O' {
				new = append(new, crow*cols+(ccol-1))
			}
			//向右边
			if ccol < cols-1 && (!visited[crow*cols+(ccol+1)]) && board[crow][ccol+1] == 'O' {
				new = append(new, crow*cols+(ccol+1))
			}

			//向下
			if crow < rows-1 && (!visited[(crow+1)*cols+ccol]) && board[crow+1][ccol] == 'O' {
				new = append(new, (crow+1)*cols+ccol)
			}

			//向上
			if crow > 0 && (!visited[(crow-1)*cols+ccol]) && board[crow-1][ccol] == 'O' {
				new = append(new, (crow-1)*cols+ccol)
			}

		}

		cur = new

	}

	return
}

func fill(board [][]byte, points []int, cols int) {

	for _, point := range points {
		board[point/cols][point%cols] = 'X'
	}

	return
}

/*
将边界的O相连的O先变为Y
*/
func dfs(board [][]byte, start int, rows, cols int) {
	cur := []int{start}
	board[start/cols][start%cols] = 'Y'
	dir := []int{0, 1, 0, -1, 0}
	for len(cur) > 0 {

		// visited[point] = true
		point := cur[0]
		cur = cur[1:]
		crow := point / cols
		ccol := point % cols

		for i := 0; i < 4; i++ {
			newr := crow + dir[i]
			newc := ccol + dir[i+1]

			if 0 <= newr && newr < rows && 0 <= newc && newc < cols && board[newr][newc] == 'O' {
				cur = append(cur, newr*cols+newc)
				board[newr][newc] = 'Y'
			}
		}

	}

	return
}

/*
只能dfs，bfs会alloc fail
*/
func solve(board [][]byte) {

	rows := len(board)
	if rows == 0 {
		return
	}

	cols := len(board[0])
	// visited := make([]bool, rows*cols)

	//将边界的O相连的O先变为Y,变为Y，同时也标记了这个点被访问过
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			if (r == 0 || r == rows-1 || c == 0 || c == cols-1) && board[r][c] == 'O' {
				dfs(board, r*cols+c, rows, cols)
			}

		}
	}

	//剩余的O变为X
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}

		}
	}

	//剩余的Y变为O
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			if board[r][c] == 'Y' {
				board[r][c] = 'O'
			}

		}
	}
	return
}
