/*
 * @Description:442
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-13 23:25:20
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-15 19:08:25
 * @FilePath: /442findDuplicates/442.go
 */
package leetcode

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
func findDuplicates(nums []int) []int {

	res := make([]int, 0)
	i := 0
	for i < len(nums) {
		rightPos := nums[i] - 1
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

	for i, v := range nums {
		if i != v-1 {
			res = append(res, v)
		}
	}

	return res

}
