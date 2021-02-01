package leetcode

import (
	"reflect"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{A: []int{1, 1}, B: []int{2, 2}}, []int{1, 2}},
		{"1r", args{A: []int{2, 2}, B: []int{1, 1}}, []int{2, 1}},
		{"2", args{A: []int{1, 2}, B: []int{2, 3}}, []int{1, 2}},
		{"3", args{A: []int{2}, B: []int{1, 3}}, []int{2, 3}},

		{"4", args{A: []int{1, 2, 5}, B: []int{2, 4}}, []int{5, 4}},
		{"4r", args{A: []int{2, 4}, B: []int{1, 2, 5}}, []int{4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwap(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
			if got := fairCandySwapHash(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
