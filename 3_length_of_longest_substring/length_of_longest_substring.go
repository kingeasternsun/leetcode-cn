package lengthoflongestsubstring

/*
	在这个实现方法中，通过记录每个字符最近一次在s中的位置，从而快速判断字符是否在子串中存在

┌─┬─┬─┬─┬─┬─┐
│ │i│ │x│ │j│
└─┴─┴─┴─┴─┴─┘
如图，假设双指标分别指向i和j，记录s从i到j-1的子串(不包含j)，s[j]最近一次在s中出现的位置为x，
如果 x>=i, 那么说明 s[j] 在子串s[i:j)中存在
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	ret := 0
	left, right := -1, 0
	prePos := make([]int, 128)
	for i := range prePos {
		prePos[i] = -1
	}

	for right < len(s) {
		b := s[right]
		preID := prePos[b]
		if preID >= left {
			ret = max(ret, right-left)
			left = preID + 1
		}

		prePos[b] = right
		right++
	}

	return max(ret, right-left)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
