package main

func main() {

}

/*

nums:	1 3 5 2 4
maxQ:	5 4
minQ:	1 2 4
*/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	count := len(nums)
	if count <= 1 {
		return false
	}

	maxQ := []int{nums[0]} //维护左边的第一项是最大值，里面的队列单调递减
	minQ := []int{nums[0]} //维护左边的第一项是最小值，里面的队列单调递减
	for i := 1; i < k; i++ {

	}

}
