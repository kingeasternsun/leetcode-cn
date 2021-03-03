/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-02 11:26:04
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-02 11:55:50
 * @FilePath: \304SumRegion\304_test.go
 */

package leetcode

import "testing"

func TestNumMatrix_SumRegion(t *testing.T) {

	m := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	obj := Constructor(m)
	type args struct {
		row1 int
		col1 int
		row2 int
		col2 int
	}
	tests := []struct {
		name string
		this *NumMatrix
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", &obj, args{2, 1, 4, 3}, 8},
		{"2", &obj, args{1, 1, 2, 2}, 11},
		{"3", &obj, args{1, 2, 2, 4}, 12},
		{"4", &obj, args{0, 0, 1, 1}, 14},
		{"5", &obj, args{1, 0, 2, 0}, 6},
		{"6", &obj, args{0, 1, 0, 3}, 5},
		{"7", &obj, args{1, 1, 1, 1}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.SumRegion(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2); got != tt.want {
				t.Errorf("NumMatrix.SumRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}
