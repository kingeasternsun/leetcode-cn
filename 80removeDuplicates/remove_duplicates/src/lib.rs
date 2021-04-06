/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-04-06 23:08:31
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-04-06 23:23:27
 * @FilePath: /remove_duplicates/src/lib.rs
 */

 /*
 i指向新数组的最后一个数字
 j指向原来数组的当前数字
 比较 nums[j] ，nums[i] 是否相同，如果相同再判断 nums[j] ,nums[i-1]是否相同
 */
pub struct Solution;
impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        if nums.len()<=2{
            return nums.len() as i32;
        }

        let mut i = 1;
        let mut j = 2;
        while j < nums.len(){
            if nums[j]!=nums[i] ||nums[j]!=nums[i-1] {
                nums[i+1]=nums[j];
                i+=1;
                j+=1;
            }else{
                j+=1;
            }
        }
        return i as i32 + 1;

    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;
    #[test]
    fn it_works() {
        assert_eq!(Solution::remove_duplicates(& mut vec![1,1,1,2,2,3]),5);
        assert_eq!(Solution::remove_duplicates(& mut vec![0,0,1,1,1,1,2,3,3]),7);
    }
}
