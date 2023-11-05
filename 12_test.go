package main

import "testing"

func Test_minGroupsForValidAssignment(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{10, 10, 10, 3, 1, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minGroupsForValidAssignment(tt.args.nums); got != tt.want {
				t.Errorf("minGroupsForValidAssignment() = %v, want %v", got, tt.want)
			}
		})
	}
}
