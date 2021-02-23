/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2020-07-01 12:44:30
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-23 16:17:45
 * @FilePath: \74\74_test.go
 */
package leetcode

import "testing"

func Test_searchMatrix(t *testing.T) {

	matrix := make([][]int, 3)
	matrix[0] = []int{1, 3, 5, 7}
	matrix[1] = []int{10, 11, 16, 20}
	matrix[2] = []int{23, 30, 34, 50}

	matrix1 := make([][]int, 1)
	matrix1[0] = []int{1}

	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "1", args: args{matrix: matrix, target: 3}, want: true},
		{name: "2", args: args{matrix: matrix, target: 13}, want: false},
		{name: "2", args: args{matrix: matrix1, target: 1}, want: true},
		{name: "2", args: args{matrix: matrix1, target: 13}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
