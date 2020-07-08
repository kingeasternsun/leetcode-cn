package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {

	h := getDepth(root)
	if h <= 1 {
		return int(h)
	}

	// fmt.Println(h)

	nodeID := getNodeID(root, h)

	return (1 << (h - 1)) + int(nodeID)
}

func getDepth(root *TreeNode) uint {

	var d uint = 0
	for root != nil {
		d++
		root = root.Left
	}

	return d
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
