/*
 * @Description: 697
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-20 10:50:19
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-20 11:48:18
 * @FilePath: \697findShortestSubArray\697.go
 */
package leetcode

/*
给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。

你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/degree-of-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

记录每个元素的开始位置，结束位置，次数
*/

type Value struct {
	Beg, End int
	Count    int
}

func findShortestSubArray(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	m := make(map[int]Value, 0)
	for i, n := range nums {
		v, ok := m[n]
		if !ok {
			m[n] = Value{Beg: i, End: i, Count: 1}
			continue
		}

		v.Count++
		v.End = i
		m[n] = v
	}

	res := Value{}
	for _, v := range m {
		if v.Count > res.Count {
			res = v
		} else if v.Count == res.Count && (v.End-v.Beg) < (res.End-res.Beg) {
			res = v
		}
	}

	return res.End - res.Beg + 1
}
