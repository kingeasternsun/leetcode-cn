package main

//158ms 3MB ，可以优化的地方在于 满足条件再左移，再右移动，不必全部比较判断
func minWindow1(s string, t string) string {

	// 重要
	if s == t {
		return t
	}

	count := len(s)
	if count == 0 {
		return ""
	}

	tmap := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		tmap[t[i]] = tmap[t[i]] + 1
	}

	left := 0
	right := 0

	smap := make(map[byte]int, 0)
	smap[s[0]] = 1      //到这里 s肯定不等于t 所以t[0]!=s[0]或者 t[1]还有值
	minLen := count + 1 //这里最大值一定要加一
	maxLeft := 0
	maxRight := -1

	if contain(smap, tmap) {

		tmpLen := right - left + 1
		if tmpLen < minLen {
			minLen = tmpLen
			maxLeft = left
			maxRight = right
		}
	}

	for right < count && left <= right {
		//右移动满足条件，跳出
		right++
		for ; right < count; right++ {
			smap[s[right]] = smap[s[right]] + 1

			if contain(smap, tmap) {

				tmpLen := right - left + 1
				if tmpLen < minLen {
					minLen = tmpLen
					maxLeft = left
					maxRight = right
				}

				break
			}
		}

		//表示right已经到达边界了还找不到满足条件的，可以直接退出了
		if right == count {
			break
		}

		//移除左边 直到不满足
		left++
		for ; left <= right; left++ {

			smap[s[left-1]] = smap[s[left-1]] - 1

			if contain(smap, tmap) {

				tmpLen := right - left + 1
				if tmpLen < minLen {
					minLen = tmpLen
					maxLeft = left
					maxRight = right
				}

			} else {
				break
			}
		}

	}

	return s[maxLeft : maxRight+1]
}

//32ms 3MB
func minWindow(s string, t string) string {

	// 重要
	if s == t {
		return t
	}

	count := len(s)
	if count == 0 {
		return ""
	}

	tmap := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		tmap[t[i]] = tmap[t[i]] + 1
	}

	left := 0
	right := 0
	minLen := count + 1 //这里最大值一定要加一
	maxLeft := 0
	maxRight := -1

	smap := make(map[byte]int, 0)

	var missByte byte //表示因为哪个字节缺一导致没有覆盖
	for right = 0; right < count; right++ {
		smap[s[right]] = smap[s[right]] + 1

		if contain(smap, tmap) {
			missByte = s[right]
			smap[s[right]] = smap[s[right]] - 1
			right-- // 后退-； 因为后面的for虚幻中， right需要先++，所以这里要回退下

			break
		}
	}

	//找不到
	if right == count {
		return ""
	}

	for right < count {
		//右移动直到 满足条件，也就是新增了一个missbyte就肯定满足条件，跳出
		right++
		for ; right < count; right++ {
			smap[s[right]] = smap[s[right]] + 1

			//当前这个right填补了缺失的missbyte，就肯定包含了t
			if s[right] == missByte {
				tmpLen := right - left + 1
				if tmpLen < minLen {
					minLen = tmpLen
					maxLeft = left
					maxRight = right
				}

				break
			}

		}

		//表示right已经到达边界了还找不到满足条件的，可以直接退出了
		if right == count {
			break
		}

		//到这里 还是包含t的， 移除左边 直到不满足； 也就是第一次遇到 smap[left-1] < tmap[left-1]
		left++
		for ; left <= right; left++ {

			smap[s[left-1]] = smap[s[left-1]] - 1

			//left-1这个byte去掉后就不满足了
			if tmap[s[left-1]] > 0 && smap[s[left-1]] < tmap[s[left-1]] {
				missByte = s[left-1]
				break
			}

			tmpLen := right - left + 1
			if tmpLen < minLen {
				minLen = tmpLen
				maxLeft = left
				maxRight = right
			}

		}

	}

	return s[maxLeft : maxRight+1]
}

func contain(smap, tmap map[byte]int) bool {

	for k, v := range tmap {
		if smap[k] < v {
			return false
		}

	}

	return true

}
