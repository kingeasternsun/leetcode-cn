package main

import "testing"

func Test_kthSmallest(t *testing.T) {

	matrix1 := [][]int{
		[]int{01, 05, 9},
		[]int{10, 11, 13},
		[]int{12, 13, 15},
	}

	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{matrix1, 1}, 1},
		{"2", args{matrix1, 2}, 5},
		{"3", args{matrix1, 3}, 9},
		{"8", args{matrix1, 8}, 13},
		{"9", args{matrix1, 9}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
