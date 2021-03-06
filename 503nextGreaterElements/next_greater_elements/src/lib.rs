/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-06 14:06:20
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-06 14:15:15
 * @FilePath: /503nextGreaterElements/next_greater_elements/src/lib.rs
 */

pub struct Solution;
impl Solution {
    pub fn next_greater_elements(nums: Vec<i32>) -> Vec<i32> {
       
        let mut res = vec![-1;nums.len()];
        let mut decre_stack = Vec::new();

        for i in 0..nums.len()*2{
            while let Some(v) = decre_stack.last(){
                if nums[*v] >= nums[i%nums.len()]{
                    break;
                }

                res[*v] = nums[i%nums.len()];
                decre_stack.pop();
            }

            decre_stack.push(i%nums.len());
        }
        res
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::next_greater_elements(vec![1,2,1]),vec![2,-1,2]);
    }
}
