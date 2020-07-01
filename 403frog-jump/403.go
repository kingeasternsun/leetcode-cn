package main

func main() {

}

/*
403. 青蛙过河 https://leetcode-cn.com/problems/frog-jump/
*/
func canCross(stones []int) bool {

	count := len(stones)
	if count == 0 {
		return false
	}

	hasStone := make(map[int]int, 0) //记录某个位置是第几个石头
	for i := 0; i < count; i++ {
		hasStone[stones[i]] = i
	}

	dp := make([][]int, stones[count-1]) //记录每个石头是青蛙从前面的石头跳了哪些步这个石头上，不用管是从哪个石头过来的
	dp[0] = []int{0}

	for i := 0; i < count; i++ {
		steps := dp[i]
		for _, step := range steps { // 第i个石头可以通过跳 step 过来

			//从第i个石头 可以往后跳的步骤
			for newStep := step - 1; newStep < step+2; newStep++ {
				if newStep > 0 {

					//第i个石头在stones【i】个位置上，再跳 newStep 到位置 stones[i]+newStep  位置上
					stoneID, ok := hasStone[stones[i]+newStep] // 如果 stones[i]+newStep 位置上 包含有石头 stoneID
					if ok {
						dp[stoneID] = append(dp[stoneID], newStep)

					}

				}
			}

		}
	}

	return len(dp[count-1]) > 0

}
