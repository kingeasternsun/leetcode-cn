package main

func main() {

}

// https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/
// partition-array-into-three-parts-with-equal-sum
func canThreePartsEqualSum(A []int) bool {

	sum := 0
	count := len(A)
	if count == 0 {
		return false
	}

	for i := 0; i < count; i++ {
		sum += A[i]
	}

	eachSum := sum / 3
	return help(A, count, sum, 3, eachSum)
}

/*
 aLen: 当前数组的长度
 sum ：当前数组的总数
 count:   分为几个部分
 eachNum:  每一部分的和多少
*/
func help(A []int, aLen, sum, count, eachNum int) bool {

	if sum != count*eachNum {
		return false
	}

	if aLen == 0 && count == 0 {
		return true
	}

	curSum := 0
	for i := 0; i < aLen; i++ {
		curSum += A[i]
		if curSum == eachNum {
			result := help(A[i+1:], aLen-i-1, sum-curSum, count-1, eachNum) //剩下的数组是否可以 组成 count-1 个 eachNum
			if result == true {
				return true
			}
		}

	}

	return false
}
