/*
 * @Description: 1052
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-23 09:04:46
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-23 09:30:57
 * @FilePath: \1052maxSatisfied\1052.go
 */
package leetcode

/*
 - 计算不考虑X时候所能得到的满意
 - 滑动窗口 计算这个区间所能得到额外满意度
*/
func maxSatisfied(customers []int, grumpy []int, X int) int {

	//不考虑X的满意度
	totalSum := 0
	for i := range customers {
		if grumpy[i] == 0 { //没有生气
			totalSum += customers[i]
		}
	}

	resDelta := 0
	tmp := 0
	for i := 0; i < X; i++ {
		if grumpy[i] == 1 { //这个时刻本来是生气，变为了不生气
			tmp += customers[i]
		}
	}
	resDelta = tmp
	for i := X; i < len(customers); i++ {

		if grumpy[i] == 1 { //这个时刻本来是生气，变为了不生气
			tmp += customers[i]
		}

		if grumpy[i-X] == 1 {
			tmp -= customers[i-X]
		}

		if tmp > resDelta {
			resDelta = tmp
		}

	}

	return resDelta + totalSum

}
