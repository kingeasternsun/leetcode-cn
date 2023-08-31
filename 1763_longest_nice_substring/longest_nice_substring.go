package longestnicesubstring

// 1763. 最长的美好子字符串 8ms 2.13mb
func longestNiceSubstring(s string) string {
	left := 0
	right := 0
	max := 0
	for i := 0; i < len(s); i++ {

		niceMap := make(map[byte]struct{})   //记录大小写已经出现的字符
		notNicMap := make(map[byte]struct{}) //记录不存在同时大小写的字符

		for j := i; j < len(s); j++ {

			// 如果已经是美丽字符了，就肯定还是美丽字符，所以这里只考虑不是美丽字符的情况
			if _, ok := niceMap[s[j]]; !ok {
				if _, ok := notNicMap[switchCase2(s[j])]; ok {
					// 如果另外一种形式存在，就从 不美丽 中移除，放入到美丽字符中
					delete(notNicMap, switchCase2(s[j]))
					niceMap[switchCase2(s[j])] = struct{}{}
					niceMap[s[j]] = struct{}{}
				} else {
					// 如果另外一种形式不存在，就还是 不美丽 字符
					notNicMap[s[j]] = struct{}{}
				}
			}

			// 都是美丽字符
			if len(notNicMap) == 0 && j-i+1 > max {
				max = j + 1 - i
				left = i
				right = j + 1
			}

		}
	}
	return s[left:right]
}

// 大小写转换函数
func switchCase2(ch byte) byte {
	if ch >= 'a' && ch <= 'z' {
		return ch - 'a' + 'A' // 将小写字母转为大写
	} else if ch >= 'A' && ch <= 'Z' {
		return ch - 'A' + 'a' // 将大写字母转为小写
	}
	return ch // 非字母字符不做处理
}
