package captureforts

// 0ms 1.9MB
func captureForts(forts []int) int {
	max := 0
	pre := -1
	for i := 0; i < len(forts)-1; i++ {
		if forts[i] != 0 && forts[i+1] == 0 {
			pre = i
		} else if forts[i] == 0 && forts[i+1] != 0 {
			if pre >= 0 && forts[pre]+forts[i+1] == 0 {
				if i-pre > max {
					max = i - pre
				}
			}

			pre = -1
		}
	}

	return max
}
