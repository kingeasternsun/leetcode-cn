/*
 * @Description:1
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-17 22:13:37
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-17 22:32:20
 * @FilePath: /1twoSum/1.go
 */
package leetcode

/*
 因为要返回下标，所以只能使用hashmap
 不然就可以先排序然后使用双指针来做
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		if j, ok := m[target-v]; ok {

			if v != target-v {
				return []int{i, j}
			}

			if j != i {
				return []int{i, j}
			}
		}
	}
	return nil
}
