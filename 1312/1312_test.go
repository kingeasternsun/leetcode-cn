package main

import "testing"

func Test_minInsertions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{"zzazz"}, 0},
		{"2", args{"mbadm"}, 2},
		{"3", args{"leetcode"}, 5},
		{"4", args{"g"}, 0},
		{"5", args{"no"}, 1},
		{"6", args{"tldjbqjdogipebqsohdypcxjqkrqltpgviqtqz"}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.args.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}
		})
	}
}
