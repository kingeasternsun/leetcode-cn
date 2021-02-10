/*
 * @Description:567
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-10 13:08:55
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-10 13:58:59
 * @FilePath: /567checkInclusion/567.go
 */
package leetcode

/*
转为在s2上执行固定长度为len(s1)的滑动窗口
*/
func checkInclusion1(s1 string, s2 string) bool {
	winLen := len(s1)
	sLen2 := len(s2)

	if winLen > sLen2 {
		return false
	}
	m1 := make([]int, 26)
	m2 := make([]int, 26)

	for i := 0; i < winLen; i++ {
		m1[s1[i]-'a']++
		m2[s2[i]-'a']++
	}

	left := 0
	right := winLen - 1
	for right < sLen2 {

		if match(m1, m2) {
			return true
		}

		if right == sLen2-1 {
			return false
		}

		m2[s2[left]-'a']--
		left++
		right++
		m2[s2[right]-'a']++

	}
	return false

}

func match(m1, m2 []int) bool {
	for i := 0; i < 26; i++ {
		if m2[i] != m1[i] {
			return false
		}
	}

	return true
}

/*
转为在s2上执行固定长度为len(s1)的滑动窗口
*/
func checkInclusion(s1 string, s2 string) bool {
	winLen := len(s1)
	sLen2 := len(s2)

	if winLen > sLen2 {
		return false
	}
	m1 := [26]int{}
	m2 := [26]int{}

	for i := 0; i < winLen; i++ {
		m1[s1[i]-'a']++
		m2[s2[i]-'a']++
	}

	left := 0
	right := winLen - 1
	for right < sLen2 {

		if m1 == m2 {
			return true
		}

		if right == sLen2-1 {
			return false
		}

		m2[s2[left]-'a']--
		left++
		right++
		m2[s2[right]-'a']++

	}
	return false

}
