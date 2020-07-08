package main

import (
	"reflect"
	"testing"
)

func Test_divingBoard(t *testing.T) {
	type args struct {
		shorter int
		longer  int
		k       int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{shorter: 1, longer: 1, k: 0}, nil},
		{"2", args{shorter: 1, longer: 2, k: 3}, []int{3, 4, 5, 6}},
		{"3", args{shorter: 1, longer: 1, k: 3}, []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divingBoard(tt.args.shorter, tt.args.longer, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divingBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
