package main

import "testing"

/*
houses = [0,0,0,0,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/paint-house-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test_minCost(t *testing.T) {
	type args struct {
		houses []int
		cost   [][]int
		m      int
		n      int
		target int
	}

	houses1 := []int{0, 0, 0, 0, 0}
	cost1 := [][]int{
		[]int{1, 10},
		[]int{10, 1},
		[]int{10, 1},
		[]int{1, 10},
		[]int{5, 1},
	}

	houses2 := []int{0, 2, 1, 2, 0}
	cost2 := [][]int{
		[]int{1, 10},
		[]int{10, 1},
		[]int{10, 1},
		[]int{1, 10},
		[]int{5, 1},
	}

	houses3 := []int{0, 0, 0, 0, 0}
	cost3 := [][]int{
		[]int{1, 10},
		[]int{10, 1},
		[]int{1, 10},
		[]int{10, 1},
		[]int{1, 10},
	}
	houses4 := []int{3, 1, 2, 3}
	cost4 := [][]int{
		[]int{1, 1, 1},
		[]int{1, 1, 1},
		[]int{1, 1, 1},
		[]int{1, 1, 1},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{houses1, cost1, 5, 2, 3}, 9},
		{"2", args{houses2, cost2, 5, 2, 3}, 11},
		{"3", args{houses3, cost3, 5, 2, 5}, 5},
		{"4", args{houses4, cost4, 4, 3, 3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.houses, tt.args.cost, tt.args.m, tt.args.n, tt.args.target); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
