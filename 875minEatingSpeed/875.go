/*
 * @Description: 875
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-10 15:00:51
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-10 15:51:43
 * @FilePath: /875minEatingSpeed/875.go
 */

package leetcode

/*
 二分查找时间
*/
func minEatingSpeed(piles []int, H int) int {

	left := 0
	right := 0
	for _, p := range piles {
		if p > right {
			right = p
		}
	}
	best := 0
	for left <= right {
		// {"4", args{[]int{2, 2}, 2}, 2}, //导致最后 left：0， right：1， mid得到0
		mid := (right-left)/2 + left
		if eatup(piles, H, mid) {
			best, right = mid, mid-1
		} else {
			left = mid + 1
		}
	}

	return best
}

func eatup(piles []int, H int, speed int) bool {
	if speed == 0 {
		return false
	}

	h := 0
	for pID := range piles {
		if h >= H { //已经过去了h个小时 还没有吃完
			return false
		}
		h += (piles[pID] / speed)
		if piles[pID]%speed > 0 {
			h++
		}
	}

	return !(h > H) //h 到这里表示吃完耗费的小时，如果h==H，也算是满足条件的，这里要注意

}
