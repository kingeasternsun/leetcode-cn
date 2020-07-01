package main

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 8}, want: []int{3, 4}},
		{name: "12", args: args{nums: []int{5, 7, 7, 8, 8, 10}, target: 6}, want: []int{-1, -1}},
		{name: "13", args: args{nums: []int{8}, target: 6}, want: []int{-1, -1}},
		{name: "14", args: args{nums: []int{8}, target: 8}, want: []int{0, 0}},
		{name: "15", args: args{nums: []int{8, 8}, target: 8}, want: []int{0, 1}},
		{name: "16", args: args{nums: []int{8, 8, 8}, target: 8}, want: []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
