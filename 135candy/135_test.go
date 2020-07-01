package main

import "testing"

func Test_candy(t *testing.T) {

	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 0, 2}}, 5},
		{"2", args{[]int{1, 2, 2}}, 4},
		{"3", args{[]int{1, 6, 10, 8, 7, 3, 2}}, 18},
		{"4", args{[]int{1, 1, 1, 1, 1}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
