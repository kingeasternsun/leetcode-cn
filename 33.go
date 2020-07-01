package main

// func main() {
// 	nums := []int{1, 3}
// 	fmt.Println(findPoint(nums, 0, len(nums)-1))
// 	fmt.Println(search(nums, 1))
// }

func search(nums []int, target int) int {
	sLen := len(nums)
	if sLen == 0 {
		return -1
	}

	mid := findPoint(nums, 0, sLen-1)
	if mid < 0 {
		return binSearch(nums, 0, sLen-1, target)
	}
	if target >= nums[0] {
		return binSearch(nums, 0, mid, target)
	}

	return binSearch(nums, mid, sLen-1, target)

}

func findPoint(nums []int, beg, end int) int {

	if beg == end {
		return beg
	}
	if end == beg+1 {
		if nums[beg] > nums[end] {
			return beg
		}

		return -1
	}

	mid := (beg + end) / 2
	if nums[mid] > nums[beg] {
		return findPoint(nums, mid, end)
	}

	return findPoint(nums, beg, mid)

}

func binSearch(nums []int, beg, end int, target int) int {

	// fmt.Println(beg, end)
	if beg == end {
		if nums[beg] == target {
			return beg
		}

		return -1
	}

	if end == beg+1 {
		if nums[beg] == target {
			return beg
		}

		if nums[end] == target {
			return end
		}

		return -1
	}

	mid := (beg + end) / 2

	if nums[mid] == target {
		return mid
	}

	//mid在左边的折叠段
	if nums[mid] > target {
		return binSearch(nums, beg, mid-1, target)
	}

	return binSearch(nums, mid+1, end, target)

}

/*
class Solution {
    public int search(int[] nums, int target) {
        int left = 0;
        int right = nums.length-1;
        int mid = left + (right-left)/2;

        while(left <= right){
            if(nums[mid] == target){
                return mid;
            }

            if(nums[left] <= nums[mid]){  //左边升序
                if(target >= nums[left] && target <= nums[mid]){//在左边范围内
                    right = mid-1;
                }else{//只能从右边找
                    left = mid+1;
                }

            }else{ //右边升序
                if(target >= nums[mid] && target <= nums[right]){//在右边范围内
                    left = mid +1;
                }else{//只能从左边找
                    right = mid-1;
                }

            }
            mid = left + (right-left)/2;
        }

        return -1;  //没找到
    }
}
*/
