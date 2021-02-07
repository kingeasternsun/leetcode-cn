/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-03 09:26:28
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-03 09:40:45
 * @FilePath: \490medianSlidingWindow\490_test.go
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_medianSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"1", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3}, []float64{1, -1, -1, 3, 5, 6}},
		{"2", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 1}, []float64{1.00000, 3.00000, -1.00000, -3.00000, 5.00000, 3.00000, 6.00000, 7.00000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("medianSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
