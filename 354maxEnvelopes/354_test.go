package main

import "testing"

func Test_maxEnvelopes(t *testing.T) {

	a := [][]int{
		[]int{5, 4},
		[]int{6, 4},
		[]int{6, 7},
		[]int{2, 3},
	}

	b := [][]int{
		[]int{4, 5},
		[]int{4, 6},
		[]int{6, 7},
		[]int{2, 3},
		[]int{1, 1},
	}
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{a}, 3},
		{"2", args{b}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
