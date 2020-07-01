package main

import (
	"testing"
)

func Test_biSearch(t *testing.T) {
	type args struct {
		d []int
		n int
	}
	tests := []struct {
		name  string
		args  args
		wantC int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2}, 2}, 1},
		{"2", args{[]int{1, 2}, 3}, 2},
		{"3", args{[]int{1, 2, 3}, 3}, 2},
		{"4", args{[]int{1}, 3}, 1},
		{"5", args{[]int{3}, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := biSearch(tt.args.d, tt.args.n); gotC != tt.wantC {
				t.Errorf("biSearch() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func Test_sqrtN(t *testing.T) {
	type args struct {
		d int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{2, 3}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sqrtN(tt.args.d, tt.args.n); got != tt.want {
				t.Errorf("sqrtN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_atMostNGivenDigitSet(t *testing.T) {
	type args struct {
		D []string
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]string{"1", "3", "5", "7"}, 100}, 20},
		{"11", args{[]string{"1", "3", "5", "7"}, 101}, 20},
		{"111", args{[]string{"1", "3", "5", "7"}, 7}, 4},
		{"2", args{[]string{"1", "4", "9"}, 1000000000}, 29523},
		{"3", args{[]string{"7"}, 8}, 1},
		{"32", args{[]string{"7"}, 7}, 1},
		{"333", args{[]string{"7"}, 88}, 2},
		{"33", args{[]string{"9"}, 55}, 1},
		{"4", args{[]string{"3", "4", "8"}, 4}, 2},
		{"7", args{[]string{"1", "2", "3", "4", "6", "7", "9"}, 333}, 171},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := atMostNGivenDigitSet(tt.args.D, tt.args.N); got != tt.want {
				t.Errorf("atMostNGivenDigitSet() = %v, want %v", got, tt.want)
			}
		})
	}

}
