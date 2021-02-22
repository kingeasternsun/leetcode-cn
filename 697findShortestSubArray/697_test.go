/*
 * @Description: 697
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-20 10:50:19
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-20 11:48:27
 * @FilePath: \697findShortestSubArray\697_test.go
 */

package leetcode

import "testing"

func Test_findShortestSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 2, 3, 1}}, 2},
		{"2", args{[]int{1, 2, 2, 3, 1, 4, 2}}, 6},
		{"3", args{[]int{1, 2}}, 1},
		{"4", args{[]int{2}}, 1},
		{"5", args{[]int{1, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray(tt.args.nums); got != tt.want {
				t.Errorf("findShortestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
