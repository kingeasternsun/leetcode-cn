/*
 * @Author: your name
 * @Date: 2020-03-16 08:55:56
 * @LastEditTime: 2020-03-16 09:05:15
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \leetcode\interview0106\compress-string-lcci.go
 */

package main

import "testing"

func Test_compressString(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"aabcccccaaa"}, "a2b1c5a3"},
		{"2", args{"abbccd"}, "abbccd"},
		{"3", args{"a"}, "a"},
		{"4", args{"aa"}, "aa"},
		{"5", args{"aaa"}, "a3"},
		{"6", args{"ab"}, "ab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString(tt.args.S); got != tt.want {
				t.Errorf("compressString() = %v, want %v", got, tt.want)
			}
		})
	}
}
