package main

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{nums: []int{2, 0, 2, 1, 1, 0}},
			want: []int{0, 0, 1, 1, 2, 2}},
		{name: "2", args: args{nums: []int{2}},
			want: []int{2}},
		{name: "3", args: args{nums: []int{1}},
			want: []int{1}},
		{name: "3", args: args{nums: []int{0}},
			want: []int{0}},
		{name: "2", args: args{nums: []int{2, 1}},
			want: []int{1, 2}},
		{name: "3", args: args{nums: []int{1, 0}},
			want: []int{0, 1}},
		{name: "3", args: args{nums: []int{0, 2}},
			want: []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if sortColors(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("longestPalindrome() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
