package leetcode

import "testing"

func Test_maxNumEdgesToRemove(t *testing.T) {

	edges1 := [][]int{[]int{3, 1, 2}, []int{3, 2, 3}, []int{1, 1, 3}, []int{1, 2, 4}, []int{1, 1, 2}, []int{2, 3, 4}}
	edges2 := [][]int{[]int{3, 1, 2}, []int{3, 2, 3}, []int{1, 1, 4}, []int{2, 1, 4}}
	edges3 := [][]int{[]int{3, 2, 3}, []int{1, 1, 2}, []int{2, 3, 4}}
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{4, edges1}, 2},
		{"2", args{4, edges2}, 0},
		{"3", args{4, edges3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumEdgesToRemove(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("maxNumEdgesToRemove() = %v, want %v", got, tt.want)
			}
		})
	}
}
