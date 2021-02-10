/*
 * @Author: your name
 * @Date: 2021-02-08 11:59:22
 * @LastEditTime: 2021-02-08 12:00:42
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /978maxTurbulenceSize/978_test.go
 */
package leetcode

import "testing"

func Test_maxTurbulenceSize(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}}, 5},
		{"2", args{[]int{4, 8, 12, 16}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSize(tt.args.arr); got != tt.want {
				t.Errorf("maxTurbulenceSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
