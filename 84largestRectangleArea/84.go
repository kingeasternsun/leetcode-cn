package leetcode

/*
 利用单调栈记录，以当前柱子为最低值的，左右范围区间
*/

func largestRectangleArea(heights []int) int {

	leftStack := []int{}
	rightStack := []int{}
	leftPos := make([]int, len(heights))
	rightPos := make([]int, len(heights))

	for i, h := range heights {

		for len(leftStack) > 0 {

			tail := leftStack[len(leftStack)-1]
			if heights[tail] >= h {
				leftStack = leftStack[:len(leftStack)-1]
			} else {
				break
			}
		}

		if len(leftStack) == 0 {
			leftPos[i] = -1
		} else {
			leftPos[i] = leftStack[len(leftStack)-1]
		}

		leftStack = append(leftStack, i)

	}

	for j := len(heights) - 1; j >= 0; j-- {
		h := heights[j]
		for len(rightStack) > 0 {

			tail := rightStack[len(rightStack)-1]
			if heights[tail] >= h {
				rightStack = rightStack[:len(rightStack)-1]
			} else {
				break
			}
		}

		if len(rightStack) == 0 {
			rightPos[j] = len(heights)
		} else {
			rightPos[j] = rightStack[len(rightStack)-1]
		}

		rightStack = append(rightStack, j)

	}

	res := 0
	for i, h := range heights {

		res = max(res, (rightPos[i]-leftPos[i]-1)*h)

	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
