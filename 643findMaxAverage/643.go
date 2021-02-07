/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-04 10:03:59
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-04 10:13:15
 * @FilePath: \643findMaxAverage\643.go
 */
package leetcode

func findMaxAverage(nums []int, k int) float64 {

	sum := 0
	for _, v := range nums[:k] {
		sum += v
	}

	res := float64(sum) / float64(k)
	for i := k; i < len(nums); i++ {
		sum += (nums[i] - nums[i-k])
		tmp := float64(sum) / float64(k)
		if tmp > res {
			res = tmp
		}
	}

	return res

}
