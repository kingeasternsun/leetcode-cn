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

    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut n = 0;
        let mut cnt = 0;
        nums.into_iter().for_each(|x|{
            if cnt ==0{
                n = x;
            }else if x ==n{
                cnt +=1;
            }else{
                cnt -=1;
            }
        });
        if cnt ==0{
            return  -1;
        }
        n
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
