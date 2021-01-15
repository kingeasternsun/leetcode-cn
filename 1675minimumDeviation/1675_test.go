package leetcode

import "testing"

func Test_minimumDeviation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 3, 4}}, 1},
		{"2", args{[]int{4, 1, 5, 20, 3}}, 3},
		{"3", args{[]int{2, 10, 8}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeviation(tt.args.nums); got != tt.want {
				t.Errorf("minimumDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
