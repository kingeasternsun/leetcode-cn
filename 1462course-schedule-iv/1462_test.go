package main

import (
	"reflect"
	"testing"
)

func Test_checkIfPrerequisite(t *testing.T) {
	type args struct {
		n             int
		prerequisites [][]int
		queries       [][]int
	}

	prerequisites1 := [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 3}, []int{3, 4}}
	queries1 := [][]int{[]int{0, 4}, []int{4, 0}, []int{1, 3}, []int{3, 0}}
	result1 := []bool{true, false, true, false}

	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
		{"1", args{5, prerequisites1, queries1}, result1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPrerequisite(tt.args.n, tt.args.prerequisites, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkIfPrerequisite() = %v, want %v", got, tt.want)
			}
		})
	}
}
