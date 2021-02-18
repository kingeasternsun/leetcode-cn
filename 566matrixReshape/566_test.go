/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-17 22:05:58
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-17 22:09:42
 * @FilePath: /566matrixReshape/566_test.go
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	nums1 := [][]int{
		{1, 2},
		{3, 4},
	}
	res1 := [][]int{
		{1, 2, 3, 4},
	}

	nums3 := [][]int{
		{1},
	}
	type args struct {
		nums [][]int
		r    int
		c    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"1", args{nums: nums1, r: 1, c: 4}, res1},
		{"2", args{nums: nums1, r: 2, c: 4}, nums1},
		{"3", args{nums: nums3, r: 2, c: 4}, nums3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixReshape(tt.args.nums, tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixReshape() = %v, want %v", got, tt.want)
			}
		})
	}
}
