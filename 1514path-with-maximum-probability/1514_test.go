package main

import "testing"

func Test_maxProbability(t *testing.T) {

	type args struct {
		n        int
		edges    [][]int
		succProb []float64
		start    int
		end      int
	}
	edges1 := [][]int{
		[]int{0, 1},
		[]int{1, 2},
		[]int{0, 2},
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"1", args{n: 3, edges: edges1, succProb: []float64{0.5, 0.5, 0.2}, start: 0, end: 2}, 0.25000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProbability(tt.args.n, tt.args.edges, tt.args.succProb, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("maxProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
