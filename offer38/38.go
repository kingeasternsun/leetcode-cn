package leetcode

import (
	"sort"
)

func permutation(s string) []string {
	var res []string
	var pre []byte
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	dfs(b, pre, &res)
	return res
}

//要考虑包含重复字符的情况,只有前一个重复的选中了，当前才能选
func dfs(s []byte, pre []byte, res *[]string) {

	if len(s) == 0 {
		*res = append(*res, string(pre))
		return
	}

	bMap := make(map[byte]struct{}, 0)
	for i, b := range s {

		if _, ok := bMap[b]; ok {
			continue
		}

		bMap[b] = struct{}{}
		s[0], s[i] = s[i], s[0]
		dfs(s[1:], append(pre, b), res)
		s[0], s[i] = s[i], s[0]
	}
	return
}
