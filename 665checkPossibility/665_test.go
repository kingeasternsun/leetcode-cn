/*
 * @Description: 665
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-07 09:18:33
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-07 09:45:41
 * @FilePath: \665checkPossibility\665_test.go
 */

package leetcode

import "testing"

func Test_checkPossibility(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{[]int{4, 2, 3}}, true},
		{"2", args{[]int{4, 2, 1}}, false},
		{"3", args{[]int{4, 2}}, true},
		{"4", args{[]int{2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
