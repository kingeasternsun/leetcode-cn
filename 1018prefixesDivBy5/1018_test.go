package leetcode

import (
	"reflect"
	"testing"
)

func Test_prefixesDivBy5(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
		{"1", args{A: []int{0, 1, 1}}, []bool{true, false, false}},
		{"2", args{A: []int{1, 1, 1}}, []bool{false, false, false}},
		{"3", args{A: []int{0, 1, 1, 1, 1, 1}}, []bool{true, false, false, false, true, false}},
		{"4", args{A: []int{1, 1, 1, 0, 1}}, []bool{false, false, false, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixesDivBy5(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixesDivBy5() = %v, want %v", got, tt.want)
			}
		})
	}
}
