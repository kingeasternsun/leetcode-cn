/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-23 15:34:05
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-23 15:50:07
 * @FilePath: \704search\search\src\lib.rs
 */
pub struct Solution;

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let mut left: usize = 0;
        let mut right: usize = nums.len() - 1;

        while left <= right {
            let mid = (right - left) / 2 + left;
            if nums[mid] == target {
                return mid as i32;
            }

            if nums[mid] > target {
                if mid == 0{
                    return -1
                }
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }

        return -1;
    }

    pub fn search1(nums: Vec<i32>, target: i32) -> i32 {
        let res = nums.binary_search(&target);
        match res{
            Ok(id)=> id as i32,
            Err(_)=> -1,
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::search(vec![-1,0,3,5,9,12], 9),4);
        assert_eq!(Solution::search(vec![-1,0,3,5,9,12], 2),-1);
        assert_eq!(Solution::search(vec![-1,0,3,5,9,12], -2),-1);
        assert_eq!(Solution::search(vec![-1,0,3,5,9,12], 30),-1);

        assert_eq!(Solution::search1(vec![-1,0,3,5,9,12], 9),4);
        assert_eq!(Solution::search1(vec![-1,0,3,5,9,12], 2),-1);
        assert_eq!(Solution::search1(vec![-1,0,3,5,9,12], -2),-1);
        assert_eq!(Solution::search1(vec![-1,0,3,5,9,12], 30),-1);
    }
}
