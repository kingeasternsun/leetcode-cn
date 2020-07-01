package main

func main() {

}

//https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
//https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/solution/san-jin-zhi-zhuang-tai-ji-by-muyids/

func singleNumber(nums []int) int {

	nLen := len(nums)
	if nLen == 1 {
		return nums[0]
	}

	one := 0
	two := 0

	//构建三进制
	for i := 0; i < nLen; i++ {
		one = (one ^ i) & (^two + 1)
		two = (two ^ i) & (^one + 1)
	}
	return one
}
