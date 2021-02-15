/*
 * @Description: 442
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-15 19:01:04
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-15 19:13:46
 * @FilePath: /442findDuplicates/find_duplicates/src/lib.rs
 */
pub struct Solution;

impl Solution {
    pub fn find_duplicates(nums: Vec<i32>) -> Vec<i32> {
        let mut nums = nums;
        let mut res = Vec::new();
        let mut i = 0;
        let mut right_pos = 0;
        while i< nums.len() {
            right_pos = nums[i] as usize -1;
            if nums[i]==nums[right_pos]{
                i+=1;
            }else{
                nums.swap(i, right_pos);
            }
        }

        for i in 0..nums.len(){
            if i + 1 != nums[i] as usize{
                res.push(nums[i]);
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
        assert_eq!(Solution::find_duplicates(vec![4,3,2,7,8,2,3,1]),vec![3,2]);
    }
}
