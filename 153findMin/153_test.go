package main

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{nums: []int{3, 4, 5, 1, 2}}, 1},
		{"2", args{nums: []int{4, 5, 6, 7, 0, 1, 2}}, 0},
		{"3", args{nums: []int{6, 7, 0, 1, 2, 3, 4, 5}}, 0},
		{"4", args{nums: []int{6, 7, 0, 1, 2, 3, 4, 5}}, 0},
		{"5", args{nums: []int{6, 7}}, 6},
		{"6", args{nums: []int{9, 7}}, 7},
		{"7", args{nums: []int{3, 1, 2}}, 1},
		{"8", args{nums: []int{4, 5, 1, 2, 3}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
