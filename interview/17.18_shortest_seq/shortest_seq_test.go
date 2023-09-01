package shortestseq

import (
	"reflect"
	"testing"
)

func Test_shortestSeq(t *testing.T) {
	type args struct {
		big   []int
		small []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// Add test cases.
		{"1", args{big: []int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7}, small: []int{1, 5, 9}}, []int{7, 10}},
		{"2", args{big: []int{1, 2, 3}, small: []int{4}}, nil},
		{"3", args{big: []int{1, 1, 1}, small: []int{4}}, nil},
		{"4", args{big: []int{1, 1, 1}, small: []int{1}}, []int{0, 0}},
		{"5", args{big: []int{1}, small: []int{1}}, []int{0, 0}},
		{"6", args{big: []int{1}, small: []int{4}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestSeq(tt.args.big, tt.args.small); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}
