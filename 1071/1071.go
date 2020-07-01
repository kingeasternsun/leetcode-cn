package main

func main() {

}

/*
 1. 先查找最大前缀
 2. 根据两个字符串的长度能否整除字串的长度 ，提前判断返回
 3. str1[i] != subStr[i%subLen]
*/ 

// 0ms 2.3mb
func gcdOfStrings(str1 string, str2 string) string {

	len1 := len(str1)
	len2 := len(str2)
	minLen := len1
	// maxLen:=len2
	if len2 < minLen {
		minLen = len2
		// maxLen = len1
	}

	i := 0
	for ; i < minLen; i++ {
		if str1[i] != str2[i] {
			break
		}
	}

	if i == 0 {
		return ""
	}

	var result bool
	for j := i; j > 0; j-- {
		result = help((str1[j:]),( str2[j:]), len1-j, len2-j, (str1[:j]), j)
		if result {
			return str1[:j]
		}
	}

	return ""
}

func help(str1 string, str2 string, len1, len2 int, subStr string, subLen int) bool {

	if len1 > 0 {

		//提前判断返回
		if len1/subLen*subLen != len1 {
			return false
		}

		for i := 0; i < len1; i++ {
			if str1[i] != subStr[i%subLen] {
				return false
			}
		}
	}

	if len2 > 0 {
		//提前判断返回
		if len2/subLen*subLen != len2 {

			return false
		}

		for i := 0; i < len2; i++ {
			if str2[i] != subStr[i%subLen] {
				return false
			}
		}
	}

	return true

}



// 4ms 6.4mb
func gcdOfStrings1(str1 string, str2 string) string {

	len1 := len(str1)
	len2 := len(str2)
	minLen := len1
	// maxLen:=len2
	if len2 < minLen {
		minLen = len2
		// maxLen = len1
	}

	i := 0
	for ; i < minLen; i++ {
		if str1[i] != str2[i] {
			break
		}
	}

	if i == 0 {
		return ""
	}

	var result bool
	for j := i; j > 0; j-- {
		result = help1([]byte(str1[j:]),[]byte( str2[j:]), len1-j, len2-j, []byte(str1[:j]), j)
		if result {
			return str1[:j]
		}
	}

	return ""
}

func help1(str1 []byte, str2 []byte, len1, len2 int, subStr []byte, subLen int) bool {

	if len1 > 0 {

		//提前判断返回
		if len1/subLen*subLen != len1 {
			return false
		}

		for i := 0; i < len1; i++ {
			if str1[i] != subStr[i%subLen] {
				return false
			}
		}
	}

	if len2 > 0 {
		//提前判断返回
		if len2/subLen*subLen != len2 {

			return false
		}

		for i := 0; i < len2; i++ {
			if str2[i] != subStr[i%subLen] {
				return false
			}
		}
	}

	return true

}
