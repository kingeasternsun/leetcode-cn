package main

import "testing"

func Test_cherryPickup(t *testing.T) {
	type args struct {
		grid [][]int
	}
	g1 := [][]int{[]int{3, 1, 1}, []int{2, 5, 1}, []int{1, 5, 5}, []int{2, 1, 1}}
	g2 := [][]int{[]int{1, 0, 0, 0, 0, 0, 1}, []int{2, 0, 0, 0, 0, 3, 0}, []int{2, 0, 9, 0, 0, 0, 0}, []int{0, 3, 0, 5, 4, 0, 0}, []int{1, 0, 2, 3, 0, 0, 6}}
	g3 := [][]int{[]int{1, 0, 0, 3}, []int{0, 0, 0, 3}, []int{0, 0, 3, 3}, []int{9, 0, 3, 3}}
	g4 := [][]int{
		[]int{0, 8, 7, 10, 9, 10, 0, 9, 6},
		[]int{8, 7, 10, 8, 7, 4, 9, 6, 10},
		[]int{8, 1, 1, 5, 1, 5, 5, 1, 2},
		[]int{9, 4, 10, 8, 8, 1, 9, 5, 0},
		[]int{4, 3, 6, 10, 9, 2, 4, 8, 10},
		[]int{7, 3, 2, 8, 3, 3, 5, 9, 8},
		[]int{1, 2, 6, 5, 6, 2, 0, 10, 0},
	}

	// [[0,0,10,2,8,4,0],[7,9,3,5,4,8,3],[6,9,8,3,5,6,0],[0,4,1,1,9,3,7],[5,6,9,8,8,10,10],[9,2,9,7,4,8,3],[1,6,1,2,0,9,9]]
	g5 := [][]int{
		[]int{0, 0, 10, 2, 8, 4, 0},
		[]int{7, 9, 3, 5, 4, 8, 3},
		[]int{6, 9, 8, 3, 5, 6, 0},
		[]int{0, 4, 1, 1, 9, 3, 7},
		[]int{5, 6, 9, 8, 8, 10, 10},
		[]int{9, 2, 9, 7, 4, 8, 3},
		[]int{1, 6, 1, 2, 0, 9, 9},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{g1}, 24},
		{"2", args{g2}, 28},
		{"3", args{g3}, 22},
		{"4", args{g4}, 96},
		{"5", args{g5}, 96},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cherryPickup(tt.args.grid); got != tt.want {
				t.Errorf("cherryPickup() = %v, want %v", got, tt.want)
			}
		})
	}
}
