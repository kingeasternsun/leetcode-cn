package main

/* 42. 接雨水  https://leetcode-cn.com/problems/trapping-rain-water/
从左到右，碰到更低的就会蓄水，碰到更高的柱子，就和左边第二高的构成一个水坑
*/
func trap(height []int) int {

	count := len(height)
	if count <= 2 {
		return 0
	}
	sum := 0
	midSum := 0 //中间柱子占用的空间
	// leftH := height[0]
	leftID := 0
	maxHeigthID := 0 // 关键点 要记录最高的位置
	for i := 1; i < count; i++ {
		if height[i] >= height[leftID] {
			sum = sum + (i-leftID-1)*height[leftID] - midSum
			// leftH = height[i]
			leftID = i
			midSum = 0
			maxHeigthID = i
		} else {
			midSum += height[i]
		}
	}

	//从右向左
	// rightH := height[count-1]
	rightID := count - 1
	midSum = 0
	for i := count - 2; i >= maxHeigthID; i-- {
		if height[i] >= height[rightID] {
			sum = sum + (rightID-i-1)*height[rightID] - midSum
			// rightH = height[i]
			rightID = i
			midSum = 0
		} else {
			midSum += height[i]
		}
	}

	return sum
}
