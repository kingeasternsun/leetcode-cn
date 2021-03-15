/*
 * @Description: 1047
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-09 09:13:18
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-09 10:02:52
 * @FilePath: \1047removeDuplicates\1047_test.go
 */

package letcode

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"abbaca"}, "ca"},
		{"2", args{"abba"}, ""},
		{"3", args{"abbb"}, "ab"},
		{"4", args{"bbb"}, "b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.S); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
