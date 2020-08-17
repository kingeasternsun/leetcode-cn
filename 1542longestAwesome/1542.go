package main

func longestAwesome(s string) int {

	count := len(s)
	if count < 2 {
		return count
	}

	var maxLen = 1
	posMap := make(map[int]int, 0) //记录某种状态最早出现的位置
	var status int = 0
	posMap[0] = -1

	for i := 0; i < len(s); i++ {

		var mask = 1 << (s[i] - '0')
		if status&mask > 0 {
			status = status & (^mask)
		} else {
			status = status | mask
		}

		if prePos, ok := posMap[status]; !ok {
			posMap[status] = i
		} else {

			if i-prePos > maxLen {
				maxLen = i - prePos
			}

		}

		// fmt.Printf("%x\n", status)
		//关键 处理奇数的情况,对每一位进行翻转
		tmpStatus := status
		for j := uint16(0); j < 10; j++ {
			tmpMask := 1 << j
			if tmpStatus&tmpMask > 0 {
				tmpStatus = status & (^tmpMask)
			} else {
				tmpStatus = status | tmpMask
			}
			// fmt.Printf("     %x\n", tmpStatus)

			if prePos, ok := posMap[tmpStatus]; ok {
				if i-prePos > maxLen {
					maxLen = i - prePos
				}
			}

		}

	}

	return maxLen
}
