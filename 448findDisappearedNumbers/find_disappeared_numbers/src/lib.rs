use core::num;
use std::mem::swap;

/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-13 11:26:49
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-15 18:56:13
 * @FilePath: /find_disappeared_numbers/src/lib.rs
 */
pub struct Solution;

impl Solution {
    pub fn find_disappeared_numbers1(nums: Vec<i32>) -> Vec<i32> {
        let mut tmp = vec![false; nums.len()];
        let mut res = Vec::new();

        for v in nums {
            tmp[v as usize - 1] = true
        }
    

        for i in 0..tmp.len() {
            if !tmp[i] {
                res.push(i as i32 + 1);
            }
        }

        res
    }

    pub fn find_disappeared_numbers(nums: Vec<i32>) -> Vec<i32> {
        let mut res = Vec::new();
        let mut nums = nums;

        let mut i = 0;
        while i < nums.len(){
            let right_pos = nums[i] as usize -1;
            if nums[i]==nums[right_pos]{
                i+=1;
            }else{
                nums.swap(i, right_pos);
            }
        }

        for i in 0..nums.len(){
            if i != nums[i] as usize -1{
                res.push(i as i32 +1);
            }
        }

        res
    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::find_disappeared_numbers1(vec![4, 3, 2, 7, 8, 2, 3, 1]),vec![5, 6]);
        assert_eq!(Solution::find_disappeared_numbers(vec![4, 3, 2, 7, 8, 2, 3, 1]),vec![5, 6]);
    }
}
