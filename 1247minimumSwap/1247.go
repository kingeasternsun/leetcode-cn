package main

func main() {

}

/*
  根据样例我们可以知道
  - s1:[x x] s2:[y y] 这种只需要交换一次
  - s1:[y y] s2:[x x] 这种只需要交换一次
  - s1:[x y] s2:[y x] 这种只需要交换两次

  我们遍历s1和s2，查询
  - 相同位置上  s1上是x，s2上时y的字符对 (x,y) 个数 xyCnt
  - 相同位置上  s1上是y，s2上时x的字符对 (y,x) 个数 yxCnt
  - 位置上相同的肯定就不需要移动了

  我们每次选取两对(x,y),对应 s1:[x x] s2:[y y] 的场景， 交换一次
  我们每次选取两对(y,x),对应 s1:[y y] s2:[x x] 的场景， 交换一次

  最后只能是下面的情况：
  1. 只剩下 (x,y) //无法相同
  2. 只剩下 (y,x) //无法相同
  3. 剩下 (y,x) 和 (x,y) ，对应 s1:[x y] s2:[y x] ,需要交换两次

*/

func minimumSwap(s1 string, s2 string) int {

	if len(s1) != len(s2) {
		return -1
	}

	xyCnt := 0
	yxCnt := 0

	for i := 0; i < len(s1); i++ {

		if s1[i] != s2[i] {

			if s1[i] == 'x' {

				xyCnt++
			} else {
				yxCnt++
			}

		}

	}

	if (xyCnt+yxCnt)&1 > 0 { //对应注释中的 只剩下 (x,y)或(y,x)
		return -1
	}

	res := xyCnt/2 + yxCnt/2
	if xyCnt&1 > 0 {
		res += 2
	}

	return res
}
