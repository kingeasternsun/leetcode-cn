package main

import "math"

func main() {

}

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//使用dp，跟 robber 的那个题很像
func minCameraCover(root *TreeNode) int {

	if root.Left == nil && root.Right == nil {
		return 1
	}

	inMap := make(map[*TreeNode]int, 0)
	unMap := make(map[*TreeNode]int, 0)
	un2Map := make(map[*TreeNode]int, 0)
	return min(help(root, 1, inMap, unMap, un2Map), help(root, 3, inMap, unMap, un2Map))
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

//当前节点 装摄像头；当前节点没有装摄像头，但是当前节点被父节点监控；当前节点没有装摄像头，但是当前节点没有被父亲节点监控
func help(root *TreeNode, install int, inMap, un2, un3 map[*TreeNode]int) int {

	if root == nil {
		return 0
	}

	num := 0

	//当前节点安装摄像头
	if install == 1 {
		//已经计算过 就缓存直接返回
		if result, ok := inMap[root]; ok {
			return result
		}

		//当前安装，子节点可以安装或不安装
		num := min(help(root.Left, 1, inMap, un2, un3), help(root.Left, 2, inMap, un2, un3)) + 1 +
			min(help(root.Right, 1, inMap, un2, un3), help(root.Right, 2, inMap, un2, un3))

		inMap[root] = num
		return num

	}

	// 当前节点没有装摄像头，但是当前节点被父节点监控；
	if install == 2 {
		//已经计算过 就缓存直接返回
		if result, ok := un2[root]; ok {
			return result
		}

		//子节点可以安装，也可以不安装
		num = min(help(root.Left, 1, inMap, un2, un3), help(root.Left, 3, inMap, un2, un3)) +
			min(help(root.Right, 1, inMap, un2, un3), help(root.Right, 3, inMap, un2, un3))

		un2[root] = num
		return num

	}

	// 当前节点没有装摄像头，而且当前节点没有被父亲节点监控,这里最关键点，在于子节点只要有一个安装监控就可以
	// if install == 3 {
	//已经计算过 就缓存直接返回
	if result, ok := un3[root]; ok {
		return result
	}

	if root.Left == nil && root.Right == nil {
		num = 1
	} else {

		num = math.MaxInt32
		if root.Left != nil { //左边安装监控，右边可安装，可不安装
			num = help(root.Left, 1, inMap, un2, un3) + min(help(root.Right, 1, inMap, un2, un3), help(root.Right, 3, inMap, un2, un3))
		}

		if root.Right != nil { //右边安装监控，左边可安装，可不安装
			num2 := help(root.Right, 1, inMap, un2, un3) + min(help(root.Left, 1, inMap, un2, un3), help(root.Left, 3, inMap, un2, un3))

			if num2 < num {
				num = num2
			}

		}

	}

	un3[root] = num
	return num

	// }

}
