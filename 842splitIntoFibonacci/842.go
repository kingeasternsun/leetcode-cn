package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(splitIntoFibonacci("123456579"))
}
func splitIntoFibonacci(S string) []int {

	var res []int

	if help([]byte(S), 0, 0, 0, &res) {
		return res
	}

	return res
}

// curID 表示接下来要处理第几个字符  preSum表示前面两个数字的和，pre表示前一个数字, res用来存放结果
func help(s []byte, curID int, preSum int, pre int, res *[]int) bool {

	if curID == len(s) {
		return len(*res) >= 3
	}

	cur := 0 //新的数字

	for i := curID; i < len(s); i++ {

		if i > curID && s[curID] == '0' { //如果cur这个数字长度大于1，而且第一个字符(curID所在字符)是0，就非法 可以提前剪枝了
			break
		}

		cur = cur*10 + int(s[i]-'0')

		if cur > math.MaxInt32 { //如果当前值过大，可以提前剪枝了
			break
		}

		//如果前面已经超过两个数字 ,就要判断是否等于前两个之和
		if len(*res) >= 2 {
			if cur < preSum {
				continue
			}

			if cur > preSum { //当前cur已经大于前两个之和了，就可以提前剪枝
				break
			}
		}

		// res中还不够2个 或 满足条件
		*res = append(*res, cur)

		if help(s, i+1, pre+cur, cur, res) {
			return true
		}

		*res = (*res)[:len(*res)-1]

	}

	return false

}
