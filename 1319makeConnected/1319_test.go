package leetcode

import "testing"

func Test_makeConnected(t *testing.T) {
	con1 := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{1, 2},
	}
	con2 := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{0, 3},
		[]int{1, 2},
		[]int{1, 3},
	}

	con3 := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{0, 3},
		[]int{1, 2},
	}
	con4 := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{3, 4},
		[]int{2, 3},
	}
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{4, con1}, 1},
		{"2", args{6, con2}, 2},
		{"3", args{6, con3}, -1},
		{"4", args{5, con4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeConnected(tt.args.n, tt.args.connections); got != tt.want {
				t.Errorf("makeConnected() = %v, want %v", got, tt.want)
			}
		})
	}
}
