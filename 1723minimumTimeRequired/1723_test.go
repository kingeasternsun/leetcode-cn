package leetcode

import "testing"

func Test_minimumTimeRequired(t *testing.T) {
	out := []int{10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000}
	out2 := []int{643, 526, 589, 976, 986, 730, 345, 926, 798, 618, 827, 873}
	// 11
	type args struct {
		jobs []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{jobs: []int{1, 2, 4, 7, 8}, k: 2}, 11},
		{"2", args{jobs: []int{1, 2, 4, 7}, k: 2}, 7},
		{"3", args{jobs: []int{1, 7, 2, 4}, k: 2}, 7},
		{"4", args{jobs: out, k: 2}, 60000000},
		{"5", args{jobs: out2, k: 11}, 986},
		//[]

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTimeRequired(tt.args.jobs, tt.args.k); got != tt.want {
				t.Errorf("minimumTimeRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
