package main

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {

	a := &ListNode{
		Val: 1,
	}
	b := &ListNode{
		Val: 4,
	}
	a.Next = b

	c := &ListNode{
		Val: 1,
	}
	d := &ListNode{
		Val: 3,
	}
	c.Next = d

	var lists []*ListNode
	lists = append(lists, a)
	lists = append(lists, c)

	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"1", args{lists: lists}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
