/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-04-12 21:26:49
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-04-12 21:58:29
 * @FilePath: /179largestNumber/179_test.go
 */
package leetcode

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{[]int{10, 2}}, "210"},
		{"2", args{[]int{20, 2}}, "220"},
		{"3", args{[]int{3, 30, 34, 5, 9}}, "9534330"},
		{"4", args{[]int{1, 2}}, "21"},
		{"5", args{[]int{10}}, "10"},
		{"6", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}}, "9876543210"},
		{"7", args{[]int{0, 0}}, "0"},
		{"8", args{[]int{0, 7, 0}}, "700"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
