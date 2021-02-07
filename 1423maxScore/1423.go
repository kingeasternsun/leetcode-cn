/*
 * @Description:1423
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-06 10:06:07
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-06 11:41:29
 * @FilePath: \1423maxScore\1423.go
 */
package leetcode

func maxScore1(cardPoints []int, k int) int {
	//将cardPoints首位相接， 可以认为给 cardPoints 添加一个虚拟的k的长度的数据。
	if k == 1 {
		if cardPoints[0] > cardPoints[len(cardPoints)-1] {
			return cardPoints[0]
		}
		return cardPoints[len(cardPoints)-1]
	}

	tmpRes := 0
	for i := len(cardPoints) - k; i < len(cardPoints); i++ {
		tmpRes += cardPoints[i]
	}

	if k == len(cardPoints) {
		return tmpRes
	}

	//执行滑窗
	res := tmpRes
	for i := len(cardPoints); i < len(cardPoints)+k; i++ {
		tmpRes += cardPoints[i%len(cardPoints)]
		tmpRes -= cardPoints[i-k]
		if tmpRes > res {
			res = tmpRes
		}
	}

	return res
}

func maxScore(cardPoints []int, k int) int {

	//将cardPoints首位相接， 可以认为给 cardPoints 添加一个虚拟的k的长度的数据。
	tmpRes, res := 0, 0
	for i := len(cardPoints) - k; i < len(cardPoints)+k; i++ {
		if i < len(cardPoints) {
			tmpRes += cardPoints[i]
			res = tmpRes
		} else {
			tmpRes += cardPoints[i%len(cardPoints)]
			tmpRes -= cardPoints[i-k]
			if tmpRes > res {
				res = tmpRes
			}
		}

	}

	return res
}
