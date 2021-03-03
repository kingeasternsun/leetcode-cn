/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-27 14:30:22
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-03 15:26:24
 * @FilePath: \395longestSubstring\longestSubstring_test.go
 */
package main

import "testing"

func Test_longestSubstring(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{"aaabb", 3}, 3},
		{"2", args{"ababbc", 2}, 5},
		{"3", args{"aaabb", 4}, 0},
		{"4", args{"aaabbb", 3}, 6},
		{"5", args{"a", 1}, 1},
		{"6", args{"ab", 1}, 2},
		{"7", args{"aa", 2}, 2},
		{"8", args{"abab", 2}, 4},
		{"9", args{"abcabc", 2}, 6},
		{"10", args{"aabbcc", 2}, 6},
		{"10", args{"weitong", 2}, 0},
		{"11", args{"abcde", 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstring(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("longestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
