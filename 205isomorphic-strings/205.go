package main

//就是map的用法
func isIsomorphic(s string, t string) bool {

	sLen := len(s)
	tLen := len(t)
	if sLen != tLen {
		return false
	}

	transMap := make(map[byte]byte, 0)
	for i := 0; i < sLen; i++ {
		transByte, ok := transMap[s[i]]
		if !ok {
			transMap[s[i]] = t[i]
			continue
		}

		if transByte != t[i] {
			return false
		}

	}

	//反向
	transMap = make(map[byte]byte, 0)
	for i := 0; i < tLen; i++ {
		transByte, ok := transMap[t[i]]
		if !ok {
			transMap[t[i]] = s[i]
			continue
		}

		if transByte != s[i] {
			return false
		}

	}

	return true

}
