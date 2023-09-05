package minflipsmonoincr

import "testing"

func Test_minFlipsMonoIncr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{s: "00110"}, 1},
		{"2", args{s: "010110"}, 2},
		{"3", args{s: "00011000"}, 2},
		{"4", args{s: "00"}, 0},
		{"5", args{s: "11"}, 0},
		{"6", args{s: "01"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlipsMonoIncr(tt.args.s); got != tt.want {
				t.Errorf("minFlipsMonoIncr() = %v, want %v", got, tt.want)
			}
		})
	}
}
