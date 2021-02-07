/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-04 10:18:09
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-04 10:34:41
 * @FilePath: \find-max-average\src\lib.rs
 */

pub struct Solution;
impl Solution {
    pub fn find_max_average(nums: Vec<i32>, k: i32) -> f64 {
        let k = k as usize;
        if k == 1 {
            let r = nums.iter().max().unwrap();
            return *r as f64;
        }

        let mut sum: i32 = nums.iter().take(k as usize).sum();
        let mut res = sum as f64 / k as f64;

        for i in k..nums.len() {
            sum += &nums[i] - &nums[i - k];
            let tmp = sum as f64 / k as f64;
            if tmp > res {
                res = tmp;
            }
        }

        return res;
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::find_max_average(vec![1, 12, -5, -6, 50, 3], 4),
            12.75
        );
        assert_eq!(
            Solution::find_max_average(vec![1, 12, -5, -6, 50, 1], 1),
            50.00
        );
    }
}
