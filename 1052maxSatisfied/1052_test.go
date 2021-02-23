/*
 * @Description: 1052
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-23 09:04:46
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-23 09:31:07
 * @FilePath: \1052maxSatisfied\1052_test.go
 */

package leetcode

import "testing"

func Test_maxSatisfied(t *testing.T) {
	type args struct {
		customers []int
		grumpy    []int
		X         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{customers: []int{1, 0, 1, 2, 1, 1, 7, 5}, grumpy: []int{0, 1, 0, 1, 0, 1, 0, 1}, X: 3}, 16},
		{"2", args{customers: []int{2}, grumpy: []int{0}, X: 1}, 2},
		{"3", args{customers: []int{2}, grumpy: []int{1}, X: 1}, 2},
		{"4", args{customers: []int{2, 3}, grumpy: []int{0, 1}, X: 1}, 5},
		{"5", args{customers: []int{4, 10, 10}, grumpy: []int{1, 1, 0}, X: 2}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfied(tt.args.customers, tt.args.grumpy, tt.args.X); got != tt.want {
				t.Errorf("maxSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
