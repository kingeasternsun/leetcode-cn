/*
 * @Description: 867
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-25 08:57:04
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-25 09:10:35
 * @FilePath: \867transpose\867_test.go
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_transpose(t *testing.T) {
	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res1 := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}

	m2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	res2 := [][]int{
		{1, 4},
		{2, 5},
		{3, 6},
	}

	m3 := [][]int{
		{1, 2, 3},
	}
	res3 := [][]int{
		{1},
		{2},
		{3},
	}

	m4 := [][]int{
		{1},
		{2},
		{3},
	}

	res4 := [][]int{
		{1, 2, 3},
	}

	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"1", args{m1}, res1},
		{"2", args{m2}, res2},
		{"3", args{m3}, res3},
		{"4", args{m4}, res4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transpose(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
