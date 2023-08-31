package longestnicesubstring

import "testing"

func Test_longestNiceSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Add test cases.
		{name: "1", args: args{"YazaAay"}, want: "aAa"},
		{name: "2", args: args{"Bb"}, want: "Bb"},
		{name: "3", args: args{"C"}, want: ""},
		{name: "4", args: args{"dDzeE"}, want: "dD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestNiceSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestNiceSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
