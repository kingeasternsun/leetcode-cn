/*
 * @Description:704
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-23 15:20:25
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-23 15:23:50
 * @FilePath: \704search\704_test.go
 */

package leetcode

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{-1, 0, 3, 5, 9, 12}, 9}, 4},
		{"2", args{[]int{-1, 0, 3, 5, 9, 12}, 2}, -1},
		{"3", args{[]int{-1, 0, 3, 5, 9, 12}, -2}, -1},
		{"4", args{[]int{-1, 0, 3, 5, 9, 12}, 20}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
