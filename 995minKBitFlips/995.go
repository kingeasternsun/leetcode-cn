/*
 * @Description: 995
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-18 19:18:41
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-18 20:08:33
 * @FilePath: /995minKBitFlips/995.go
 */
package leetcode

//https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips/solution/k-lian-xu-wei-de-zui-xiao-fan-zhuan-ci-s-dseq/
/*
如何快速知道当前位置已经被翻转了多少次， 当前位置i的被动翻转次数只跟 (i-k, i-1]的翻转次数有关
*/
func minKBitFlips(A []int, K int) int {

	res := 0
	//滑动窗口
	win := make([]int, 0)
	for i := range A {

		// 如果左边超过范围了，就移除左边那个，这个要放在 循环的最开始
		if len(win) > 0 && win[0] <= i-K {
			win = win[1:]
		}
		//原本是0，翻转偶数次还是0，需要主动翻转下当前
		//原本是1，翻转奇数次变为0，需要主动翻转下当前
		if len(win)&1 == A[i] {
			//如果需要翻转当前i，但是没有足够的可以翻转
			if i+K > len(A) {
				return -1
			}

			res++
			//加入窗口
			win = append(win, i)

		}
		//当前变为1了不需要翻转

	}

	return res
}
