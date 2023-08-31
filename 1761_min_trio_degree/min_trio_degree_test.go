package mintriodegree

import "testing"

func Test_minTrioDegree(t *testing.T) {
	edges1 := [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}}
	edges2 := [][]int{{1, 3}, {4, 1}, {4, 3}, {2, 5}, {5, 6}, {6, 7}, {7, 5}, {2, 6}}
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
		{name: "1", args: args{n: 6, edges: edges1}, want: 3},
		{name: "2", args: args{n: 7, edges: edges2}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTrioDegree(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("minTrioDegree() = %v, want %v", got, tt.want)
			}
		})
	}
}
