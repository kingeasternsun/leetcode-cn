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

func solve(board [][]byte) {

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


https://leetcode-cn.com/problems/surrounded-regions/solution/24ms6mb-he-bing-ji-he-jie-jue-130-bei-wei-rao-de-q/