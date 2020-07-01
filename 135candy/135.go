package main

import "fmt"

func main() {

}

/*
 135. 分发糖果 https://leetcode-cn.com/problems/candy/
*/
func candy1(ratings []int) int {

	count := len(ratings)
	if count <= 1 {
		return count
	}

	candys := make([]int, count)
	candys[0] = 1
	for i := 1; i < count; i++ { //dp[i]之前的是已经遵循规则的了
		if ratings[i] == ratings[i-1] { //如果相等，那么新来这个就给最少的
			candys[i] = 1
		} else if ratings[i] > ratings[i-1] { //如果比左边的分数高，那么这个就多给一个就可以了
			candys[i] = candys[i-1] + 1
		} else { //比左边的分数低，

			//如果左边的糖果大于1 ，那么当前这个给最少就可以了
			if candys[i-1] > 1 {
				candys[i] = 1
			} else { //如果左边的只有1个，那么左边的就要加一，然后要再往左边判断
				candys[i] = 1
				update(ratings, candys, i-1)

			}

		}

		fmt.Println(candys)
	}

	sum := 0
	for i := 0; i < count; i++ {
		sum = sum + candys[i]
	}

	fmt.Println(candys)
	return sum
}

//第i个当前为1， 要加一
func update(ratings []int, candys []int, i int) {
	for i >= 0 {

		candys[i] = candys[i] + 1
		if i == 0 {
			return
		}

		// 如果i-1的分数比较低，当前i个糖果增加了，依然满足规则，就可以返回了
		// 如果i-1的分数和当前的分数相同，依然满足规则，可以直接返回了
		if ratings[i-1] <= ratings[i] {
			return
		}

		//到这里，表示  ratings[i-1] > ratings[i] , 如果当前i的糖果数目增加了，i-1 的糖果数目依然很多，那么依然满足规则，可以返回
		if candys[i-1] > candys[i] {
			return
		}

		i--

	}

	return
}

//20ms 6.1MB
/*
1. 一个数据left，只要比左边大就加一，否则就置为1
2. 一个数组right，从右到左，只要比右边大，就加一，否则置为1
3. 现在需要糖果既要对左边满足规则，又要对右边满足规则，取max即可
*/
func candy(ratings []int) int {

	count := len(ratings)
	if count <= 1 {
		return count
	}

	left := make([]int, count)
	right := make([]int, count)
	left[0] = 1
	right[count-1] = 1

	for i := 1; i < count; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}

		if ratings[count-i-1] > ratings[count-i] {
			right[count-i-1] = right[count-i] + 1
		} else {
			right[count-i-1] = 1
		}
	}

	sum := 0
	for i := 0; i < count; i++ {
		sum += max(left[i], right[i])
	}
	return sum

}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
