package main

func maxSlidingWindow(nums []int, k int) []int {

	if k == 0 {
		return nil
	}

	count := len(nums)
	if count == 0 {
		return nil
	}
	if k >= count {
		maxV := nums[0]
		for _, v := range nums {
			if v > maxV {
				maxV = v
			}
		}

		return []int{maxV}
	}

	monoMax := []int{nums[0]} //单调递减数量，因为从左边退出，最左边是最大值
	result := make([]int, 0)
	for i := 1; i < k; i++ {
		j := len(monoMax) - 1
		for ; j >= 0; j-- {
			if monoMax[j] >= nums[i] {
				break
			}
		}
		//下面的部分必须放到外面，考虑j为-1的情况
		monoMax = monoMax[:j+1]
		monoMax = append(monoMax, nums[i])

	}

	result = append(result, monoMax[0])
	//窗口向右移动
	for right := k; right < count; right++ {
		//窗口 添加 right
		j := len(monoMax) - 1
		for ; j >= 0; j-- {
			//找到大于等于 nums[right] 的，追加到后面
			if monoMax[j] >= nums[right] {
				break
			}
		}

		//下面的部分必须放到外面，考虑j为-1的情况
		monoMax = monoMax[:j+1]
		monoMax = append(monoMax, nums[right])

		//窗口移除 left
		left := right - k
		if nums[left] == monoMax[0] {
			monoMax = monoMax[1:]
		}
		result = append(result, monoMax[0])
	}

	return result

}
