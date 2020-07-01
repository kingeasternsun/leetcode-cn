package main

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}

	intervals1 := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}

	intervals2 := [][]int{
		[]int{1, 4},
		[]int{2, 3},
	}

	want1 := [][]int{
		[]int{1, 6},
		[]int{8, 10},
		[]int{15, 18},
	}

	want2 := [][]int{
		[]int{1, 4},
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"1", args{intervals1}, want1},
		{"2", args{intervals2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
