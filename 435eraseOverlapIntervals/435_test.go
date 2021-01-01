package leetcode

import "testing"

func Test_eraseOverlapIntervals(t *testing.T) {

	in1 := [][]int{
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 4},
		[]int{1, 3},
	}
	in2 := [][]int{
		[]int{1, 2},
		[]int{1, 2},
		[]int{1, 2},
	}
	in3 := [][]int{
		[]int{1, 2},
		[]int{2, 3},
	}

	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{in1}, 1},
		{"2", args{in2}, 2},
		{"3", args{in3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
