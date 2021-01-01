/*
 * @Author: your name
 * @Date: 2020-03-14 21:34:31
 * @LastEditTime: 2020-03-14 21:40:15
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \leetcode\300\300_test.go
 */
package leetcode

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{nums: []int{1, 2, 3}}, 3},
		{"2", args{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
		{"3", args{nums: []int{10}}, 1},
		{"4", args{nums: []int{10, 10}}, 1},
		{"5", args{nums: []int{10, 9, 2, 5, 3, 4}}, 3},

		//[10,9,2,5,3,4]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
