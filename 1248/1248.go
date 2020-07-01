package main

//https://leetcode-cn.com/problems/count-number-of-nice-subarrays/

/*
 preID,begID,endID,nextID  ,curNum:= (begID-preID)*(nextID-endID)

 [begID,endID] has k

 preID is the pre begID
 nextID is the next endID
172ms 7.3MB
*/
func numberOfSubarrays1(nums []int, k int) int {

	if k == 0 {
		return 0
	}

	num := 0
	preID := -1
	begID := -1
	endID := -1
	nextID := -1

	// get begID
	for i := range nums {
		if nums[i]&1 == 1 {
			begID = i
			break
		}
	}

	if begID == -1 {
		return num
	}

	//get endID
	kNum := 1
	if k > 1 {
		i := begID + 1
		for i < len(nums) {
			if nums[i]&1 == 1 {
				kNum++

				if kNum == k {
					endID = i
					break
				}
			}

			i++
		}

		if endID == -1 {
			return 0
		}

	} else {
		endID = begID
	}

	// compute and move the preID,beID,endID

	for endID < len(nums) {

		// get nextID
		i := endID + 1
		for i < len(nums) {
			if nums[i]&1 == 1 {
				break
			}
			i++
		}

		nextID = i

		num = num + (begID-preID)*(nextID-endID)
		// if k == 1 {
		// 	num = num - 1
		// }

		endID = nextID
		preID = begID
		i = begID + 1
		for i < len(nums) {
			if nums[i]&1 == 1 {
				break
			}
			i++
		}

		begID = i

	}

	return num
}

//132ms 7.6MB
func numberOfSubarrays2(nums []int, k int) int {

	oddIndexs := make([]int, len(nums)+2)
	oddIndexs[0] = -1
	oddCnt := 0

	for i := range nums {
		if nums[i]&1 == 1 {
			oddCnt++
			oddIndexs[oddCnt] = i
		}
	}

	if oddCnt < k {
		return 0
	}

	oddIndexs[oddCnt+1] = len(nums)

	cnt := 1
	num := 0
	for cnt+k <= oddCnt+1 {
		num = num + (oddIndexs[cnt]-oddIndexs[cnt-1])*(oddIndexs[cnt+k]-oddIndexs[cnt+k-1])
		cnt++
	}

	return num
}
