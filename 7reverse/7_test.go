/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-05-03 21:47:11
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-05-03 21:53:06
 * @FilePath: /7reverse/7_test.go
 */
package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{1}, 1},
		{"2", args{123}, 321},
		{"3", args{-123}, -321},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
