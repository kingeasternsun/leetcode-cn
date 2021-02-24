/*
 * @Description: 1011
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-24 10:24:31
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-24 10:53:43
 * @FilePath: \1011shipWithinDays\1011_test.go
 */

package leetcode

import "testing"

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		D       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5}, 15},
		{"2", args{[]int{3, 2, 2, 4, 1, 4}, 3}, 6},
		{"3", args{[]int{1, 2, 3, 1, 1}, 4}, 3},
		{"4", args{[]int{1, 2}, 1}, 3},
		{"5", args{[]int{1}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.args.weights, tt.args.D); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
