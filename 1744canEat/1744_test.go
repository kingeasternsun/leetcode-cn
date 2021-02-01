package leetcode

import (
	"reflect"
	"testing"
)

func Test_canEat(t *testing.T) {
	candiesCount2 := []int{5, 2, 6, 4, 1}
	queries2 := [][]int{
		[]int{3, 1, 2},
		[]int{4, 10, 3},
		[]int{3, 10, 100},
		[]int{4, 100, 30},
		[]int{1, 3, 1},
	}
	type args struct {
		candiesCount []int
		queries      [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
		{"2", args{candiesCount2, queries2}, []bool{false, true, true, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canEat(tt.args.candiesCount, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canEat() = %v, want %v", got, tt.want)
			}
		})
	}
}

// [true,true,true,true,true,true,true,true,false,false,true,true,false,false,false,true,true,false,false,false,false,false,false,false,false,true,false,false,false,false,,true,false,true,true,false,false,false,true,false,false,false,false,false,true,true,true,false,false,false,false,true,false,false,true,false,true,true,false,true,false,false,true,true,true,true,true,false,false,false,true,true,false,false,true,false,true]
// [true,true,true,true,true,true,true,true,false,false,true,true,false,false,false,true,true,false,false,false,false,false,false,false,false,true,false,false,false,false,,true,false,true,true,false,false,false,true,false,false,false,false,false,true,true,true,false,false,false,false,true,false,false,true,false,true,true,false,true,false,false,true,true,true,true,true,false,false,false,true,true,false,false,true,false,true]
