/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-05-02 20:29:04
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-05-02 20:39:36
 * @FilePath: /554leastBricks/554_test.go
 */
package leetcode

import "testing"

func Test_leastBricks(t *testing.T) {
	type args struct {
		wall [][]int
	}

	wall1 := [][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}
	wall2 := [][]int{{1}, {1}, {1}}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{wall1}, 2},
		{"2", args{wall2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastBricks(tt.args.wall); got != tt.want {
				t.Errorf("leastBricks() = %v, want %v", got, tt.want)
			}
		})
	}
}
