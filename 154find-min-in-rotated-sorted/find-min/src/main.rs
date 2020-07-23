fn main() {
    println!("Hello, world!");
}

pub struct Solution{}
impl Solution {
    pub fn find_min(nums: Vec<i32>) -> i32 {

        let mut beg = 0;
        if nums.len()==0{
            return 0
        }

        // if nums.len()==1{
        //     return nums[0]
        // }

        let mut end = nums.len()-1;
        
        
        while beg < end{

            let mid = (beg+end)/2;
            if nums[mid] > nums[end]{ // which means mid is in the left part that bigger, so min should in right
                beg = mid+1
            }else if nums[mid]< nums[end]{ // which means mid is in the right part that smaller. so min should in left
                end = mid  // mid may be the smallest
            }else{
                end = end-1
            }

        }

        nums[beg]
    }
}



#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_154() {
        assert_eq!(Solution::find_min(vec![1, 2, 2, 2, 2, 2]), 1);
        assert_eq!(Solution::find_min(vec![1, 3, 3]), 1);
        assert_eq!(Solution::find_min(vec![3, 1, 3, 3]), 1);
        assert_eq!(Solution::find_min(vec![1]), 1);
    }
}