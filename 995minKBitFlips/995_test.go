/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-18 19:48:10
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-18 20:07:19
 * @FilePath: /995minKBitFlips/995_test.go
 */
package leetcode

import "testing"

func Test_minKBitFlips(t *testing.T) {
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
		{"1", args{[]int{0, 1, 0}, 1}, 2},
		{"2", args{[]int{1, 1, 0}, 2}, -1},
		{"3", args{[]int{0, 0, 0, 1, 0, 1, 1, 0}, 3}, 3},
		{"4", args{[]int{0, 1, 1}, 2}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minKBitFlips(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("minKBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
