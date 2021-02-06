/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-05 09:12:56
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-05 10:08:11
 * @FilePath: \1208equalSubstring\1208_test.go
 */

package leetcode

import "testing"

/*
"krpgjbjjznpzdfy"
"nxargkbydxmsgby"
14
*/
func Test_equalSubstring(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{s: "abcd", t: "bcdf", maxCost: 3}, 3},
		{"2", args{s: "abcd", t: "cdef", maxCost: 3}, 1},
		{"3", args{s: "abcd", t: "acde", maxCost: 0}, 1},
		{"4", args{s: "krpgjbjjznpzdfy", t: "nxargkbydxmsgby", maxCost: 14}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSubstringBinary(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}

			if got := equalSubstring(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
