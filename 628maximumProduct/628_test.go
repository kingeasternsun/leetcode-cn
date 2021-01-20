package leetcode

import "testing"

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 3}}, 6},
		{"2", args{[]int{-1, 2, 3, 4}}, 24},
		{"3", args{[]int{-1, -2, 3, 4}}, 8},
		{"4", args{[]int{-1, -2, -3, 4}}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
