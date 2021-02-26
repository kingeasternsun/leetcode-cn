/*
 * @Description: 80
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-26 16:27:11
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-26 17:44:13
 * @FilePath: \80removeDuplicates\80.go
 */
package leetcode

//之前只要判断和前一个是否相同，这里加个判断是否和 j-2相同就可以了
// j-2有可能被覆盖，所以要提前保存
func removeDuplicates(nums []int) int {

	if len(nums) <= 2 {
		return len(nums)
	}

	i, j, prePre := 1, 2, nums[0]
	for j < len(nums) {
		if nums[j] != nums[j-1] || nums[j] != prePre {
			prePre = nums[j-1]
			i++
			nums[i] = nums[j] //这里有可能覆盖掉 nums[j-1] 所以要提前保存
			j++
			continue
		}
		prePre = nums[j-1]
		j++

	}

	return i + 1
}
