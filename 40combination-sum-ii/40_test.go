package main

import (
	"reflect"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}

	want1 := [][]int{
		[]int{1, 1, 6},
		[]int{1, 2, 5},
		[]int{1, 7},
		[]int{2, 6},
	}
	want2 := [][]int{
		[]int{1, 2, 2},
		[]int{5},
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"1", args{candidates: []int{10, 1, 2, 7, 6, 1, 5}, target: 8}, want1},
		{"2", args{candidates: []int{2, 5, 2, 1, 2}, target: 5}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
