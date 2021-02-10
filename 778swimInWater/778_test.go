package leetcode

import "testing"

func Test_swimInWater(t *testing.T) {
	g1 := [][]int{
		[]int{0, 2},
		[]int{1, 3},
	}

	g3 := [][]int{
		[]int{3, 2},
		[]int{0, 1},
	}
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// 　{"1",args{g1},3},
		{"1", args{g1}, 3},
		{"3", args{g3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swimInWater(tt.args.grid); got != tt.want {
				t.Errorf("swimInWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
