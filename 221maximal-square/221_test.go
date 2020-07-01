package main

import "testing"

func Test_maximalSquare(t *testing.T) {

	m := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}
	m2 := [][]byte{
		[]byte{'0'},
	}
	m3 := [][]byte{
		[]byte{'1'},
	}
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{matrix: m}, 4},
		{"2", args{matrix: m2}, 0},
		{"3", args{matrix: m3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
