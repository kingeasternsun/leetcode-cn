package maxdisttoclosest

import "testing"

func Test_maxDistToClosest(t *testing.T) {
	type args struct {
		seats []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{seats: []int{1, 0, 0, 0, 1, 0, 1}}, want: 2},
		{name: "2", args: args{seats: []int{1, 0, 0, 0}}, want: 3},
		{name: "3", args: args{seats: []int{0, 1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistToClosest(tt.args.seats); got != tt.want {
				t.Errorf("maxDistToClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
