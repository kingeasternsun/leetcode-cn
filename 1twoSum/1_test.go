/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-17 22:17:58
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-17 22:19:02
 * @FilePath: /1twoSum/1_test.go
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"2", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"3", args{[]int{3, 3}, 6}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
