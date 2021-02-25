/*
 * @Description: 27
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-25 09:24:07
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-25 09:30:55
 * @FilePath: \27removeElement\27_test.go
 */

package leetcode

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{3, 2, 2, 3}, 3}, 2},
		{"2", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, 5},
		{"3", args{[]int{2}, 2}, 0},
		{"4", args{[]int{3}, 2}, 1},
		{"5", args{[]int{2, 2}, 2}, 0},
		{"6", args{[]int{2, 3}, 2}, 1},
		{"7", args{[]int{3, 2}, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
