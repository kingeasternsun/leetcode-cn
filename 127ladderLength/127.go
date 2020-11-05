package main

import (
	"math"
)

func main() {

}

/*
https://leetcode-cn.com/problems/word-ladder/
127. 单词接龙

dfs  排列组合
*/

func ladderLengthdfs(beginWord string, endWord string, wordList []string) int {
	// wordList:=[]string
	minWordCount := math.MaxInt32

	has := false
	for i := range wordList {
		if wordList[i] == endWord {
			has = true
			break
		}
	}
	if !has {
		return 0
	}

	//放入了第一个字符 beginWord
	dfs(1, beginWord, endWord, wordList, &minWordCount)

	if minWordCount == math.MaxInt32 {
		return 0
	}
	return minWordCount

}

func dfs(tmpWordCount int, tmpEndword, endWord string, wordList []string, minWordCount *int) {

	if tmpWordCount > *minWordCount {
		return
	}

	if tmpEndword == endWord {

		if tmpWordCount < *minWordCount {
			*minWordCount = tmpWordCount
		}

		return
	}

	//相差一个字符
	// if hasOneDiff(tmpEndword, endWord) {
	// 	tmpWordCount++
	// 	if tmpWordCount < *minWordCount {
	// 		*minWordCount = tmpWordCount
	// 	}

	// 	return
	// }

	for i := range wordList {

		// 每次只能更改一个字符，所以前提是 wordList[i] 和 序列中最后一个字符串 tmpEndword 相差一个字符
		if !hasOneDiff(wordList[i], tmpEndword) {
			// fmt.Printf("%v %v not match \n", wordList[i], tmpEndword)
			continue
		}

		// 选取第i个word之前，先和第0个换下位置，这样就可以做到每次取一个就少一个 不重复
		wordList[0], wordList[i] = wordList[i], wordList[0]
		// 加入更换后词典的第一个world， wordlist然后去除第一个word
		dfs(tmpWordCount+1, wordList[0], endWord, wordList[1:], minWordCount)

		//再恢复回来
		wordList[0], wordList[i] = wordList[i], wordList[0]

	}

}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	return bbfs(beginWord, endWord, wordList)
}

// 是否只有一个字符不一样
func hasOneDiff(a, b string) bool {

	dif := 0

	for i := range a {
		if a[i] != b[i] {
			dif++

			if dif > 1 {
				return false
			}
		}
	}

	return dif == 1

}

func bfs(beginWord string, endWord string, wordList []string) int {
	has := false
	for i := range wordList {
		if wordList[i] == endWord {
			has = true
			break
		}
	}
	if !has {
		return 0
	}

	wordLen := len(wordList)
	used := make([]bool, wordLen)

	q := make([]string, 0)
	// m := make(map[string]struct{}, 0)

	q = append(q, beginWord)
	qLen := 1
	step := 1
	// m[beginWord] = struct{}{}

	for qLen > 0 {

		newLen := 0
		for i := 0; i < qLen; i++ {

			cur := q[i]
			if cur == endWord {
				return step
			}

			for j := range wordList {

				if used[j] == false && hasOneDiff(cur, wordList[j]) {
					q = append(q, wordList[j])
					newLen++
					used[j] = true
				}

			}

		}
		q = q[qLen:]
		qLen = newLen
		step++
	}

	return 0

}

//双边 bfs 参考 https://leetcode-cn.com/problems/word-ladder/solution/shuang-xiang-bfs-by-joy-teng/
func bbfs(beginWord string, endWord string, wordList []string) int {
	has := false
	for i := range wordList {
		if wordList[i] == endWord {
			has = true
			break
		}
	}
	if !has {
		return 0
	}

	wordLen := len(wordList)
	used := make([]int, wordLen)
	for i := range wordList {
		if wordList[i] == endWord {
			used[i] = 2
		} else if wordList[i] == beginWord {
			used[i] = 1
		}
	}

	q1 := make([]string, 0)
	q1 = append(q1, beginWord)
	qLen1 := 1

	q2 := make([]string, 0)
	q2 = append(q2, endWord)
	qLen2 := 1

	step := 1 //总的步长
	step++
	// m[beginWord] = struct{}{}

	for qLen1 > 0 && qLen2 > 0 {

		qLen := qLen1
		q := q1
		flag := 1

		// 从小的开始遍历
		if qLen2 < qLen1 {
			qLen = qLen2
			q = q2
			flag = 2
		}

		newLen := 0
		for i := 0; i < qLen; i++ {

			cur := q[i]
			// if cur == endWord {
			// 	return step
			// }

			for j := range wordList {

				// used[j]&flag  used为1 表示从q1遍历，为2表示从q2遍历 ，为3表示都访问了; used[j]&flag==0 表示这个方向的bfs没有访问过
				if (used[j]&flag) == 0 && hasOneDiff(cur, wordList[j]) {
					q = append(q, wordList[j])
					newLen++
					used[j] = (used[j] | flag)
				}

				if used[j] == 3 {
					return step
				}

			}

		}

		q = q[qLen:]
		qLen = newLen
		step++

		if flag == 1 {
			q1 = q
			qLen1 = qLen
		} else {
			q2 = q
			qLen2 = qLen
		}

	}

	return 0

}
