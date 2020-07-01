package main

func findTheLongestSubstring(s string) int {

	posMap := make(map[uint]int, 0) //记录状态所在最早出现的位置
	posMap[0] = -1                  //表示如果什么字符都不选的话 -1位置，用来和全部元音都是偶数的匹配

	var status uint
	var maxLen int

	for i := range s {

		switch s[i] {
		case 'a':
			status ^= 1
		case 'e':
			status ^= (1 << 1)
		case 'i':
			status ^= (1 << 2)
		case 'o':
			status ^= (1 << 3)
		case 'u':
			status ^= (1 << 4)
		}

		pos, ok := posMap[status]
		if ok { //前面一个位置的各个字母的奇偶数目 和 当前位置的一样
			maxLen = max(maxLen, i-pos)
		} else {
			posMap[status] = i
		}

	}

	return maxLen
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
