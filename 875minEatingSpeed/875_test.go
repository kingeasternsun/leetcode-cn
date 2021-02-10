/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-10 15:23:11
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-10 15:42:20
 * @FilePath: /875minEatingSpeed/875_test.go
 */
package leetcode

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		H     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{3, 6, 7, 11}, 8}, 4},
		{"2", args{[]int{30, 11, 23, 4, 20}, 5}, 30},
		{"3", args{[]int{30, 11, 23, 4, 20}, 6}, 23},
		{"4", args{[]int{2, 2}, 2}, 2}, //导致最后 left：0， right：1， mid得到0
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.H); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
