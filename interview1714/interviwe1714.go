package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{62577, -220, -8737, -22, -6, 59956, 5363, -16699, 0, -10603, 64, -24528, -4818, 96, 5747, 2638, -223, 37663, -390, 35778, -4977, -3834, -56074, 7, -76, 601, -1712, -48874, 31, 3, -9417, -33152, 775, 9396, 60947, -1919, 683, -37092, -524, -8, 1458, 80, -8, 1, 7, -355, 9, 397, -30, -21019, -565, 8762, -4, 531, -211, -23702, 3, 3399, -67, 64542, 39546, 52500, -6263, 4, -16, -1, 861, 5134, 8, 63701, 40202, 43349, -4283, -3, -22721, -6, 42754, -726, 118, 51, 539, 790, -9972, 41752, 0, 31, -23957, -714, -446, 4, -61087, 84, -140, 6, 53, -48496, 9, -15357, 402, 5541, 4, 53936, 6, 3, 37591, 7, 30, -7197, -26607, 202, 140, -4, -7410, 2031, -715, 4, -60981, 365, -23620, -41, 4, -2482, -59, 5, -911, 52, 50068, 38, 61, 664, 0, -868, 8681, -8, 8, 29, 412}
	a := smallestK(arr, 131)
	sort.Ints(a)
	b := []int{-61087, -60981, -56074, -48874, -48496, -37092, -33152, -26607, -24528, -23957, -23702, -23620, -22721, -21019, -16699, -15357, -10603, -9972, -9417, -8737, -7410, -7197, -6263, -4977, -4818, -4283, -3834, -2482, -1919, -1712, -911, -868, -726, -715, -714, -565, -524, -446, -390, -355, -223, -220, -211, -140, -76, -67, -59, -41, -30, -22, -16, -8, -8, -8, -6, -6, -4, -4, -3, -1, 0, 0, 0, 1, 3, 3, 3, 4, 4, 4, 4, 4, 5, 6, 6, 7, 7, 7, 8, 8, 9, 9, 29, 30, 31, 31, 38, 51, 52, 53, 61, 64, 80, 84, 96, 118, 140, 202, 365, 397, 402, 412, 531, 539, 601, 664, 683, 775, 790, 861, 1458, 2031, 2638, 3399, 5134, 5363, 5541, 5747, 8681, 8762, 9396, 35778, 37591, 37663, 39546, 40202, 41752, 42754, 43349, 50068, 52500}
	sort.Ints(b)

	for i := range a {
		if a[i] != b[i] {
			fmt.Println("not equal")
			return
		}
	}
}

/*
面试题 17.14. 最小K个数 https://leetcode-cn.com/problems/smallest-k-lcci/
*/

func smallestK(arr []int, k int) []int {
	count := len(arr)
	if count <= k {
		return arr
	}

	if k == 0 {
		return nil
	}

	return quick(0, count, k, arr)

}

//不包含 end
func quick(beg, end int, k int, arr []int) []int {

	mid := partion(beg, end, arr)
	if mid == k-1 {
		return arr[:k]
	}

	//左边的少了
	if mid < k-1 {
		return quick(mid+1, end, k, arr)
	}

	//左边的多了
	return quick(beg, mid, k, arr)

}

// 不包含  end
func partion(beg, end int, arr []int) int {

	i := beg
	j := end - 1
	v := arr[beg] //以v进行分割

	for i < j {

		//必须先判断右边,这样返回的beg才准确
		//从右边到左查找 小于 v 的
		for i < j && arr[j] >= v {
			j--
			// 1.1  这里  i左边的（包含i）都是小于或等于 v；
			//     如果这个地方因为 i==j 退出了，j是从右往左过来的，所以 j右边的（不包含j）都是大于或等于v的。 那么1.3的规则依然满足
			//     如果这个地方因为 arr[j]<v 退出了， 那么1.3的规则依然满足
		}

		//从左边到右边查找  大于 v 的
		for i < j && arr[i] <= v {
			// 1.2 这里 j右边的（不包含j）都是大于或等于v的，但是j是小于等于v的
			//     i向右查询 如果这个地方因为 i==j 退出了，那么这里j==i，所以 i左边的（包含i）都是小于或等于 v；那么1.3的规则依然满足
			//     i向右查询 如果这个地方因为 arr[i]> v 退出了， 后面进行了交换， 那么1.3的规则依然满足
			i++
		}

		//走到这一步，只有在i < j的条件下，才表示 arr[i]>v && arr[j]<v , 进行交换
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	// 1.3 走到这里 ，i左边的（包含i）都是小于或等于 v； j右边的（不包含j）都是大于或等于v的。 这里的理解很重要
	// 见上面1.1 和  1.2的注释

	//然后将用于分割的数放到i这里，i上的树放到beg
	arr[beg] = arr[i]
	arr[i] = v

	return i

}

/*
 求最小的k数，使用最大堆  结果超时
*/
func smallestK_timeout(arr []int, k int) []int {
	count := len(arr)
	if count <= k {
		return arr
	}

	if k == 0 {
		return nil
	}

	kHeap := make([]int, k)
	copy(kHeap, arr[:k])

	buildKHeap(kHeap, k)

	for j := k; j < count; j++ {

		//如果 当前的值 arr【j】 比最大堆的最大值还要大，说明这个值就肯定不在 最小的k里面 ，跳过
		if arr[j] >= kHeap[0] {
			continue
		}

		// 当前的值 arr【j】 比最大堆的最大值要小，就替换堆顶的元素，下沉调整
		kHeap[0] = arr[j]
		buildKHeap(kHeap, k)

	}

	return kHeap

}

// 下沉调整 满足最大堆的条件
func downHeap(kHeap []int, i int) {

	k := len(kHeap)
	if k <= 1 {
		return
	}

	for i < k/2 {
		childIndex := 2*i + 1
		if childIndex+1 < k && kHeap[childIndex] < kHeap[childIndex+1] { //先求出子节点中的最大值
			childIndex++
		}

		//子树已经满足了 ，就可以退出了
		if kHeap[i] >= kHeap[childIndex] {
			return
		}

		// 交换
		kHeap[i], kHeap[childIndex] = kHeap[childIndex], kHeap[i]

		i = childIndex
	}

	return

}

// buildKHeap 构建 kHeap
func buildKHeap(kHeap []int, k int) {

	i := k/2 - 1
	for i >= 0 {

		downHeap(kHeap, i)
		i--
	}

}
