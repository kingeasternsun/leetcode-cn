package main

func firstUniqChar(s string) int {

	charMap := make([][2]int, 26)
	for i := 0; i < 26; i++ {
		charMap[i][0] = -1
		charMap[i][1] = -1
	}

	for i := 0; i < len(s); i++ {
		if charMap[s[i]-'a'][0] == -1 {
			charMap[s[i]-'a'][0] = i
		} else { //已经出现过
			charMap[s[i]-'a'][1] = i
		}

	}

	res := len(s)

	for i := 0; i < 26; i++ {

		//只出现了一次
		if charMap[i][0] >= 0 && charMap[i][1] == -1 {
			if charMap[i][0] < res {
				res = charMap[i][0]
			}
		}

	}

	if res == len(s) {
		return -1
	}
	return res
}
