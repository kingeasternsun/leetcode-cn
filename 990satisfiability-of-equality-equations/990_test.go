package main

import "testing"

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{[]string{"a==b", "b!=a"}}, false},
		{"2", args{[]string{"b==a", "a==b"}}, true},
		{"3", args{[]string{"a==b", "b==c", "a==c"}}, true},
		{"4", args{[]string{"a==b", "b!=c", "c==a"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
