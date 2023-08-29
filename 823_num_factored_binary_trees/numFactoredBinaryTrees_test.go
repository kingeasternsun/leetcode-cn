package numfactoredbinarytrees

import "testing"

func Test_numFactoredBinaryTrees(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{arr: []int{2, 4}}, 3},
		{"2", args{arr: []int{2, 4, 5, 10}}, 7},
		{"3", args{arr: []int{2, 4, 8}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numFactoredBinaryTrees(tt.args.arr); got != tt.want {
				t.Errorf("numFactoredBinaryTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
