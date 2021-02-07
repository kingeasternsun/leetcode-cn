/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-06 10:06:07
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-06 11:41:22
 * @FilePath: \1423maxScore\1423_test.go
 */

package leetcode

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 3, 4, 5, 6, 1}, 3}, 12},
		{"2", args{[]int{2, 2, 2}, 2}, 4},
		{"3", args{[]int{9, 7, 7, 9, 7, 7, 9}, 7}, 55},
		{"4", args{[]int{1, 1000, 1}, 1}, 1},
		{"5", args{[]int{1, 79, 80, 1, 1, 1, 200, 1}, 3}, 202},
		{"6", args{[]int{1, 1000, 2}, 1}, 2},
		{"7", args{[]int{3, 1000, 2}, 1}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
			if got := maxScore1(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
