/*
 * @Description: 448
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-13 11:21:17
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-15 18:26:35
 * @FilePath: /448findDisappearedNumbers/448.go
 */
package leetcode

func findDisappearedNumbers1(nums []int) []int {
	tmp := make([]bool, len(nums))

	for _, v := range nums {
		tmp[v-1] = true
	}

	res := make([]int, 0)

	for i, v := range tmp {
		if !v {
			res = append(res, i+1)
		}
	}
	return res
}

/*
移形换影
把数组中的value放入到index为value-1的位置上
然后找出数组中index != nums[index]-1 的index+1
这种写法就有点像换座位，A换到B的座位，B暂时先搬到一个临时位置，然后B搬到C的位置，C搬到临时位置
*/
func findDisappearedNumbers2(nums []int) []int {

	for i := range nums {
		if i == nums[i]-1 {
			continue
		}

		ch := nums[i]
		nextID := ch - 1
		//注意这里的终止条件 有可能回到i，也有可能在其他地方就提前遇到 nextID == nums[nextID]-1
		for nextID != i && nextID != nums[nextID]-1 {
			nums[nextID], ch = ch, nums[nextID]
			nextID = ch - 1
		}
		nums[i] = ch

	}
	res := make([]int, 0)
	for i, v := range nums {
		if i != v-1 {
			res = append(res, i+1)
		}
	}

	return res
}

/*
相似题目 442 448

参考了 DDY1024 https://talkgo.org/t/topic/1495/43 的写法 太漂亮了

这种写法就是 ，A要换到B的座位，A就和B互相换位置，这样A就换好了；然后B要搬到C的位置，B就和C换位置，B就换好了
A->B
B->C
C->A

原来的位置 		|A|B|C|
然后A和B互换	|B|A|C|
然后B和C互换	|C|A|B|
*/
func findDisappearedNumbers(nums []int) []int {

	i := 0
	for i < len(nums) {
		rightPos := nums[i] - 1 //表示nums[i]要放入的新位置
		/*这里特别巧妙，
		如果当前i已经是要放入的新位置，就原地不动; 也就是 rightPos == i
		如果在rightPos上的value和 当前 nums[i] 相同，那么也没有必要再进行移动了
		nums[rightPos] == nums[i] 直接就可以覆盖上面的两种情况
		*/
		if nums[rightPos] == nums[i] {
			i++
		} else {
			nums[rightPos], nums[i] = nums[i], nums[rightPos]
			//注意，这里i是要保持不变的
		}

	}

	res := make([]int, 0)
	for i, v := range nums {
		if i != v-1 {
			res = append(res, i+1)
		}
	}

	return res
}
