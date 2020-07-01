package main

import "testing"

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}

	g1 := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	g2 := [][]int{
		[]int{1},
		[]int{1},
		[]int{1},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{g1}, 7},
		{"2", args{g2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
