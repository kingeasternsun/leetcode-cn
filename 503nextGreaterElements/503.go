/*
 * @Description: 503
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-06 13:04:11
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-06 14:00:10
 * @FilePath: /503nextGreaterElements/503.go
 */

/*
构造递减栈，保存的是nums中的下标，
当要加入一个新 nums[i] value的时候，从栈顶到栈地逐个比较，
	如果栈顶比value小，这个value就是比栈顶 nums[j]大的第一个元素，然后弹出栈顶
	最后终止后，将i加入到栈
*/
package leetcode

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	cnt := len(nums)
	res := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		res[i] = -1
	}
	decreStack := make([]int, 0)
	for i := 0; i < cnt*2; i++ {

		for len(decreStack) > 0 && nums[decreStack[len(decreStack)-1]] < nums[i%cnt] {
			res[decreStack[len(decreStack)-1]] = nums[i%cnt]
			decreStack = decreStack[:len(decreStack)-1]
		}

		decreStack = append(decreStack, i%cnt)

	}
	return res
}
