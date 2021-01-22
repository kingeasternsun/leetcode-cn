package leetcode

import "testing"

func Test_largestSubmatrix(t *testing.T) {

	matrix1 := [][]int{[]int{0, 0, 1}, []int{1, 1, 1}, []int{1, 0, 1}}
	matrix2 := [][]int{[]int{1, 0, 1, 0, 1}}
	matrix3 := [][]int{[]int{1, 1, 0}, []int{1, 0, 1}}
	matrix4 := [][]int{[]int{0, 0}, []int{0, 0}}
	matrix5 := [][]int{[]int{1}}
	matrix6 := [][]int{[]int{0}}

	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{matrix1}, 4},
		{"2", args{matrix2}, 3},
		{"3", args{matrix3}, 2},
		{"4", args{matrix4}, 0},
		{"5", args{matrix5}, 1},
		{"6", args{matrix6}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSubmatrix(tt.args.matrix); got != tt.want {
				t.Errorf("largestSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
