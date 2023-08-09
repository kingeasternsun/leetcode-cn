struct Solution;

impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let mut max_rob = 0;
        Self::dfs(&nums[0..], 0, &mut max_rob);
        max_rob
    }
    // 对于每一户，对应偷或不偷两种状态，最终是求在不同组合下的结果，典型的DFS题型
    // LC: timeout
    fn dfs(nums: &[i32], cur_sum: i32, max_rob: &mut i32) {
        if nums.len() == 0 {
            *max_rob = cur_sum.max(*max_rob);
            return;
        }

        if nums.len() == 1 {
            *max_rob = (cur_sum + nums[0]).max(*max_rob);
            return;
        }

        // rob nums[0], should skip nums[1]
        Self::dfs(&nums[2..], cur_sum + nums[0], max_rob);
        // not rob nums[0]
        Self::dfs(&nums[1..], cur_sum, max_rob);
    }


    pub fn rob1(nums: Vec<i32>) -> i32 {
        Self::dp(&nums[0..])
    }

    // 从第 i 户开始可以得到的最大值 跟 
    //  从第 i+1 户开始的最大值， 表示不抢第 i 户
    //  从第 i+2 户开始的最大值， 表示抢第 i+1 户
    // 的最大值相关
    // dp[i] = max( dp[i+1] ,nums[i] + dp[i-2] )
    // LC: timeout
    fn dp(nums: &[i32]) -> i32 {
        if nums.len() <= 1 {
            return nums.iter().sum();
        }

        // rob cur
        let ret1 = Self::dp(&nums[2..]) + nums[0];
        // not rob cur
        let ret2 = Self::dp(&nums[1..]);
        ret1.max(ret2)
    }

    // 4ms 1.92mb
    pub fn rob2(nums: Vec<i32>) -> i32 {
        if nums.len()==1{
            return nums[0]
        }

        let mut a1 =  nums[nums.len()-1];
        let mut a2 = 0;
        let mut cur = 0;
        for id in (0..nums.len()-1).rev(){
            cur = a1.max(nums[id] + a2);
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
        assert_eq!(Solution::rob(vec![1, 2, 3, 1]), 4);
        assert_eq!(Solution::rob(vec![2, 7, 9, 3, 1]), 12);
        assert_eq!(Solution::rob(vec![1]), 1);
        assert_eq!(Solution::rob(vec![1, 2]), 2);
    }

    #[test]
    fn it_works1() {
        assert_eq!(Solution::rob1(vec![1, 2, 3, 1]), 4);
        assert_eq!(Solution::rob1(vec![2, 7, 9, 3, 1]), 12);
        assert_eq!(Solution::rob1(vec![1]), 1);
        assert_eq!(Solution::rob1(vec![1, 2]), 2);
    }

    #[test]
    fn it_works2() {
        assert_eq!(Solution::rob2(vec![1, 2, 3, 1]), 4);
        assert_eq!(Solution::rob2(vec![2, 7, 9, 3, 1]), 12);
        assert_eq!(Solution::rob2(vec![1]), 1);
        assert_eq!(Solution::rob2(vec![1, 2]), 2);
    }
}
