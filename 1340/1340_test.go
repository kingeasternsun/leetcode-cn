//1340. 跳跃游戏 V

package main

import "testing"

func Test_maxJumps(t *testing.T) {
	type args struct {
		arr []int
		d   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{arr: []int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, d: 2}, want: 4},
		{name: "1", args: args{arr: []int{3, 3, 3, 3, 3}, d: 3}, want: 1},
		{name: "1", args: args{arr: []int{7, 6, 5, 4, 3, 2, 1}, d: 1}, want: 7},
		{name: "1", args: args{arr: []int{7, 1, 7, 1, 7, 1}, d: 2}, want: 2},
		{name: "1", args: args{arr: []int{66}, d: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxJumps(tt.args.arr, tt.args.d); got != tt.want {
				t.Errorf("maxJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
