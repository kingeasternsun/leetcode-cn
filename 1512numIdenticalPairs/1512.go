package leetcode

func numIdenticalPairs(nums []int) int {

	numMap := make(map[int]int, 0)
	for _, v := range nums {
		numMap[v]++
	}

	res := 0
	for _, v := range numMap {
		if v == 1 {
			continue
		}
		res += (v) * (v - 1) / 2
	}

	return res

}
