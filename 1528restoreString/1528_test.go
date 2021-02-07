/*
 * @Description: 1528
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-06 10:24:45
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-06 11:03:10
 * @FilePath: \1528restoreString\1528_test.go
 */

package leetcode

import "testing"

func Test_restoreString(t *testing.T) {
	type args struct {
		s       string
		indices []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}}, "leetcode"},
		{"2", args{"abc", []int{0, 1, 2}}, "abc"},
		{"3", args{"aiohn", []int{3, 1, 4, 2, 0}}, "nihao"},
		{"4", args{"aaiougrt", []int{4, 0, 2, 6, 7, 3, 1, 5}}, "arigatou"},
		{"5", args{"art", []int{1, 0, 2}}, "rat"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreString(tt.args.s, tt.args.indices); got != tt.want {
				t.Errorf("restoreString() = %v, want %v", got, tt.want)
			}
		})
	}
}
