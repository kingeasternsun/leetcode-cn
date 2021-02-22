/*
 * @Description: 766
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-22 09:53:17
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-22 10:09:24
 * @FilePath: \766isToeplitzMatrix\766_test.go
 */

package leetcode

import "testing"

func Test_isToeplitzMatrix(t *testing.T) {
	matrix1 := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}
	matrix2 := [][]int{
		{1, 2},
		{2, 2},
	}

	matrix3 := [][]int{
		{1, 2},
	}
	matrix4 := [][]int{
		{1},
	}
	matrix5 := [][]int{
		{1},
		{2},
	}
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{matrix1}, true},
		{"2", args{matrix2}, false},
		{"3", args{matrix3}, true},
		{"4", args{matrix4}, true},
		{"5", args{matrix5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isToeplitzMatrix(tt.args.matrix); got != tt.want {
				t.Errorf("isToeplitzMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
