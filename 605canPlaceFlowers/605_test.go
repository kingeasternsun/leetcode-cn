package leetcode

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 0, 0, 0, 1}, 1}, true},
		{"2", args{[]int{1, 0, 0, 0, 1}, 2}, false},
		{"3", args{[]int{0, 0, 1, 0, 0}, 2}, true},
		{"4", args{[]int{0, 0, 1, 0, 0}, 2}, true},
		{"5", args{[]int{0, 0, 1, 0, 0}, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
