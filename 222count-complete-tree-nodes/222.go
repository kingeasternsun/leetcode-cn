package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodesBiSearch(root *TreeNode) int {

	h := getDepth(root)
	if h <= 1 {
		return int(h)
	}

	// fmt.Println(h)

	nodeID := getNodeID(root, h)

	return (1 << (h - 1)) + int(nodeID)
}

func exist(root *TreeNode, id uint, h uint) bool {

	h = h - 1
	for root != nil {
		h = h - 1
		// 从左到右 判断 bitFlag 位上的 如果是1 表示右节点
		var bitFlag uint = 1 << h
		// fmt.Println(bitFlag)
		if (id & bitFlag) == bitFlag {
			root = root.Right
		} else {
			root = root.Left
		}

		if h == 0 {

			if root == nil {
				return false
			}

			return true
		}
	}

	return false
}

func getNodeID(root *TreeNode, h uint) uint {
	var end uint = (1 << (h - 1)) - 1
	var start uint = 0
	for start >= 0 && start < end {

		mid := (start + end) / 2
		// fmt.Println("mid", mid)
		if !exist(root, mid, h) {
			end = mid - 1
			continue
		}

		//mid找得到
		if start == end-1 {
			if exist(root, end, h) {
				return end
			}

			return start
		}

		start = mid

	}

	return start
}

func getDepth(root *TreeNode) uint {

	var d uint = 0
	for root != nil {
		d++
		root = root.Left
	}

	return d
}

// https://leetcode-cn.com/problems/count-complete-tree-nodes/solution/c-san-chong-fang-fa-jie-jue-wan-quan-er-cha-shu-de/
// 如果根节点的左子树深度等于右子树深度，则说明左子树为满二叉树
// 如果根节点的左子树深度大于右子树深度，则说明右子树为满二叉树
func countNodesRecu(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftDepth := getDepth(root.Left)
	rightDepth := getDepth(root.Right)
	if leftDepth == rightDepth {
		return (1 << leftDepth) + countNodesRecu(root.Right)
	}
	return (1 << rightDepth) + countNodesRecu(root.Left)

}

//上面继续优化，每次递归，左边节点的高度肯定是父节点的高度减一，所以可以优化
func countNodesRecuOpt(root *TreeNode, leftDepth uint) int {

	if root == nil {
		return 0
	}

	rightDepth := getDepth(root.Right)
	if leftDepth == rightDepth {
		return (1 << leftDepth) + countNodesRecuOpt(root.Right, leftDepth-1)
	}
	return (1 << rightDepth) + countNodesRecuOpt(root.Left, leftDepth-1)

}

func countNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	h := getDepth(root.Left)

	return countNodesRecuOpt(root, h)
}
