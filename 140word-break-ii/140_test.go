package main

import (
	"reflect"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.

		{"1", args{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}}, []string{"cats and dog", "cat sand dog"}},

		{"2", args{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}}, []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}},
		{"3", args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
