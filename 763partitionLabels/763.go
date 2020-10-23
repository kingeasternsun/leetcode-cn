package main

import (
	"math"
	"sort"
)

func main() {
	partitionLabels("ababcbacadefegdehijhklij")
}

func partitionLabels(S string) []int {

	part := make([][2]int, 26) //26个英文字母 在S中的最左边的位置  和 最右边的位置
	for i := 0; i < 26; i++ {
		part[i][0] = math.MaxInt32
		part[i][1] = -1
	}

	//得到每个字母所占用的区间
	cnt := len(S)
	for i := 0; i < cnt; i++ {
		partID := S[i] - 'a'
		if i < part[partID][0] {
			part[partID][0] = i
		}

		if i > part[partID][1] {
			part[partID][1] = i
		}
	}
	// fmt.Printf("%+v", part)

	//根据开始位置进行排序
	sort.Slice(part, func(i, j int) bool {

		return part[i][0] < part[j][0]
	})

	// fmt.Printf("%+v", part)
	//最后将交叉的区间合并
	for i := 1; i < 26; i++ {
		if part[i][1] == -1 { //表示没有区间了
			break
		}

		if part[i][0] > part[i-1][1] { //不重叠 直接掉过
			continue
		}

		if part[i][1] < part[i-1][1] { //当前区间在上一个区间里面
			part[i] = part[i-1] //抛弃小的区间
			part[i-1][1] = -1   //前面交叉的区间变为无效
		}

		//区间交叉, 合并两个区间
		part[i][0] = part[i-1][0]
		part[i-1][1] = -1 //前面交叉的区间变为无效

	}
	// fmt.Printf("%+v", part)
	//选取part[i][1]不为-1的，计算宽度返回
	result := make([]int, 0)
	for _, v := range part {

		if v[1] == -1 {
			continue
		}

		result = append(result, v[1]-v[0]+1)
	}
	// fmt.Printf("%+v", result)

	return result
}
