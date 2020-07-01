package main

/*
126. 单词接龙 II https://leetcode-cn.com/problems/word-ladder-ii/
*/
func main() {

}

//因为 wordlist中各个单词是唯一的，所以使用索引代表 单词 300ms 5.9MB
func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	has := false
	count := len(wordList)
	endID := 0
	for endID = 0; endID < count; endID++ {
		if wordList[endID] == endWord {
			has = true
			break
		}
	}
	if !has {
		return nil
	}

	has = false
	i := 0
	for i = 0; i < count; i++ {
		if wordList[i] == beginWord {
			has = true
			break
		}

	}

	// 如果 begigword 不在 wordlist,就加入
	if i == count {
		wordList = append(wordList, beginWord)
	}

	preMap := make(map[int][]int, 0)
	preMap = bfs(beginWord, endWord, wordList) //获取链表
	if preMap == nil {
		return nil
	}

	var prePath []string
	var result [][]string
	queryLadders(endID, prePath, &result, preMap, wordList) //从后往前 再深度优先
	return result
}

func bfs(beginWord string, endWord string, wordList []string) map[int][]int {
	// has := false
	count := len(wordList)
	preMap := make(map[int][]int, 0) //记录当前单词的父单词
	used := make([]bool, count)

	q := make([]int, 0)
	qLen := 1

	// 将 begID 放入q中
	for begID := 0; begID < count; begID++ {
		if wordList[begID] == beginWord {
			used[begID] = true //这里很重要
			q = append(q, begID)
			break
		}

	}

	for qLen > 0 {

		// newLen := 0
		qMap := make(map[int]struct{}, 0) //保存下一层的节点
		for i := 0; i < qLen; i++ {

			cur := q[i]
			if wordList[cur] == endWord {
				return preMap
			}

			for j := range wordList {

				if used[j] == false && hasOneDiff(wordList[cur], wordList[j]) {
					// q = append(q, wordList[j])
					// newLen++
					qMap[j] = struct{}{} // 和127题的区别，在同一层遍历的时候，j可能会重复访问，所以要用map
					// used[j] = true 		//和127题的区别，这里不能就添加为访问过了，需要在这一层遍历过后统一设置used标记，不然就缺少了其它路径
					preMap[j] = append(preMap[j], cur)

				}

			}

		}

		//清空当前层
		q = q[:0:0]
		qLen = 0
		// 将下一层更新为当前层
		for k := range qMap {
			q = append(q, k)
			qLen++

			used[k] = true //统一设置为访问过
		}
		// q = q[qLen:]
		// qLen = len()
		// step++
	}

	return nil

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

// 当前节点 root， prePath 是root之前的节点路径，不包含root
func queryLadders(rootID int, prePath []string, result *[][]string, parentMap map[int][]int, wordList []string) {
	prePath = append([]string{wordList[rootID]}, prePath...)

	parents := parentMap[rootID]
	if len(parents) == 0 {
		prePath = append(prePath[:0:0], prePath...)
		*result = append(*result, prePath)
		return
	}
	for _, nodeID := range parents {
		queryLadders(nodeID, prePath, result, parentMap, wordList)

	}
	return

}
