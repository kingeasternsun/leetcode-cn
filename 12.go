package main

//https://leetcode-cn.com/problems/integer-to-roman/

// func intToRoman(num int) string {

// 	for{
// 		n:=num/
// 	}
// }

func minGroupsForValidAssignment(nums []int) int {

	cnt := make(map[int]int, 0)
	for _, n := range nums {
		cnt[n]++
	}

	mivCnt := len(nums)
	for _, v := range cnt {
		if v < mivCnt {
			mivCnt = v
		}
	}

	check := func(k int) (int, bool) {
		ans := 0
		for _, v := range cnt {
			m := v / k
			n := v % k
			if n > m {
				return 0, false
			}
			if v%(k+1) == 0 {
				ans += v / (k + 1)
			} else {
				ans += m
			}
		}

		return ans, true
	}

	for k := mivCnt; k >= 1; k-- {
		ans, ok := check(k)
		if ok {
			return ans
		}
	}

	return 0
}

// [1,1,2,1,1,1,3,1,2,3]

// 1: 6
// 2: 2
// 3: 2
