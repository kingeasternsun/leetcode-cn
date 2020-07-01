// 1024. 视频拼接	https://leetcode-cn.com/problems/video-stitching/

package main

import "testing"

func Test_videoStitching(t *testing.T) {

	clips := make([][]int, 16)
	clips[0] = []int{0, 1}
	clips[1] = []int{6, 8}
	clips[2] = []int{0, 2}
	clips[3] = []int{5, 6}
	clips[4] = []int{0, 4}
	clips[5] = []int{0, 3}
	clips[6] = []int{6, 7}
	clips[7] = []int{1, 3}
	clips[8] = []int{4, 7}
	clips[9] = []int{1, 4}
	clips[10] = []int{2, 5}
	clips[11] = []int{2, 6}
	clips[12] = []int{3, 4}
	clips[13] = []int{4, 5}
	clips[14] = []int{5, 7}
	clips[15] = []int{6, 9}

	/*
		[[5,7],[1,8],[0,0],[2,3],[4,5],[0,6],[5,10],[7,10]]
		5
	*/
	clips1 := make([][]int, 8)
	clips1[0] = []int{5, 7}
	clips1[1] = []int{1, 8}
	clips1[2] = []int{0, 0}
	clips1[3] = []int{2, 3}
	clips1[4] = []int{4, 5}
	clips1[5] = []int{0, 6}
	clips1[6] = []int{5, 10}
	clips1[7] = []int{7, 10}

	type args struct {
		clips [][]int
		T     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{clips, 9}, 3},
		{"2", args{clips1, 5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := videoStitching(tt.args.clips, tt.args.T); got != tt.want {
				t.Errorf("videoStitching() = %v, want %v", got, tt.want)
			}
		})
	}
}
