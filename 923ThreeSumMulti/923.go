package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8) == 20)
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 2, 2}, 5) == 12)
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 8) == 100)
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 3) == 220)
	fmt.Println(threeSumMulti([]int{0, 0, 0}, 0) == 1)

	fmt.Println(threeSumMultiMap([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8) == 20)
	fmt.Println(threeSumMultiMap([]int{1, 1, 2, 2, 2, 2}, 5) == 12)
	fmt.Println(threeSumMultiMap([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 8) == 100)
	fmt.Println(threeSumMultiMap([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 3) == 220)
	fmt.Println(threeSumMultiMap([]int{0, 0, 0}, 0) == 1)

}

// 给定一个整数数组 A，以及一个整数 target 作为目标值，返回满足 i < j < k 且 A[i] + A[j] + A[k] == target 的元组 i, j, k 的数量。

// 由于结果会非常大，请返回 结果除以 10^9 + 7 的余数。
// 3 <= A.length <= 3000
// 0 <= A[i] <= 100
// 0 <= target <= 300
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum-with-multiplicity
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//转为求两数目之和
func threeSumMultiSlow(A []int, target int) int {
	cnt := len(A)
	if cnt < 3 {
		return 0
	}

	sort.Ints(A)
	num := 0
	mod := 1000000007

	for i := 0; i < cnt-1; i++ {
		tmpTarget := target - A[i]
		if tmpTarget < 0 {
			continue
		}

		if A[i]+A[i+1] > tmpTarget && A[cnt-2]+A[cnt-1] < tmpTarget {
			continue
		}

		left := i + 1
		right := cnt - 1

		for left < right {

			if left == i {
				left++
				continue
			}

			if right == i {
				right--
				continue
			}

			switch sum := A[left] + A[right]; {
			case sum == tmpTarget:

				//两端相同，中间的肯定也相同 ,从里面选2两个出来
				if A[left] == A[right] {

					if A[i] == A[left] {
						//刚好时3分之一
						num += (right - i + 1) * (right - i) * (right - i - 1) / 6

						left = right
						i = right
					} else {
						num += (right - left + 1) * (right - left) / 2
						left = right
					}

				} else { //两端不相同

					leftNum := 0 //left相同的个数
					leftVal := A[left]

					rightNum := 0 //right 相同的个数
					rightVal := A[right]
					for left < right && A[left] == leftVal {
						leftNum++
						left++
					}
					left-- //这里要回退以下  重要 关键
					for left < right && A[right] == rightVal {
						rightNum++
						right--
					}

					// fmt.Println(num)
					num += (leftNum * rightNum)

				}

				num = num % mod
			case sum > tmpTarget:
				right--
			case sum < tmpTarget:
				left++

			}

		}

	}

	return num
}

//该题目中 A[i]<=300 的数值的大小是有限的，所以我们可以统计各个数值的分布，再进行计算.变为从0 到100里面取三个数，是否构成 target
// 4ms 3.1MB
func threeSumMulti(A []int, target int) int {

	num := 0
	mod := 1000000007

	//记录每个数字的个数
	nums := make([]int, 101)

	for _, v := range A {
		nums[v] = nums[v] + 1
	}

	for i := 0; i < 101; i++ {
		if nums[i] == 0 {
			continue
		}

		// 这里关键点在于 j可以等于i
		for j := i; j < 101; j++ {

			if nums[j] == 0 {
				continue
			}

			k := target - i - j
			if k < j || k > 100 {
				continue
			}
			if nums[k] == 0 {
				continue
			}

			// i <= j <= k
			if i == j && j == k {

				num += (nums[i]) * (nums[i] - 1) * (nums[i] - 2) / 6
				num = num % mod
			} else if i == j {
				num += (nums[i]) * (nums[i] - 1) / 2 * (nums[k])
				num = num % mod
			} else if j == k {
				num += (nums[i]) * (nums[k]) * (nums[k] - 1) / 2
				num = num % mod
			} else {
				num += (nums[i]) * (nums[j]) * (nums[k])
				num = num % mod
			}

		}
	}

	return num

}

//该题目中 A[i]<=300 的数值的大小是有限的，所以我们可以统计各个数值的分布，再进行计算.变为从0 到100里面取三个数，是否构成 target
//8ms 3.2MB
func threeSumMultiMap(A []int, target int) int {

	num := 0
	mod := 1000000007

	//记录每个数字的个数
	nums := make(map[int]int, 0)

	for _, v := range A {
		nums[v] = nums[v] + 1
	}

	for i := 0; i < 101; i++ {
		if nums[i] == 0 {
			continue
		}

		// 这里关键点在于 j可以等于i
		for j := i; j < 101; j++ {

			if nums[j] == 0 {
				continue
			}

			k := target - i - j
			if k < j || k > 100 {
				continue
			}
			if nums[k] == 0 {
				continue
			}

			// i <= j <= k
			if i == j && j == k {

				num += (nums[i]) * (nums[i] - 1) * (nums[i] - 2) / 6
				num = num % mod
			} else if i == j {
				num += (nums[i]) * (nums[i] - 1) / 2 * (nums[k])
				num = num % mod
			} else if j == k {
				num += (nums[i]) * (nums[k]) * (nums[k] - 1) / 2
				num = num % mod
			} else {
				num += (nums[i]) * (nums[j]) * (nums[k])
				num = num % mod
			}

		}
	}

	return num

}
