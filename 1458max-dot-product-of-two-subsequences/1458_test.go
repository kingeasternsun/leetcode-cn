package main

import "testing"

func Test_maxDotProduct(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// // TODO: Add test cases.
		{"1", args{nums1: []int{3, -2}, nums2: []int{2, -6, 7}}, 21},
		{"2", args{nums1: []int{-1, -1}, nums2: []int{1, 1}}, -1},
		{"3", args{nums1: []int{-3, -8, 3, -10, 1, 3, 9}, nums2: []int{9, 2, 3, 7, -9, 1, -8, 5, -1, -1}}, 200},
		// {"3", args{nums1: []int{-3,-8,3,-10,1,3,9}, nums2: []int{9,2,3,7,-9,1,-8,5,-1,-1}, 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDotProduct(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxDotProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
