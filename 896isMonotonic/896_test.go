/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-28 13:55:19
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-28 13:56:03
 * @FilePath: /896isMonotonic/896_test.go
 */
package leetcode

import "testing"

func Test_isMonotonic(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1}}, true},
		{"2", args{[]int{1, 2}}, true},
		{"3", args{[]int{1, 2, 1}}, false},
		{"4", args{[]int{2, 1, 2}}, false},
		{"5", args{[]int{2, 2, 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.args.A); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
