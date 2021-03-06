/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-06 13:53:39
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-06 13:54:09
 * @FilePath: /503nextGreaterElements/503_test.go
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 1}}, []int{2, -1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
