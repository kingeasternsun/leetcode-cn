package leetcode

import "testing"

func Test_isTransformable(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{s: "84532", t: "34852"}, true},
		{"2", args{s: "34521", t: "23415"}, true},
		{"3", args{s: "12345", t: "12435"}, false},
		{"4", args{s: "1", t: "2"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTransformable(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isTransformable() = %v, want %v", got, tt.want)
			}
		})
	}
}
