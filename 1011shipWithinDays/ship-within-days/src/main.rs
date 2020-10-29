fn main() {
    println!("Hello, world!");
}
pub struct Solution;
impl Solution {
    /*
    二分法，转为是否存在一个threshold，连续子数组的和 不大于 threadhold 的情况下，分为 m 个
    12ms 2.5MB  //跟410题型一样 
    */
    pub fn ship_within_days(nums: Vec<i32>, m: i32) -> i32 {
        if nums.len() < m as usize {
            return 0;
        }

        let mut left = 0; //left 必须要是nums中的最大值 
        let mut right = 0; 
        for v in &nums {
            right += v;

            if *v > left{
                left = *v
            }
        }

        // let mut left = 0;
        let mut mid;
        while left < right {
            mid = left + (right - left) / 2;
            if Self::exist(&nums, m, mid) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }

        left
    }

    //判断按照 threshold 是否可以分为 最多m 份 
    pub fn exist(nums: &Vec<i32>, m: i32, threshold: i32) -> bool {
        let mut cur_sum = 0;
        let mut cnt = 0;

        for v in nums {
            if cur_sum + v <= threshold {
                cur_sum = cur_sum + v;
            } else {
                cnt += 1;
                cur_sum = *v;

                // cnt 已经等于 m 了， 然后 还有剩余的 cur_sum ，说明最后的 分组数目 肯定大于 m
                if cnt == m {
                    return false;
                }
            }
        }


        //因为 threadhold 大于等于nums中的任意一个数字, 从循环出来后， cnt 肯定小于m，不然就提前返回false了， cur_sum 也肯定 <= threadhold
        true
    }
}
