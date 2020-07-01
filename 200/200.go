package main

/*
https://leetcode-cn.com/problems/number-of-islands/
200. 岛屿数量
*/

func findFatherZip(father []int, i int) int {
	j := i
	for father[i] != i {
		i = father[i]
	}

	father[j] = i
	return i

}

// i对应的组链接到j对应的组
// 如果提前知道两个组的大小，就将小的组链接到大的组上
func uniZip(father []int, i, j int) int {
	i = findFatherZip(father, i)
	j = findFatherZip(father, j)
	if i == j {
		return 0
	}
	father[i] = j
	return 1
}
func numIslands1(grid [][]byte) int {

	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])

	if n == 0 {
		return 0
	}

	father := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		father[i] = i
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}

			ID := i*n + j
			//向上
			if i > 0 && grid[i-1][j] == '1' {
				uniZip(father, ID, ID-n)
			}
			// 向左
			if j > 0 && grid[i][j-1] == '1' {
				uniZip(father, ID, ID-1)
			}
		}

	}

	var count int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}

			if father[i*n+j] == i*n+j {
				count++
			}

		}

	}
	return count
}

func numIslands(grid [][]byte) int {

	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])

	if n == 0 {
		return 0
	}

	father := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		father[i] = i
	}

	var count int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}

			// 先把当前位置作为单独的一个岛屿， 后面如果合并了就再减去
			count++
			ID := i*n + j
			//向上
			if i > 0 && grid[i-1][j] == '1' {
				count = count - uniZip(father, ID, ID-n)
			}
			// 向左
			if j > 0 && grid[i][j-1] == '1' {
				count = count - uniZip(father, ID, ID-1)
			}
		}

	}

	return count
}
