/*
 * @Description: 26
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-25 09:17:27
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-25 09:21:45
 * @FilePath: \26removeDuplicates\26_test.go
 */

package leetcode

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 1, 2}}, 2},
		{"2", args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
