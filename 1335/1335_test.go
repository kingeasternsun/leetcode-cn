package main

import "testing"

func Test_minDifficulty(t *testing.T) {
	type args struct {
		job []int
		d   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "", args: args{job: []int{6, 5, 4, 3, 2, 1}, d: 2}, want: 7},
		{name: "", args: args{job: []int{9, 9, 9}, d: 4}, want: -1},
		{name: "", args: args{job: []int{1, 1, 1}, d: 3}, want: 3},
		{name: "", args: args{job: []int{7, 1, 7, 1, 7, 1}, d: 3}, want: 15},
		{name: "", args: args{job: []int{11, 111, 22, 222, 33, 333, 44, 444}, d: 6}, want: 843},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDifficulty(tt.args.job, tt.args.d); got != tt.want {
				t.Errorf("minDifficulty() = %v, want %v", got, tt.want)
			}
		})
	}
}
