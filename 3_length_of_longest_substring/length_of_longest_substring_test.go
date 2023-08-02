package lengthoflongestsubstring

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{s: ""}, want: 0},
		{name: "2", args: args{s: "abcaefg"}, want: 6},
		{name: "3", args: args{s: "abcabcbb"}, want: 3},
		{name: "4", args: args{s: "bbbb"}, want: 1},
		{name: "5", args: args{s: "pwwkew"}, want: 3},
		{name: "6", args: args{s: "a"}, want: 1},
		{name: "7", args: args{s: "aa"}, want: 1},
		{name: "8", args: args{s: "aab"}, want: 2},
		{name: "9", args: args{s: "ab"}, want: 2},
		{name: "10", args: args{s: "abb"}, want: 2},
		{name: "11", args: args{s: "abc"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
