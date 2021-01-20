package leetcode

//递归实现
//根据题目种的提示，每次判断s中是否有字符可以通过向前冒泡移动的方式，到达s的开头和t的开头字母相同
func isTransformable_timeout(s string, t string) bool {

	return help([]byte(s), []byte(t))
}

//timeout
func help(s []byte, t []byte) bool {

	if len(s) == 0 && len(t) == 0 {
		return true
	}

	if s[0] == t[0] {
		return help(s[1:], t[1:])
	}

	target := t[0]

	for i := 0; i < len(s); i++ {

		//当前字符比 target 要小，那么s中后面即使有和target相同的字符也无法移动过来
		if s[i] < target {
			return false
		}

		//找到相同的,将 s[i]之前的往后移动
		if s[i] == target {

			for j := i; j > 0; j-- {
				s[j] = s[j-1]
			}
			s[0] = target
			return help(s[1:], t[1:])

		}
	}

	//找不到相同的
	return false

}

func isTransformable(s string, t string) bool {

	posMap := make([][]int, 10)

	for i := 0; i < len(s); i++ {
		posMap[int(s[i]-'0')] = append(posMap[int(s[i]-'0')], i) //记录s中每个字符出现的位置
	}

	//对于t中的每个字符 判断
	for i := 0; i < len(t); i++ {
		num := int(t[i] - '0')
		pos := posMap[num]
		if len(pos) == 0 { //在s中没有对应的字符 肯定无法完成
			return false
		}

		// t[i]在s中的第一个出现位置 pos[0]
		// 判断s中，是否有比num小的出现在pos[0]之前，挡住了移动的道路
		for j := 0; j < num; j++ {

			//如果数字j在s中的第一个位置在 pos[0]前面，挡住了移动的道路
			if len(posMap[j]) > 0 && posMap[j][0] < pos[0] {
				return false
			}
			//如果数字j在s中的第一个位置在 pos[0]后面，由于posMap是单调递增的，所以数字j在s中的位置肯定都是大于pos[0]的

		}

		//s中，比num小的数字都在pos[0]后面
		posMap[num] = pos[1:]
	}

	return true

}
