package main

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}

	intervals1 := [][]int{
		[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},
	}

	newInterval1 := []int{4, 8}

	want1 := [][]int{
		[]int{1, 2},
		[]int{3, 10},
		[]int{12, 16},
	}

	intervals2 := [][]int{
		[]int{1, 5},
	}

	newInterval2 := []int{2, 7}

	want2 := [][]int{
		[]int{1, 7},
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"1", args{intervals1, newInterval1}, want1},
		{"2", args{intervals2, newInterval2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
