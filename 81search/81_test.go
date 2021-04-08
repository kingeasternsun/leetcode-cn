/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-04-07 23:06:52
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-04-07 23:39:42
 * @FilePath: /81search/81_test.go
 */
package leetcode

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// // TODO: Add test cases.
		{"1", args{[]int{2, 5, 6, 0, 0, 1, 2}, 0}, true},
		{"2", args{[]int{2, 5, 6, 0, 0, 1, 2}, 3}, false},
		{"3", args{[]int{1}, 0}, false},
		{"4", args{[]int{1, 0, 1, 1, 1}, 0}, true},
		{"5", args{[]int{1, 1, 1, 1, 1}, 1}, true},
		{"6", args{[]int{1, 1, 1, 1, 1, 1}, 1}, true},
		{"7", args{[]int{0, 0}, 1}, false},
		{"8", args{[]int{0, 0, 0}, 1}, false},
		{"9", args{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2}, true},
		{"10", args{[]int{1, 3}, 2}, false},
		{"11", args{[]int{3, 1}, 1}, true},
		{"12", args{[]int{3, 1, 1}, 1}, true},
		{"13", args{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1}, true},
		{"14", args{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 13}, true},
		{"15", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, true},
		// {"15"}
		// [4,5,6,7,0,1,2]
		// 0
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
