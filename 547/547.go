//https://leetcode-cn.com/problems/friend-circles/
// 547. 朋友圈/

package main

/*
方法 可以参考 最小代价生成树中 合并子树的方法
*/
func findFather(father []int, i int) int {

	for father[i] != i {
		i = father[i]
	}

	return i

}

// 24ms 6.2MB
func uni(father []int, i, j int) {
	i = findFather(father, i)
	j = findFather(father, j)
	father[i] = j

}

func findCircleNum(M [][]int) int {

	mLen := len(M)
	if mLen <= 1 {
		return mLen
	}

	father := make([]int, mLen)
	for i := 0; i < mLen; i++ {
		father[i] = i
	}

	for i := 0; i < mLen; i++ {
		for j := i + 1; j < mLen; j++ {
			if M[i][j] == 1 {
				uni(father, i, j)
			}
		}
	}

	count := 0
	for i := 0; i < mLen; i++ {
		if father[i] == i {
			count++
		}
	}

	return count

}
func uniOp(father, height []int, i, j int) {
	i = findFather(father, i)
	j = findFather(father, j)

	if height[i] == height[j] {
		father[i] = j
		height[j]++
	}

	if height[i] < height[j] {
		father[i] = j
	} else {
		father[j] = i
	}

}

// 48ms
func findCircleNumOp(M [][]int) int {

	mLen := len(M)
	if mLen <= 1 {
		return mLen
	}

	father := make([]int, mLen)
	height := make([]int, mLen)
	for i := 0; i < mLen; i++ {
		father[i] = i
		height[i] = 1
	}

	for i := 0; i < mLen; i++ {
		for j := i + 1; j < mLen; j++ {
			if M[i][j] == 1 {
				uniOp(father, height, i, j)
			}
		}
	}

	count := 0
	for i := 0; i < mLen; i++ {
		if father[i] == i {
			count++
		}
	}

	return count

}

func findFatherZip(father []int, i int) int {

	j := i
	for father[i] != i {
		i = father[i]
	}

	//通过压缩，树最高三层，所以把最开始的指向树顶就可以
	father[j] = i
	return i

}
func uniZip(father []int, i, j int) {
	i = findFatherZip(father, i)
	j = findFatherZip(father, j)
	father[i] = j

}
func findCircleNumZip(M [][]int) int {

	mLen := len(M)
	if mLen <= 1 {
		return mLen
	}

	father := make([]int, mLen)
	for i := 0; i < mLen; i++ {
		father[i] = i
	}

	for i := 0; i < mLen; i++ {
		for j := i + 1; j < mLen; j++ {
			if M[i][j] == 1 {
				uniZip(father, i, j)
			}
		}
	}

	count := 0
	for i := 0; i < mLen; i++ {
		if father[i] == i {
			count++
		}
	}

	return count

}
