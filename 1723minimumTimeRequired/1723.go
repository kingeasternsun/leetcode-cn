package leetcode

//二分法加DFS
func minimumTimeRequired(jobs []int, k int) int {

	beg := 0
	end := 0
	for _, job := range jobs {
		if job > beg {
			beg = job
		}

		end += job
	}

	if len(jobs) == k {
		return beg
	}

	// end++
	mid := 0
	for beg < end {

		mid = (beg + end) / 2
		worker := make([]int, k)
		//满足条件
		if dfsBi(0, mid, worker, jobs, k) {
			end = mid
		} else {
			beg = mid + 1
		}

	}

	return beg
}

// 这个dfsBi算法，每次选第curID个工作，然后给某个工人，看是否满足threshould条件. worker 标识每个工人的总工作时长
func dfsBi(curID int, threshold int, worker []int, jobs []int, k int) bool {

	if curID == len(jobs) {
		return true
	}

	for i := 0; i < k; i++ {
		//如果这个工作给工人i是满足条件的,就给他
		if worker[i]+jobs[curID] <= threshold {
			worker[i] += jobs[curID]
			if dfsBi(curID+1, threshold, worker, jobs, k) {
				return true
			}
			//恢复
			worker[i] -= jobs[curID]
		}

		//关键点在这里，提前返回判断，不然就超时了. 这个没有太懂
		if worker[i] == 0 {
			return false
		}

	}

	return false

}

func minimumTimeRequiredDFS(jobs []int, k int) int {

	res := 0
	sum := 0
	for _, job := range jobs {
		if job > res {
			res = job
		}

		sum += job
	}

	if len(jobs) == k {
		return res
	}

	res = sum

	dfs(jobs, k, -1, &res, sum, -1, -1, 0)

	return res
}

func getMax(jobs []int) int {
	res := 0

	for _, job := range jobs {
		if job > res {
			res = job
		}

	}

	return res

}

// jobs当前剩余任务,k 工人数量 ,curMax :当前最大完成时间,res 目前为止最小的最大完成时间,leftSum: 剩余总数;
// preMin:上一个工人最小的工作，preMax上一个工人最大的工作,preSum上一个工人工作时间总和
func dfs(jobs []int, k int, curMax int, res *int, leftSum int, preMin, preMax, preSum int) {

	//如果剩余工作的数目 不够分到k个工人 直接返回
	if len(jobs) < k {
		return
	}

	if len(jobs) == 0 && k == 0 {
		if curMax < *res {
			*res = curMax
		}
		return
	}

	//如果工作数量等于工人数目，那么只能给剩下的工人每人分一个
	if len(jobs) == k {

		curMax = max(curMax, getMax(jobs))
		if curMax < *res {
			*res = curMax
		}

		return
	}

	// 注意 k为1的时候，这里不能直接把剩下工作的都给当前这个工人，因为jobs里面还可以继续分到上一个工人
	// if k == 1 {

	// 	if leftSum > curMax {
	// 		curMax = leftSum
	// 	}

	// 	if curMax < *res {
	// 		*res = curMax
	// 	}

	// 	return
	// }

	//当前的最大值已经比当前的一种方案大了 就可以直接返回了，不用考虑了
	if curMax > *res {
		return
	}

	addMap := make(map[int]struct{}, 0)
	for i, jobH := range jobs {

		if _, ok := addMap[jobH]; ok {
			continue
		}

		//第一种情况，当前jobH仍然给上一个工人使用
		//加限制条件，给一个工人分配的工作，必须是递增的，减少重复
		if jobH >= preMax && preSum != 0 {
			jobs[0], jobs[i] = jobs[i], jobs[0]
			// fmt.Printf(" %v", jobH)
			dfs(jobs[1:], k, max(curMax, preSum+jobH), res, leftSum-jobH, preMin, jobH, preSum+jobH)

			jobs[0], jobs[i] = jobs[i], jobs[0]

			addMap[jobH] = struct{}{}
		}

	}
	addMap = make(map[int]struct{}, 0)
	for i, jobH := range jobs {
		if _, ok := addMap[jobH]; ok {
			continue
		}

		//第二种情况，当前jobH分配给新的工人使用
		//加限制条件，后面工人拿到的第一份工作的时长也是递增的,减少重复
		if jobH >= preMin && k > 0 {
			jobs[0], jobs[i] = jobs[i], jobs[0]

			dfs(jobs[1:], k-1, max(curMax, jobH), res, leftSum-jobH, jobH, jobH, jobH)

			jobs[0], jobs[i] = jobs[i], jobs[0]
			addMap[jobH] = struct{}{}
		}

	}

}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
