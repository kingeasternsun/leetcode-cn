/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-15 22:52:32
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-15 22:53:51
 * @FilePath: /485findMaxConsecutiveOnes/485_test.go
 */
package leetcode

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 1, 0, 1, 1, 1}}, 3},
		{"2", args{[]int{1}}, 1},
		{"3", args{[]int{0}}, 0},
		{"4", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
