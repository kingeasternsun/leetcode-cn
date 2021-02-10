/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-10 13:21:58
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-10 13:22:29
 * @FilePath: /567checkInclusion/567_test.go
 */
package leetcode

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{s1: "ab", s2: "eidbaooo"}, true},
		{"2", args{s1: "ab", s2: "eidboaoo"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
