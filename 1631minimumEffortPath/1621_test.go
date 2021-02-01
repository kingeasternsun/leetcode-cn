package leetcode

import "testing"

func Test_minimumEffortPath(t *testing.T) {
	heights1 := [][]int{[]int{1, 2, 2}, []int{3, 8, 2}, []int{5, 3, 5}}
	heights2 := [][]int{[]int{1, 2, 3}, []int{3, 8, 4}, []int{5, 3, 5}}
	heights3 := [][]int{{1, 2, 1, 1, 1}, []int{1, 2, 1, 2, 1}, []int{1, 2, 1, 2, 1}, []int{1, 2, 1, 2, 1}, []int{1, 1, 1, 2, 1}}
	type args struct {
		heights [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{heights1}, 2},
		{"2", args{heights2}, 1},
		{"3", args{heights3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumEffortPathDFS(tt.args.heights); got != tt.want {
				t.Errorf("minimumEffortPath() = %v, want %v", got, tt.want)
			}
			if got := minimumEffortPathBFS(tt.args.heights); got != tt.want {
				t.Errorf("minimumEffortPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
