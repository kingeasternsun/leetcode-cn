package main

import "fmt"

func main() {

	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
	fmt.Println(reversePairs([]int{2, 3, 1}))
}

func reversePairs(nums []int) int {

	cnt := len(nums)
	if cnt <= 1 {
		return 0
	}

	newNums := make([]int, cnt)
	copy(newNums, nums)

	leftNum := reversePairs(newNums[:cnt/2])
	rightNum := reversePairs(newNums[cnt/2:])

	//计算后， nums[:cnt/2]  nums[cnt/2:] 也都已经排好了顺序

	//等于 左边数组里面反转对的数量，加上右边数组里面反转对的数量
	// 再加上 i在左半边数据，j在右半边数组组成的反转对的数量
	num := leftNum + rightNum

	//由于两边都已经排序好了，直接利用双指针查找

	j := cnt / 2
	for _, v := range newNums[:cnt/2] {

		//找到不满足 v > 2*newNums[j]的最小的j
		for j < cnt && v > 2*newNums[j] {
			j++
		}

		num += (j - cnt/2)

	}

	//排序
	i := 0
	j = cnt / 2
	k := 0

	//i和j不能同时到达2右边边界，但是有一个到达右边边界是可以的
	for i < cnt/2 || j < cnt {

		if i == cnt/2 {
			nums[k] = newNums[j]
			j++

		} else if j == cnt {
			nums[k] = newNums[i]
			i++
		} else if newNums[i] <= newNums[j] {
			nums[k] = newNums[i]
			i++
		} else {
			nums[k] = newNums[j]
			j++
		}

		k++
	}

	return num
}
