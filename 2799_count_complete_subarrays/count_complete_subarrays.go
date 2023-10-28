package countcompletesubarrays

// 12mb 0.65mb
func countCompleteSubarrays(nums []int) int {

	set := make(map[int]struct{}, 0)
	for _, n := range nums {
		set[n] = struct{}{}
	}
	cnt := len(set)

	ans := 0
	subMap := make(map[int]int, 0)
	left := 0
	for right, v := range nums {
		subMap[v]++
		for len(subMap) == cnt {
			ans += len(nums) - right // 以left为左边界的子数组有多少个
			if subMap[nums[left]] == 1 {
				delete(subMap, nums[left])
			} else {
				subMap[nums[left]]--
			}
			left++
		}

	}
	return ans

}
