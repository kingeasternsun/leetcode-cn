//  115. 不同的子序列 https://leetcode-cn.com/problems/distinct-subsequences/

package main

import "testing"

// "adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc"
// "bcddceeeebecbc"
func Test_numDistinct(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{s: "rabbbit", t: "rabbit"}, 3},
		{"2", args{s: "babgbag", t: "bag"}, 5},
		{"3", args{s: "adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc",
			t: "bcddceeeebecbc"}, 700531452},
		{"4", args{s: "", t: ""}, 1},
		{"5", args{s: "", t: "bag"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
