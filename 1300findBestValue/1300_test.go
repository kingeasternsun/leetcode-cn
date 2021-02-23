/*
 * @Description:1300
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-23 17:12:38
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-23 19:25:20
 * @FilePath: \1300findBestValue\1300_test.go
 */

package leetcode

import "testing"

func Test_findBestValue(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{4, 9, 3}, 10}, 3},
		{"2", args{[]int{2, 3, 5}, 10}, 5},
		{"3", args{[]int{60864, 25176, 27249, 21296, 20204}, 56803}, 11361},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBestValue(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("findBestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
