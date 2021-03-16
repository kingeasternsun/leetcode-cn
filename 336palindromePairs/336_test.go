/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-16 16:10:57
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-16 16:55:39
 * @FilePath: /336palindromePairs/336_test.go
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_palindromePairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"1", args{[]string{"abcd", "dcba", "lls", "s", "sssll"}}, [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromePairs(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("palindromePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
