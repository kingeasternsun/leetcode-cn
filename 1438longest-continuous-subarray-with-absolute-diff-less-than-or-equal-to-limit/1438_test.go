package main

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{nums: []int{8, 2, 4, 7}, limit: 4}, 2},
		{"2", args{nums: []int{10, 1, 2, 4, 7, 2}, limit: 5}, 4},
		{"3", args{nums: []int{4, 2, 2, 2, 4, 4, 2, 2}, limit: 0}, 3},
		//
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
