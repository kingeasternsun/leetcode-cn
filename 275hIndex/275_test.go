package leetcode

import "testing"

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{0, 1, 3, 5, 6}}, 3},
		{"2", args{[]int{}}, 0},
		{"3", args{[]int{0}}, 0},
		{"4", args{[]int{1, 2, 4, 5, 6}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
