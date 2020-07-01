package main

func productExceptSelf(nums []int) []int {

	count:=len(nums)
	if count <2{
		return nil
	}

	left:=make([]int,count)
	right:=make([]int,count)
	result:=make([]int,count)

	for i:=range nums{

		if i == 0{
			left[i]= 1
		}else{
			left[i] = left[i-1]*nums[i-1]
		}

	}

	for i:=count-1;i>=0;i--{

		if i == count-1{
			right[i]= 1
		}else{
			right[i] = right[i+1]*nums[i+1]
		}

	}

	for i:=0;i<count;i++{
		result[i] = left[i]*right[i]
	}

	return result

}