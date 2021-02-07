/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-05 17:36:11
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-06 09:52:18
 * @FilePath: \541reverseStr\541_test.go
 */

package leetcode

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"abcdefg", 2}, "bacdfeg"},
		{"2", args{"abcdef", 2}, "bacdfe"},
		{"3", args{"abcde", 2}, "bacde"},
		{"4", args{"abcd", 2}, "bacd"},
		{"5", args{"abcdefg", 8}, "gfedcba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

// "gfedcba"
