/*
 * @Description: 766
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-22 09:53:17
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-22 10:05:53
 * @FilePath: \766isToeplitzMatrix\766.go
 */
package leetcode

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) <= 1 {
		return true
	}

	if len(matrix[0]) == 1 {
		return true
	}

	for rowID := 0; rowID < len(matrix); rowID++ {

		x := rowID
		y := 0
		for x < len(matrix) && y < len(matrix[0]) {
			if matrix[x][y] != matrix[rowID][0] {
				return false
			}
			x++
			y++
		}
	}

	for colID := 1; colID < len(matrix[0]); colID++ {

		x := 0
		y := colID
		for x < len(matrix) && y < len(matrix[0]) {
			if matrix[x][y] != matrix[0][colID] {
				return false
			}
			x++
			y++
		}
	}

	return true
}
