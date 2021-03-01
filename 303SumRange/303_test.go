/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-01 09:25:13
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-01 09:29:01
 * @FilePath: /303SumRange/303_test.go
 */
package leetcode

import "testing"

func TestNumArray_SumRange(t *testing.T) {

	this := Constructor([]int{-2, 0, 3, -5, 2, -1})
	type fields struct {
		PreSum []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{i: 0, j: 2}, 1},
		{"2", args{i: 2, j: 5}, -1},
		{"3", args{i: 0, j: 5}, -3},
		{"4", args{i: 0, j: 0}, -2},
		{"5", args{i: 1, j: 1}, 0},
		{"6", args{i: 2, j: 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := this.SumRange(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("NumArray.SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
