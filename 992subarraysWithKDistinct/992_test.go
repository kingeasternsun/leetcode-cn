/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-09 21:06:28
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-09 21:47:10
 * @FilePath: /992subarraysWithKDistinct/992_test.go
 */
package leetcode

import "testing"

func Test_subarraysWithKDistinct(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 1, 2, 3}, 2}, 7},
		{"2", args{[]int{1, 2, 1, 3, 4}, 3}, 3},
		{"3", args{[]int{1, 1, 1}, 1}, 6},
		{"4", args{[]int{1, 1, 1, 1}, 2}, 0},
		{"5", args{[]int{1, 1, 1, 1}, 1}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysWithKDistinct(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
