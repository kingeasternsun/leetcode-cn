package main

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{3, 4, 5, 2}}, 12},
		{"2", args{[]int{1, 5, 4, 5}}, 16},
		{"3", args{[]int{3, 7}}, 12},
		{"4", args{[]int{8, 8, 7, 4, 2, 8, 1, 7, 7}}, 49},

		//[8,8,7,4,2,8,1,7,7]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
