package leetcode

import "testing"

func Test_regionsBySlashes(t *testing.T) {

	b1 := []string{
		" /",
		"/ ",
	}
	b2 := []string{

		" /",
		"  ",
	}
	b3 := []string{
		"\\/",
		"/\\",
	}
	b4 := []string{
		"/\\",
		"\\/",
	}
	b5 := []string{
		"//",
		"/ ",
	}
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{b1}, 2},
		{"2", args{b2}, 1},
		{"3", args{b3}, 4},
		{"4", args{b4}, 5},
		{"5", args{b5}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regionsBySlashes(tt.args.grid); got != tt.want {
				t.Errorf("regionsBySlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}
