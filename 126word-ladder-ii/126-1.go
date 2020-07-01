package main

/*
126. 单词接龙 II https://leetcode-cn.com/problems/word-ladder-ii/
*/

func findLadders1(beginWord string, endWord string, wordList []string) [][]string {

	preMap := make(map[string][]string, 0)
	preMap = bfs1(beginWord, endWord, wordList) //获取链表
	if preMap == nil {
		return nil
	}

	var prePath []string
	var result [][]string
	queryLadders1(endWord, prePath, &result, preMap) //从后往前 再深度优先
	return result
}

func bfs1(beginWord string, endWord string, wordList []string) map[string][]string {
	has := false
	for i := range wordList {
		if wordList[i] == endWord {
			has = true
			break
		}

	}
	if !has {
		return nil
	}

	preMap := make(map[string][]string, 0) //记录当前单词的父单词

	wordLen := len(wordList)
	used := make([]bool, wordLen)

	q := make([]string, 0)
	// m := make(map[string]struct{}, 0)

	q = append(q, beginWord)
	qLen := 1
	step := 1

	for i := range wordList {
		if wordList[i] == beginWord {
			used[i] = true //这里很重要
			break
		}

	}

	// m[beginWord] = struct{}{}

	for qLen > 0 {

		// newLen := 0
		qMap := make(map[int]struct{}, 0) //保存下一层的节点
		// indexQ := make([]int, 0)
		for i := 0; i < qLen; i++ {

			cur := q[i]
			if cur == endWord {
				return preMap
			}

			for j := range wordList {

				if used[j] == false && hasOneDiff(cur, wordList[j]) {
					// q = append(q, wordList[j])
					// newLen++
					qMap[j] = struct{}{} // 和127题的区别，在同一层遍历的时候，j可能会重复访问，所以要用map
					// used[j] = true 		//和127题的区别，这里不能就添加为访问过了，需要在这一层遍历过后统一设置used标记，不然就缺少了其它路径
					preMap[wordList[j]] = append(preMap[wordList[j]], cur)

				}

			}

		}

		//清空当前层
		q = q[:0:0]
		qLen = 0
		// 将下一层更新为当前层
		for k := range qMap {
			q = append(q, wordList[k])
			qLen++

			used[k] = true //统一设置为访问过
		}
		// q = q[qLen:]
		// qLen = len()

		step++
	}

	return nil

}

// 当前节点 root， prePath 是root之前的节点路径，不包含root
func queryLadders1(root string, prePath []string, result *[][]string, parentMap map[string][]string) {
	prePath = append([]string{root}, prePath...)

	parents := parentMap[root]
	if len(parents) == 0 {
		prePath = append(prePath[:0:0], prePath...)
		*result = append(*result, prePath)
		return
	}
	for _, node := range parents {
		queryLadders1(node, prePath, result, parentMap)

	}
	return

}
