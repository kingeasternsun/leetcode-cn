package main

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	r1 := [][]int{
		[]int{-5, -4, -3, 1},
	}

	r2 := [][]int{
		[]int{-4, 0, 1, 2},
		[]int{-1, -1, 0, 1},
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"1", args{nums: []int{1, -2, -5, -4, -3, 3, 3, 5}, target: -11}, r1},
		{"2", args{nums: []int{-1, 0, 1, 2, -1, -4}, target: -1}, r2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
