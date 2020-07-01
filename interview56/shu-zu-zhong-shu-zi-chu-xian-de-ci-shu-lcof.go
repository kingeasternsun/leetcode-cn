package main

func main() {

}
/*

https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
*/
/*

1. 最终的异或结果c是要查找的两个数字a和b的异或
2. c中某一位为1，那么a和b中，其中一个在这一位上也是1，另外一个数在这个位上肯定不是1；假设a在这一位上是1，
3. 把数组nums中这个位是1的数 和 c 进行异或。由于除了a和b，其他数字都出现了两次，所以异或的结果对c不会有影响。
   而a只出现了一次，所以c和a异或就得到了b。

*/
func singleNumbers(nums []int) []int {
	nLen := len(nums)
	if nLen <= 2 {
		return nil
	}

	c := 0
	for i := 0; i < nLen; i++ {
		c = c ^ nums[i]
	}

	// 查询c中某位是1的位置
	var p uint
	for {
		if ((1 << p) & c) > 0 {
			break
		}
		p++
	}

	b := c
	for i := 0; i < nLen; i++ {

		if ((1 << p) & nums[i]) > 0 {
			b = b ^ nums[i]
		}

	}

	a := b ^ c

	return []int{a, b}
}
