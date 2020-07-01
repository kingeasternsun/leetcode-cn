package main
import(
	"sort"
	"math"
)

func abs(a int)int{
	if a>=0{
		return a
	}

	return -a
}

func threeSumClosest(nums []int, target int) int {

	count := len(nums)
	if count < 3 {
		return 0
	}

	sort.Ints(nums)

	min:=math.MaxInt32
	minSum:=0

	for i := 0; i < count-2; i++ {

		tmpt := target - nums[i]

		beg := i + 1
		end := count - 1

		for beg<end{

			for beg < end && (nums[beg]+nums[end] < tmpt) {
				if abs(nums[beg]+nums[end]-tmpt)<min{
					min = abs(nums[beg]+nums[end]-tmpt)
					minSum = nums[beg]+nums[end]+nums[i]
				}
				beg++
			}
	
			if beg < end && nums[beg]+nums[end] == tmpt {
				return target
			}
	
			for beg <end && (nums[beg]+nums[end] > tmpt){
				if abs(nums[beg]+nums[end]-tmpt)<min{
					min = abs(nums[beg]+nums[end]-tmpt)
					minSum = nums[beg]+nums[end]+nums[i]
				}
				end--
			}
	
			if beg < end && nums[beg]+nums[end] == tmpt {
				return target
			}

		}


	}

	return minSum

}
