package leetcode

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	a := &ListNode{
		Val: 2,
	}
	b := &ListNode{
		Val:  5,
		Next: a,
	}

	c := &ListNode{
		Val:  2,
		Next: b,
	}
	d := &ListNode{
		Val:  3,
		Next: c,
	}
	e := &ListNode{
		Val:  4,
		Next: d,
	}
	f := &ListNode{
		Val:  1,
		Next: e,
	}

	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"1", args{f, 3}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
