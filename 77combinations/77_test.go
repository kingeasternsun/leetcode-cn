package main

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}

	want0 := make([][]int, 0)
	want1 := [][]int{
		[]int{1},
		[]int{2},
		[]int{3},
	}
	want2 := [][]int{
		[]int{1, 2},
		[]int{1, 3},
		[]int{2, 3},
	}
	want3 := [][]int{
		[]int{1, 2, 3},
	}
	want4 := make([][]int, 0)
	want54 := [][]int{
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3, 5},
		[]int{1, 2, 4, 5},
		[]int{1, 3, 4, 5},
		[]int{2, 3, 4, 5},
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"0", args{3, 0}, want0},
		{"1", args{3, 1}, want1},
		{"2", args{3, 2}, want2},
		{"3", args{3, 3}, want3},
		{"4", args{3, 4}, want4},
		{"54", args{5, 4}, want54},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
