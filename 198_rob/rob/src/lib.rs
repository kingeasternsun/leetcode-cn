struct Solution;

impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let mut max_rob = 0;
        Self::dp(&nums[0..], 0, &mut max_rob);
        max_rob
    }
    fn dp(nums: &[i32], cur_sum: i32, max_rob: &mut i32) {
        if nums.len() == 0 {
            *max_rob = cur_sum.max(*max_rob);
            return;
        }

        if nums.len() == 1 {
            *max_rob = (cur_sum + nums[0]).max(*max_rob);
            return;
        }

        // rob nums[0], should skip nums[1]
        Self::dp(&nums[2..], cur_sum + nums[0], max_rob);
        // not rob nums[0]
        Self::dp(&nums[1..], cur_sum, max_rob);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::rob(vec![1, 2, 3, 1]), 4);
        assert_eq!(Solution::rob(vec![2, 7, 9, 3, 1]), 12);
        assert_eq!(Solution::rob(vec![1]), 1);
        assert_eq!(Solution::rob(vec![1, 2]), 2);
    }
}
