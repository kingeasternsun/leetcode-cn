package main

import "testing"

func Test_subarraysDivByK(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {"1", args{[]int{4, 5, 0, -2, -3, 1}, 5}, 7},
		// {"2", args{[]int{-1, 2, 9}, 2}, 2},
		// {"3", args{[]int{-3, 2, 9}, 3}, 2},
		{"4", args{[]int{8, 9, 7, 8, 9}, 8}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysDivByK(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("subarraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
