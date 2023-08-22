struct Solution;
impl Solution {
    // 0ms 1.89mb
    pub fn massage(nums: Vec<i32>) -> i32 {
        // nums[i..]的最最大值
        let mut dp = vec![0; nums.len() + 2];

        for i in (0..nums.len()).rev() {
            dp[i] = dp[i + 1].max(nums[i] + dp[i + 2]);
        }
        dp[0]
    }

    // 0ms 1.9mb
    pub fn massage1(nums: Vec<i32>) -> i32 {
        let mut dp = (0, 0);

        for i in (0..nums.len()).rev() {
            let new_dp = dp.0.max(nums[i] + dp.1);
            dp.1 = dp.0;
            dp.0 = new_dp;
        }
        dp.0.max(dp.1)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::massage(vec![1, 2, 3, 1]), 4);
        assert_eq!(Solution::massage(vec![2, 7, 9, 3, 1]), 12);
        assert_eq!(Solution::massage(vec![2, 1, 4, 5, 3, 1, 1, 3]), 12);
    }

    #[test]
    fn it_works1() {
        assert_eq!(Solution::massage1(vec![1, 2, 3, 1]), 4);
        assert_eq!(Solution::massage1(vec![2, 7, 9, 3, 1]), 12);
        assert_eq!(Solution::massage1(vec![2, 1, 4, 5, 3, 1, 1, 3]), 12);
    }
}
