package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {

	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	if root == nil {
		return "[null]"
	}

	nodes := []*TreeNode{root}
	code := "[" + strconv.Itoa(root.Val)

	for len(nodes) > 0 {
		newNodes := []*TreeNode{}

		for i := range nodes {

			if nodes[i].Left != nil {
				newNodes = append(newNodes, nodes[i].Left)
				code = code + "," + strconv.Itoa(nodes[i].Left.Val)
			} else {
				code = code + ",null"
			}

			if nodes[i].Right != nil {
				newNodes = append(newNodes, nodes[i].Right)
				code = code + "," + strconv.Itoa(nodes[i].Right.Val)
			} else {
				code = code + ",null"
			}

		}

		nodes = newNodes

	}
	code = code + "]"
	return code
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	if len(data) == 0 {
		return nil
	}

	code := data[1:]
	code = code[:len(code)-1]
	nodeStr := strings.Split(code, ",")
	if len(nodeStr) == 0 {
		return nil
	}

	if nodeStr[0] == "null" {
		return nil
	}

	nodes := []*TreeNode{}
	v, _ := strconv.Atoi(nodeStr[0])
	root := &TreeNode{
		Val: v,
	}
	nodes = append(nodes, root)

	nodeStr = nodeStr[1:]
	for len(nodeStr) > 0 {

		newNodes := []*TreeNode{}
		childNum := 2 * len(nodes)
		if childNum > len(nodeStr) {
			childNum = len(nodeStr)
		}

		for j := 0; j < childNum; j++ {
			var curNode *TreeNode
			if nodeStr[j] != "null" {
				v, _ := strconv.Atoi(nodeStr[j])
				curNode = &TreeNode{
					Val: v,
				}

				newNodes = append(newNodes, curNode)
			}

			parNode := nodes[j/2]
			if j&1 == 0 {
				parNode.Left = curNode
			} else {
				parNode.Right = curNode
			}
		}

		nodeStr = nodeStr[childNum:]
		nodes = newNodes
	}

	return root
}

func main() {
	var a []interface{}
	a = append(a, 1)
	a = append(a, nil)
	c := fmt.Sprintf("%v", a)
	println(c)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
