package isinterleave

type Key struct {
	m, n int
}

// 0ms 2.85mb
func isInterleave(s1 string, s2 string, s3 string) bool {
	cache := make(map[Key]bool, 0)

	return df([]byte(s1), []byte(s2), []byte(s3), cache)

}

func df(s1 []byte, s2 []byte, s3 []byte, cache map[Key]bool) bool {

	if len(s1)+len(s2) != len(s3) {
		return false
	}

	if len(s3) == 0 {
		return true
	}

	if ret, ok := cache[Key{m: len(s1), n: len(s2)}]; ok {
		return ret
	}

	ret1 := false
	ret2 := false

	if len(s1) > 0 && s1[0] == s3[0] {
		ret1 = df(s1[1:], s2, s3[1:], cache)
	}

	if ret1 {
		cache[Key{m: len(s1), n: len(s2)}] = true
		return true
	}

	if len(s2) > 0 && s2[0] == s3[0] {
		ret2 = df(s1, s2[1:], s3[1:], cache)
	}

	cache[Key{m: len(s1), n: len(s2)}] = ret2
	return ret2

}
