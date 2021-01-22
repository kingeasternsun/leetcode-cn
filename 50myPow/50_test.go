package leetcode

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"1", args{x: 2.0000, n: 10}, 1024.0000},
		{"2", args{x: 2.10000, n: 3}, 9.26100},
		{"3", args{x: 2.00000, n: -2}, 0.25000},
		{"4", args{x: 2.00000, n: -3}, 0.12500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
