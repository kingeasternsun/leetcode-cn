/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-24 11:00:14
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-24 11:13:08
 * @FilePath: \1011shipWithinDays\ship_within_days\src\lib.rs
 */
pub struct Solution;

impl Solution {
    /*
    二分法，转为是否存在一个threshold，连续子数组的和 不大于 threadhold 的情况下，分为 m 个
    12ms 2.5MB  //跟410题型一样
    */
    pub fn ship_within_days(nums: Vec<i32>, m: i32) -> i32 {
        if nums.len()==0{
            return 0
        }
        let mut beg = nums.iter().max().unwrap().clone() ;
        let mut end: i32 = nums.iter().sum();

        let mut best = 0;
        while beg<=end{
            let mid = (end - beg)/2+beg;
            if Self::exist(&nums, m, mid){
                best = mid;
                end = mid-1;
            }else{
                beg = mid+1;
            }
        }

        return best

    }

    //判断按照 threshold 是否可以分为 最多m 份
    pub fn exist(nums: &Vec<i32>, m: i32, threshold: i32) -> bool {
        let mut cur_day = 1; //当前第几天
        let mut cur_weight = 0;

        for v in nums {
            if cur_weight + v > threshold {
                cur_day += 1;
                cur_weight = *v;
                if cur_day > m {
                    return false;
                }
            }else{
                cur_weight += v;
            }
        }

        true
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;
    #[test]
    fn it_works() {
        assert_eq!(
            Solution::ship_within_days(vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 5),
            15
        );
        assert_eq!(Solution::ship_within_days(vec![3, 2, 2, 4, 1, 4], 3), 6);
        assert_eq!(Solution::ship_within_days(vec![1, 2, 3, 1, 1], 4), 3);
    }
}
