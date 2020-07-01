package main

//https://leetcode-cn.com/problems/minimum-size-subarray-sum/
//4MS 3.9MB
func minSubArrayLen(s int, nums []int) int {
	count := len(nums)
	if count == 0 {
		return 0
	}
	minLen := count + 1

	tmpSum := 0
	right := 0
	for ; right < count; right++ {

		if nums[right] >= s {
			return 1
		}

		tmpSum += nums[right]
		if tmpSum >= s {
			break
		}

	}

	// sum of all the nums less value of s,  because all the value in nums is > 0,  so just return
	if tmpSum < s {
		return 0
	}

	// backforward one step
	minLen = right + 1

	// tmpSum -= nums[right]
	// right--

	left := 0
	for right < count {

		//move right until tmpSum >= t
		for {

			if tmpSum >= s {
				break
			}

			right++
			if right < count {
				tmpSum += nums[right]
			} else {
				break
			}

		}

		// has move the end means tmpSum < t, because all the value in nums is > 0; so we can exit now
		if right == count {
			break
		}

		if right-left+1 < minLen {
			minLen = right - left + 1
		}

		//move left until tmpSum < s, and left can equal to right
		for left <= right {

			if tmpSum < s {
				break
			}

			tmpSum = tmpSum - nums[left]
			left++

		}

		// nums[right]>t
		if left > right {
			return 1
		}

		// sum[left-1,right] 大于 s
		if (right - left + 2) < minLen {
			minLen = right - left + 2
		}

	}

	return minLen
}
