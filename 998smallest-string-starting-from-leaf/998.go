package main

import "fmt"

func main() {

	a := 'a'
	fmt.Printf("%c", a+1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
988. 从叶结点开始的最小字符串 https://leetcode-cn.com/problems/smallest-string-starting-from-leaf/
*/
func smallestFromLeaf(root *TreeNode) string {

	if root == nil {
		return ""
	}

	var curList []byte
	minStr := ""
	dfs(root, curList, &minStr)
	return minStr
}

// curList 表示当前节点之前的父节点的列表，不包含当前节点；且 cur不为nil
// 0ms 4.6MB
func dfs(cur *TreeNode, curList []byte, minStr *string) {

	// 因为 curList 进入的时候不包含当前节点，所以这里要追加
	// curList = append(curList)
	curList = append([]byte{byte('a' + cur.Val)}, curList...) //这里要注意 是从叶节点到根节点
	if cur.Left == nil && cur.Right == nil {

		if *minStr == "" {
			*minStr = string(curList)
		} else if string(curList) < *minStr {
			*minStr = string(curList)
		}
		return
	}

	if cur.Left != nil { //如果没有这个判断，那么进入函数的开始地方 就要 判断非空
		dfs(cur.Left, curList, minStr)
	}

	if cur.Right != nil { //如果没有这个判断，那么进入函数的开始地方 就要 判断非空
		dfs(cur.Right, curList, minStr)
	}
	return
}
