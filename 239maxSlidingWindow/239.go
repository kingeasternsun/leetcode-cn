package leetcode

//224ms
func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) < k {
		return nil
	}

	//使用单调队列来实现
	monoQueue := make([]int, 0)
	res := make([]int, 0) //保存结果

	for i, v := range nums {

		//从后往前，如果比monoQueue的最后一个要大，就弹出最后一个.
		for len(monoQueue) > 0 {
			if v > monoQueue[len(monoQueue)-1] {
				monoQueue = monoQueue[:len(monoQueue)-1] //弹出
			} else {
				break //注意如果遇到比monoQueue最后一个要小 或 相等 就需要停止弹出了
			}

		}

		//追加v到monoQueue里面
		monoQueue = append(monoQueue, v)

		//窗口长度还没有满k
		if i < k-1 { //注意这里是要小于k-1
			continue
		}

		//刚好满k
		if i == k-1 {
			res = append(res, monoQueue[0])
			continue
		}

		//sliding窗口最左边的要移除
		top := nums[i-k]
		//top就是最大值，那么monoque需要把左边的也移除
		if top == monoQueue[0] {
			monoQueue = monoQueue[1:]
		}

		res = append(res, monoQueue[0])

	}

	return res
}
