package main

import "testing"

func Test_minReorder(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}

	con1 := [][]int{
		[]int{0, 1},
		[]int{1, 3},
		[]int{2, 3},
		[]int{4, 0},
		[]int{4, 5},
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{6, con1}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minReorder(tt.args.n, tt.args.connections); got != tt.want {
				t.Errorf("minReorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
