package main

func main() {

}

/*
 求最小的k数，使用最大堆  结果超时
 面试题40. 最小的k个数 https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/

 类似  面试题 17.14. 最小K个数 https://leetcode-cn.com/problems/smallest-k-lcci/
*/

func getLeastNumbers(arr []int, k int) []int {
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

func getLeastNumbers1(arr []int, k int) []int {
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
