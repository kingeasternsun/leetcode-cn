/*
 * @Description: 20
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-25 09:33:49
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-25 09:40:12
 * @FilePath: \20isValid\20_test.go
 */

package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"2", args{"()[]{}"}, true},
		{"5", args{"{[]}"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
