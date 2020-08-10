package main

import (
	"reflect"
	"testing"
)

func Test_maxNumOfSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"1", args{"adefaddaccc"}, []string{"e", "f", "ccc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumOfSubstrings(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxNumOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
