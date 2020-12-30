// 108. 将有序数组转换为二叉搜索树 https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
package main

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	// sortedArrayToBST(nums)
	nums = []int{-10, -3}
	sortedArrayToBST(nums)
	// nums = []int{-10}
	// sortedArrayToBST(nums)
	// nums = []int{}
	// sortedArrayToBST(nums)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST1(nums []int) *TreeNode {

	nLen := len(nums)
	if nLen == 0 {
		return nil
	}

	return sortToBST1(nums, 0, nLen-1)
}

func sortToBST1(nums []int, beg, end int) *TreeNode {

	mid := (beg + end) / 2
	if (end-beg)%2 == 1 {
		mid = mid + 1
	}
	// fmt.Println(beg, mid, end)

	root := &TreeNode{
		Val: nums[mid],
	}

	if beg == end {
		return root
	}

	if mid-1 >= beg {
		root.Left = sortToBST1(nums, beg, mid-1)

	}

	if mid+1 <= end {
		root.Right = sortToBST1(nums, mid+1, end)

	}

	return root

}

func sortedArrayToBST(nums []int) *TreeNode {

	// if len(nums)==0{
	// 	return
	// }

	return sortToBST(nums)
}
func sortToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}
	// fmt.Println(beg, mid, end)

	mid := len(nums) / 2
	root := &TreeNode{
		Val: nums[len(nums)/2],
	}

	if len(nums) == 1 {
		return root
	}

	root.Left = sortToBST(nums[0:mid])
	root.Right = sortToBST(nums[mid+1 : len(nums)])

	// if mid > 0 {
	// 	root.Left = sortToBST(nums[0:mid])

	// }

	// if mid+1 <= end {
	// 	root.Right = sortToBST(nums, mid+1, end)

	// }

	return root

}
