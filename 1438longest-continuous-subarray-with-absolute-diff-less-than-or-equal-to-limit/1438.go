package main

//https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
func longestSubarray(nums []int, limit int) int {

	if len(nums) == 0 {
		return 0
	}

	list := []int{}    //从后面追加，从前面删除
	maxMono := []int{} //作为一个严格递减序列，最大值就是第一个
	minMono := []int{} //最小值递增队列
	maxLen := 0
	// listLen := 1

	for _, v := range nums {

		// if i == 0 {
		// 	continue
		// }

		// 当前值加入
		list = append(list, v)
		// listLen++

		//更新 maxMono 作为一个严格递减序列
		j := len(maxMono) - 1
		for j >= 0 && maxMono[j] < v { ////碰到比v大或相等的就停止
			j--
		}

		//走到这里 要么j变为了-1，也就是maxMono里面的都比v小，那么都可以清除
		//要么就是 maxMono[j] >= v,
		maxMono = maxMono[:j+1] //注意先截断 再append
		maxMono = append(maxMono, v)

		//更新 minMono 作为一个严格递增序列
		j = len(minMono) - 1
		for j >= 0 && minMono[j] > v { ////碰到比v小或相等的就停止
			j--
		}
		//走到这里 要么j变为了-1，也就是 minMono 里面的都比v大，那么都可以清除
		//要么就是 minMono[j] <= v,
		minMono = minMono[:j+1]
		minMono = append(minMono, v)

		for maxMono[0]-minMono[0] > limit {
			//需要把list最前面的剔除掉
			if list[0] == maxMono[0] {
				maxMono = maxMono[1:]
			}
			if list[0] == minMono[0] {
				minMono = minMono[1:]
			}

			list = list[1:]
			// listLen--

		}

		//最大值和最小值的差值 小于 limit
		if len(list) > maxLen {
			maxLen = len(list)
		}

	}

	return maxLen
}
