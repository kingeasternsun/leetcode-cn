/*
 * @Description: 338
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-03 11:00:27
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-03 11:05:21
 * @FilePath: \338countBits\338_test.go
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{5}, []int{0, 1, 1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
