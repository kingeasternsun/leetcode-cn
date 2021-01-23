package leetcode

//h不一定是citations里面的值,找到满足 n-i<=citations[i] 的最小i
func hIndex(citations []int) int {

	if len(citations) == 0 {
		return 0
	}

	beg := 0
	end := len(citations)
	h := 0
	for beg < end {
		mid := (end-beg)/2 + beg
		//当前有 len(citations)-mid 个论文的引用次数 大于等于 citations[mid]
		h = len(citations) - mid

		if h <= citations[mid] { //说明有 h个论文的引用次数大于等于 h

			end = mid

		} else { //h个论文中 有个引用次数小于 h 不满足
			beg = mid + 1
		}
	}

	return len(citations) - beg
}
