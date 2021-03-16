/*
 * @Description: 866
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-16 11:24:32
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-16 14:17:27
 * @FilePath: /866primePalindrome/866_test.go
 */
package leetcode

import "testing"

func Test_primePalindrome(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{13}, 101},
		{"2", args{6}, 7},
		{"3", args{8}, 11},
		{"4", args{130}, 131},
		{"5", args{1}, 2},
		{"6", args{2}, 2},
		{"7", args{11}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := primePalindrome(tt.args.N); got != tt.want {
				t.Errorf("primePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
