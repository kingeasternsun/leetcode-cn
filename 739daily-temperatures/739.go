package main

//80ms 	6.9MB
/*
1. 利用map记录每个温度都在哪些位置
2. 对于位置i上的温度t，map中查询温度大于t的位置里面的第一个位置，然后求最小值
*/
func dailyTemperatures(T []int) []int {

	count := len(T)
	tMap := make([][]int, 71) //记录某个温度的在第几天
	for i, t := range T {
		pos := tMap[t-30]
		pos = append(pos, i)
		tMap[t-30] = pos
	}

	result := make([]int, len(T))
	for i := 0; i < count; i++ {

		j := T[i] + 1
		tmpLen := count + 1
		for ; j <= 100; j++ {
			pos := tMap[j-30]
			if len(pos) == 0 {
				continue
			}

			//注意 这里不能够直接返回，要继续查找是否有别的位置的温度
			if pos[0]-i < tmpLen {
				tmpLen = pos[0] - i
			}
			// result[i] = pos[0] - i
			// break

		}

		if tmpLen == count+1 {
			result[i] = 0
		} else {
			result[i] = tmpLen
		}

		//从tMap中温度 T[i]-30 删除该天 i
		pos := tMap[T[i]-30]
		if len(pos) > 0 {
			pos = pos[1:]
			tMap[T[i]-30] = pos
		}
	}

	return result

}

//224ms 7.2MB
func dailyTemperatures1(T []int) []int {

	count := len(T)
	// tMap := make([][]int, 71) //记录某个温度的在第几天
	tMap := make(map[int][]int) //记录某个温度的在第几天
	for i, t := range T {
		pos := tMap[t]
		pos = append(pos, i)
		tMap[t] = pos
	}

	result := make([]int, len(T))
	for i := 0; i < count; i++ {

		key := T[i]
		tmpLen := count + 1
		if key != 100 {
			// j := T[i] + 1
			for j := range tMap {

				if j <= key {
					continue
				}

				pos, ok := tMap[j]
				if !ok {
					continue
				}

				//注意 这里不能够直接返回，要继续查找是否有别的位置的温度
				if pos[0]-i < tmpLen {
					tmpLen = pos[0] - i
				}
				// result[i] = pos[0] - i
				// break

			}
		}

		if tmpLen == count+1 {
			result[i] = 0
		} else {
			result[i] = tmpLen
		}

		//从tMap中温度 T[i]-30 删除该天 i
		pos := tMap[key]
		if len(pos) > 0 {
			pos = pos[1:]

			if len(pos) == 0 {
				delete(tMap, key)
			} else {
				tMap[key] = pos

			}

		}
	}

	return result

}
