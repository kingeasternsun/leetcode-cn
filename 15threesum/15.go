/*
 * @Description:15
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2020-07-01 12:44:30
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-19 11:54:00
 * @FilePath: \15threesum\15.go
 */
package leetcode

import "sort"

/*
滑窗
记得去重，去重的时候还要注意 beg > i+1 && nums[beg] == nums[beg-1]
提前返回
*/
func threeSum1(nums []int) [][]int {

	count := len(nums)
	if count < 3 {
		return nil
	}

	sort.Ints(nums)

	result := make([][]int, 0)
	for i, v := range nums {

		//去重
		if i > 0 && v == nums[i-1] {
			continue
		}

		//剩余两个要从后面找，所以和肯定要大于0
		if v > 0 {
			break
		}

		//提前判断返回
		if nums[0]+nums[1] > -v || nums[count-1]+nums[count-2] < -v {
			continue
		}

		beg := i + 1
		end := count - 1

		for beg < end {

			//跳过目标v对应的数组项 ,
			if beg > i+1 && nums[beg] == nums[beg-1] { //去重 beg=i+1的时候，可以一样
				beg++
				continue
			}

			if end < count-1 && nums[end] == nums[end+1] {
				end--
				continue
			}

			if nums[beg]+nums[end] == -v {
				tmpResult := []int{v, nums[beg], nums[end]}
				result = append(result, tmpResult)
				beg++
				end--
				continue
			}

			if nums[beg]+nums[end] > -v {
				end--
				continue
			}

			if nums[beg]+nums[end] < -v {
				beg++
				// continue
			}

		}

	}

	return result

}

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
参考 花花酱 的讲解视频
*/

func threeSumHash(nums []int) [][]int {

	res := make([][]int, 0)

	sort.Ints(nums)
	cntMap := make(map[int]int, 0)
	for _, v := range nums {
		cntMap[v]++
	}

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] { //去重
			continue
		}

		// nums[i]<=nums[j]<=nums[k] 因为nums排好顺序了，只要 j 从i+1 开始就可以
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] { //去重
				continue
			}

			c := 0 - nums[i] - nums[j]
			if c < nums[j] {
				continue
			}

			//判断 是否有c 而且c的数量足够.这里比较关键，不能只判断 cntMap中是否存在c，而且要考虑 nums[i] nums[j]也有可能是c 0+0+0 = 0
			cnt := 1
			if nums[i] == c {
				cnt++
			}
			if nums[j] == c {
				cnt++
			}
			if cntMap[c] < cnt {
				continue
			}

			res = append(res, []int{nums[i], nums[j], c})

		}

	}

	return res

}
