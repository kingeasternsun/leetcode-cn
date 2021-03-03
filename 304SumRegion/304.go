/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-02 11:26:04
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-02 11:56:48
 * @FilePath: \304SumRegion\304.go
 */
package leetcode

type NumMatrix struct {
	Sum    [][]int
	Matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{}
	}

	sum := make([][]int, len(matrix))
	for i := range sum {
		sum[i] = make([]int, len(matrix[0]))
	}

	sum[0][0] = matrix[0][0]
	for i := 1; i < len(matrix[0]); i++ {
		sum[0][i] = sum[0][i-1] + matrix[0][i]
	}

	for i := 1; i < len(matrix); i++ {
		sum[i][0] = sum[i-1][0] + matrix[i][0]
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i][j]
		}
	}

	return NumMatrix{Sum: sum, Matrix: matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	switch {
	case row1 == row2 && col1 == col2:
		return this.Matrix[row1][col2]
	case row1 == 0 && col1 == 0:
		return this.Sum[row2][col2]
	case row1 == 0:
		return this.Sum[row2][col2] - this.Sum[row2][col1-1]
	case col1 == 0:
		return this.Sum[row2][col2] - this.Sum[row1-1][col2]
	default:
		return this.Sum[row2][col2] - this.Sum[row2][col1-1] - this.Sum[row1-1][col2] + this.Sum[row1-1][col1-1]
	}

}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
