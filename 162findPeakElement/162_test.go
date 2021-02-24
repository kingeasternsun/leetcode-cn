/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-24 11:30:48
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-24 11:55:11
 * @FilePath: \162findPeakElement\162_test.go
 */

package leetcode

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 3, 1}}, 2},
		{"2", args{[]int{1, 2, 1, 3, 5, 6, 4}}, 5},
		{"3", args{[]int{1}}, 0},
		{"4", args{[]int{2, 1}}, 0},
		{"5", args{[]int{3, 2, 1}}, 0},
		{"6", args{[]int{1, 2}}, 1},
		{"7", args{[]int{1, 2, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
