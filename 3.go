package main

//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
// 	fmt.Println(lengthOfLongestSubstring("bbbbb"))
// 	fmt.Println(lengthOfLongestSubstring("pwwkew"))
// 	fmt.Println(lengthOfLongestSubstring("au"))
// }

// abbd
// 121
// 0123
//
func lengthOfLongestSubstring(s string) int {

	sLen := len(s)
	if sLen == 0 || sLen == 1 {
		return sLen
	}

	// lenArray := make([]int, sLen)
	preLen := 1
	var maxLen = 1
	// lenArray[0] = 1

	for i := 1; i < sLen; i++ {
		// fmt.Println(s[i])
		// lenArray[i]=1

		j := i - preLen
		for ; j < i; j++ {
			if s[j] == s[i] { //如果相同，那么和S[J]到S[I-1]中间的肯定不相同，所以可以直接计算得到
				preLen = i - j
				break
			}
		}

		if j == i { //都不相同
			preLen = preLen + 1
		}

		if preLen > maxLen {
			maxLen = preLen
		}

	}

	return maxLen
}
