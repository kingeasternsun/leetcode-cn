/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-04 10:03:59
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-04 10:10:00
 * @FilePath: \643findMaxAverage\643_test.go
 */

package leetcode

import "testing"

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 12, -5, -6, 50, 3}, 4}, 12.75},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
