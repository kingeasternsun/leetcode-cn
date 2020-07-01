package main

import "testing"

func Test_patternMatching(t *testing.T) {
	type args struct {
		pattern string
		value   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{pattern: "abba", value: "dogcatcatdog"}, true},
		{"2", args{pattern: "abba", value: "dogcatcatfish"}, false},
		{"3", args{pattern: "aaaa", value: "dogcatcatdog"}, false},
		{"4", args{pattern: "abba", value: "dogdogdogdog"}, true},
		{"5", args{pattern: "a", value: ""}, true},
		{"6", args{pattern: "ab", value: ""}, false},                                                                                        //a，b不能相同 所以要注意
		{"7", args{pattern: "baabab", value: "eimmieimmieeimmiee"}, true},                                                                   //a，b不能相同 所以要注意
		{"8", args{pattern: "bbbbbbbbabbbbbbbbbbbabbbbbbba", value: "zezezezezezezezezkxzezezezezezezezezezezezkxzezezezezezezezkx"}, true}, //a，b不能相同 所以要注意

		// "bbbbbbbbabbbbbbbbbbbabbbbbbba"
		// "zezezezezezezezezkxzezezezezezezezezezezezkxzezezezezezezezkx"
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := patternMatching(tt.args.pattern, tt.args.value); got != tt.want {
				t.Errorf("patternMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
