/*
 * @Description: 867
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-25 08:57:04
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-25 09:11:18
 * @FilePath: \867transpose\867.go
 */

package leetcode

func transpose(matrix [][]int) [][]int {

	rows := len(matrix)
	if rows == 0 {
		return matrix
	}
	cols := len(matrix[0])
	if cols == 0 {
		return matrix
	}

	if rows == cols {
		return inplace(matrix, rows, cols)
	}

	res := make([][]int, cols)
	for i := 0; i < cols; i++ {
		res[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res[j][i] = matrix[i][j]
		}
	}

	return res
}

func inplace(matrix [][]int, rows, cols int) [][]int {
	for i := 0; i < rows; i++ {
		for j := i + 1; j < cols; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	return matrix
}
