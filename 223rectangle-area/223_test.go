package main

import "testing"

func Test_computeArea(t *testing.T) {
	type args struct {
		A int
		B int
		C int
		D int
		E int
		F int
		G int
		H int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{-3, 0, 3, 4, 0, -1, 9, 2}, 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeArea(tt.args.A, tt.args.B, tt.args.C, tt.args.D, tt.args.E, tt.args.F, tt.args.G, tt.args.H); got != tt.want {
				t.Errorf("computeArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
