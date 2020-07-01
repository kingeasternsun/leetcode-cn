package main

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 5}, 11}, 3},
		{"2", args{[]int{2}, 3}, -1},                    //[186,419,83,408]
		{"3", args{[]int{186, 419, 83, 408}, 6249}, 20}, //[]
		{"4", args{[]int{1, 7, 10}, 14}, 2},             //[] 最先找到的不一定最优化
		// 62/49
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
