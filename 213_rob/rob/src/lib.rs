struct Solution;

impl Solution {
    // 第一个和最后一个不能同时偷，
    //   要么不偷第一个 &nums[1..]
    //   要么不偷最后一个 &nums[0..nums.len()-1]
    pub fn rob(nums: Vec<i32>) -> i32 {
        if nums.len() <= 2 {
            return nums.into_iter().max().unwrap_or_default();
        }

        Self::dp(&nums[1..]).max(Self::dp(&nums[0..nums.len() - 1]))
    }

    // dp[i]: 表示从第i户开始偷
    // dp[i] = max(dp[i+1], nums[i] + dp[i+2])
    / LC: 0ms 2mb
    fn dp(nums: &[i32]) -> i32 {
        if nums.len() == 0 {
            return 0;
        }
        let mut a1 = nums[nums.len() - 1]; // dp[i+1]
        let mut a2 = 0; // dp[i+2]

        //       cur -> a1 -> a2
        // cur -> a1 -> a2

        for i in (0..nums.len() - 1).rev() {
            let cur = a1.max(nums[i] + a2);
            a2 = a1;
            a1 = cur;
        }
        a1.max(a2)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::rob(vec![2, 3, 2]), 3);
        assert_eq!(Solution::rob(vec![1, 2, 3, 1]), 4);
        assert_eq!(Solution::rob(vec![1, 2, 3]), 3);
        assert_eq!(Solution::rob(vec![1]), 1);
        assert_eq!(Solution::rob(vec![1, 2]), 2);
    }
}
