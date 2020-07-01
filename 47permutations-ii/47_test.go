package main

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}

	r1 := [][]int{
		[]int{1},
	}
	r2 := [][]int{
		[]int{1, 2},
		[]int{2, 1},
	}
	r3 := [][]int{
		[]int{1, 1, 2},
		[]int{1, 2, 1},
		[]int{2, 1, 1},
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1}}, r1},
		{"2", args{[]int{1, 2}}, r2},
		{"3", args{[]int{1, 1, 2}}, r3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
