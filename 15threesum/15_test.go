/*
 * @Description:15
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2020-07-01 12:44:30
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-19 11:58:03
 * @FilePath: \15threesum\15_test.go
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_threeSumHash(t *testing.T) {

	res1 := [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{-1, 0, 1, 2, -1, -4}}, res1},
		{"2", args{[]int{0}}, [][]int{}},
		{"3", args{[]int{}}, [][]int{}},
		{"4", args{[]int{0, 0, 0}}, [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumHash(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSumHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
