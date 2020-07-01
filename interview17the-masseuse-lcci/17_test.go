package main

import "testing"

func Test_massage(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 3, 1}}, 4},
		{"2", args{[]int{2, 7, 9, 3, 1}}, 12},
		{"3", args{[]int{2, 1, 4, 5, 3, 1, 1, 3}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := massage(tt.args.nums); got != tt.want {
				t.Errorf("massage() = %v, want %v", got, tt.want)
			}
		})
	}
}
