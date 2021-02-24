/*
 * @Description: 832
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-24 09:07:07
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-24 09:25:56
 * @FilePath: \832flipAndInvertImage\832_test.go
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	m1 := [][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}
	r1 := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	m3 := [][]int{
		{1},
		{0},
	}
	r3 := [][]int{
		{0},
		{1},
	}
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"1", args{m1}, r1},
		{"3", args{m3}, r3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipAndInvertImage(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
