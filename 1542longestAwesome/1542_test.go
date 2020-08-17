package main

import "testing"

func Test_longestAwesome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{"3242415"}, 5},
		{"2", args{"12345678"}, 1},
		{"3", args{"213123"}, 6},
		{"4", args{"2"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestAwesome(tt.args.s); got != tt.want {
				t.Errorf("longestAwesome() = %v, want %v", got, tt.want)
			}
		})
	}
}
