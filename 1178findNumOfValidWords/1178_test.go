/*
 * @Description:1178
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-26 14:56:32
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-26 16:03:04
 * @FilePath: \1178findNumOfValidWords\1178_test.go
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_findNumOfValidWords(t *testing.T) {

	words1 := []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}
	puzzles1 := []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}

	type args struct {
		words   []string
		puzzles []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{words: words1, puzzles: puzzles1}, []int{1, 1, 3, 2, 4, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumOfValidWords(tt.args.words, tt.args.puzzles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumOfValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
