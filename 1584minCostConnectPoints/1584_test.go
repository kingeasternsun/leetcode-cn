package leetcode

import "testing"

func Test_minCostConnectPoints(t *testing.T) {

	points1 := [][]int{
		[]int{0, 0},
		[]int{2, 2},
		[]int{3, 10},
		[]int{5, 2},
		[]int{7, 0},
	}
	points2 := [][]int{
		[]int{3, 12},
		[]int{-2, 5},
		[]int{-4, 1},
	}
	points3 := [][]int{
		[]int{0, 0},
		[]int{1, 1},
		[]int{1, 0},
		[]int{-1, 1},
	}
	points4 := [][]int{
		[]int{-1000000, -1000000},
		[]int{1000000, 1000000},
	}
	points5 := [][]int{
		[]int{0, 0},
	}
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{points1}, 20},
		{"2", args{points2}, 18},
		{"3", args{points3}, 4},
		{"4", args{points4}, 4000000},
		{"5", args{points5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostConnectPoints(tt.args.points); got != tt.want {
				t.Errorf("minCostConnectPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
