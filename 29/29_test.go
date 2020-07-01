package main

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{10, 3}, 3},
		{"2", args{1, 1}, 1},
		{"3", args{1, 2}, 0},
		{"4", args{2, 2}, 1},
		{"5", args{7, -3}, -2},
		{"6", args{-7, 3}, -2},
		{"7", args{-2, 1}, -2},
		{"8", args{2, -1}, -2},
		{"9", args{-2, -1}, 2},
		{"10", args{-1, -1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
