package main

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{nums: []int{1, 2, 0}}, want: 3},
		{name: "2", args: args{nums: []int{3, 4, -1, 1}}, want: 2},
		{name: "3", args: args{nums: []int{7, 8, 9, 11, 12}}, want: 1},
		{name: "4", args: args{nums: []int{-1}}, want: 1},
		{name: "5", args: args{nums: []int{0}}, want: 1},
		{name: "6", args: args{nums: []int{1}}, want: 2},
		{name: "7", args: args{nums: []int{2}}, want: 1},
		{name: "8", args: args{nums: []int{0, -1}}, want: 1},
		{name: "9", args: args{nums: []int{1, -1}}, want: 2},
		{name: "10", args: args{nums: []int{2, -1}}, want: 1},
		{name: "11", args: args{nums: []int{2, 2}}, want: 1},
		{name: "11", args: args{nums: []int{3, 3, 3}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
