/*
 * @Description: 557
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-05 16:14:38
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-05 17:29:03
 * @FilePath: \557reverseWords\557_test.go
 */

package leetcode

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"Let's take LeetCode contest"}, "s'teL ekat edoCteeL tsetnoc"},
		{"2", args{""}, ""},
		{"3", args{"w"}, "w"},
		{"4", args{"wa"}, "aw"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
