package minflipsmonoincr

// 只需要对比当前状态和期望状态的差值
// 0 0 0 1 1 0 0 0
// left   | right
// 0 0 0 0 1 1 1 1
func minFlipsMonoIncr(s string) int {

	left := []int{0, 0}
	right := []int{0, 0}
	for _, b := range []byte(s) {
		if b == '0' {
			left[0]++
		} else {
			left[1]++
		}
	}

	// 变为都是 0 所需要的操作数
	min := left[1]
	// 变为都是 1 所需要的操作数
	if left[0] < left[1] {
		min = left[0]
	}

	// 期望状态: 从 0 到 i-1 都是0, 从 i 到 len(s)-1 都是 1
	for i := len(s) - 1; i >= 0; i-- {
		b := s[i]
		if b == '0' {
			left[0]--
			right[0]++
		} else {
			left[1]--
			right[1]++
		}

		// 左边的1变为0， 右边的0变为1
		if left[1]+right[0] < min {
			min = left[1] + right[0]
		}
	}

	return min

}
