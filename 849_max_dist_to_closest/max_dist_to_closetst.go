package maxdisttoclosest

func maxDistToClosest(seats []int) int {
	pre := -1
	ret := 0

	for i, v := range seats {
		if v == 0 {
			continue
		}

		if pre < 0 {
			ret = i
		} else if i-pre >= 2 { // 1001 => (3-0)/2=1 , 101 => (2-0)/2=1
			ret = max(ret, (i-pre)/2)
		}
		pre = i
	}

	if pre >= 0 {
		ret = max(ret, len(seats)-1-pre)
	}

	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
