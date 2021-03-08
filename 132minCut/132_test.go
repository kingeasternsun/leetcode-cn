/*
 * @Description: 132
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-08 09:42:18
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-08 11:11:58
 * @FilePath: \132minCut\132_test.go
 */
package leetcode

import "testing"

func Test_minCut(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{"aab"}, 1},
		{"2", args{"a"}, 0},
		{"3", args{"ab"}, 1},
		{"4", args{"aba"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
