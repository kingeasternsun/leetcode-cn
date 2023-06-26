package pivotinteger

import "testing"

func Test_pivotInteger(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"8", args{8}, 6},
		{"1", args{1}, 1},
		{"4", args{4}, -1},
		{"3", args{3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotInteger(tt.args.n); got != tt.want {
				t.Errorf("pivotInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
