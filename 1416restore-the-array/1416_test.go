package main

import "testing"

func Test_numberOfArrays(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{"1000", 10000}, 1},
		{"2", args{"1000", 10}, 0},
		{"3", args{"1317", 2000}, 8},
		{"4", args{"2020", 30}, 1},
		{"5", args{"1234567890", 90}, 34},
		{"6", args{"600342244431311113256628376226052681059918526204", 703}, 411743991},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArrays(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("numberOfArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
