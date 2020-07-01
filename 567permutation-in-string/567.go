package main

//16ms 2.8MB
func checkInclusion1(s1 string, s2 string) bool {
	m1 := make(map[byte]int, 0)
	m2 := make(map[byte]int, 0)

	count1 := len(s1)
	count2 := len(s2)

	if count2 < count1 {
		return false
	}

	for i := range s1 {
		m1[s1[i]]++
	}

	for i := 0; i < count1; i++ {
		m2[s2[i]]++
	}

	if match1(m1, m2) {
		return true
	}

	//向右移动  move to right
	for i := 1; i+count1 <= count2; i++ {
		m2[s2[i-1]]--        // the s2[i-1] remove from window
		m2[s2[i+count1-1]]++ // the s2[i+count1] add to window

		if m2[s2[i-1]] < m1[s2[i-1]] {
			continue
		}

		if match1(m1, m2) {
			return true
		}

	}

	return false
}

func match1(m1, m2 map[byte]int) bool {
	for k, v := range m1 {

		if v > m2[k] {
			return false
		}

	}
	return true
}

//0ms 2.7MB
func checkInclusion(s1 string, s2 string) bool {
	m1 := make([]int, 26)
	m2 := make([]int, 26)

	count1 := len(s1)
	count2 := len(s2)

	if count2 < count1 {
		return false
	}

	for i := range s1 {
		m1[s1[i]-'a']++
	}

	for i := 0; i < count1; i++ {
		m2[s2[i]-'a']++
	}

	if match(m1, m2) {
		return true
	}

	//向右移动  move to right
	for i := 1; i+count1 <= count2; i++ {
		m2[s2[i-1]-'a']--        // the s2[i-1] remove from window
		m2[s2[i+count1-1]-'a']++ // the s2[i+count1] add to window

		if m2[s2[i-1]-'a'] < m1[s2[i-1]-'a'] {
			continue
		}

		if match(m1, m2) {
			return true
		}

	}

	return false
}

func match(m1, m2 []int) bool {
	for k, v := range m1 {

		if v > m2[k] {
			return false
		}

	}
	return true
}
