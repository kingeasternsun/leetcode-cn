/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-05-01 19:47:12
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-05-01 20:16:24
 * @FilePath: /leetcode-cn/offer10-1/10_test.go
 */
package leetcode

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {"1", args{2}, 1},
		{"2", args{5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
