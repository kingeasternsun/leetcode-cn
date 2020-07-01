package main

import "testing"

func Test_maxSatisfaction(t *testing.T) {
	type args struct {
		satisfaction []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{-1, -8, 0, 5, -9}}, 14},
		{"2", args{[]int{4, 3, 2}}, 20},
		{"3", args{[]int{-1, -4, -5}}, 0},
		{"4", args{[]int{-2, 5, -1, 0, 3, -3}}, 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfaction(tt.args.satisfaction); got != tt.want {
				t.Errorf("maxSatisfaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
