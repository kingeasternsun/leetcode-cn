/*
 * @Author: kingeasternsun
 * @Date: 2020-03-15 16:59:55
 * @LastEditTime: 2020-03-15 17:06:11
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \leetcode\263\263.go
 */

//https://leetcode-cn.com/problems/ugly-number/
package main

func main() {

}

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	if num == 1 {
		return true
	}

	if num%5 == 0 {
		return isUgly(num / 5)
	}

	if num%3 == 0 {
		return isUgly(num / 3)
	}

	if num%2 == 0 {
		return isUgly(num / 2)
	}

	return false

}
