package main

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}

	want1 := [][]string{
		[]string{"a", "a", "b"},
		[]string{"aa", "b"},
	}
	want2 := [][]string{
		[]string{"a", "b", "a"},
		[]string{"aba"},
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{"1", args{"aab"}, want1},
		{"2", args{"aba"}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
