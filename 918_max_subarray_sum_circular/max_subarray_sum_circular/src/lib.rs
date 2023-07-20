struct Solution;
impl Solution {
    pub fn max_subarray_sum_circular(nums: Vec<i32>) -> i32 {
        if nums.len() == 0 {
            return 0;
        }

        let mut res = nums[0];
        let mut pre_max_sum = nums[0]; // the max subarray sum that end with id
        let mut tmp_left_sum = nums[0]; // the sum from 0 to id
        let mut max_left_sum = vec![0; nums.len()]; // the max left sum that begin from 0
        max_left_sum[0] = nums[0];

        for i in 1..nums.len() {
            pre_max_sum = nums[i].max(nums[i] + pre_max_sum);
            res = res.max(pre_max_sum);

            tmp_left_sum += nums[i];
            max_left_sum[i] = tmp_left_sum.max(max_left_sum[i-1]);
        }

        let mut right_sum = 0;
        for i in (1..nums.len()).rev() {
            right_sum += nums[i];
            res = res.max(max_left_sum[i - 1] + right_sum);
        }

        res
    }

    pub fn max_subarray_sum_circular2(nums: Vec<i32>) -> i32 {
        if nums.len() == 0 {
            return 0;
        }

        if nums.iter().all(|&x| x >= 0) {
            return nums.iter().sum();
        }

        // we need recored the starting index and the max sum that end with current index
        let mut dp = Vec::with_capacity(2 * nums.len());
        for id in 0..2 * nums.len() {
            let r_id = id % nums.len();
            if id == 0 || dp[id - 1] <= 0 {
                dp.push(nums[r_id]);
            } else {
                dp.push(dp[id - 1] + nums[r_id]);
            }
            // println!("{}", dp);
        }
        dp.into_iter().max().unwrap()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        // assert_eq!(Solution::max_subarray_sum_circular2(vec![-1]), -1);
        // assert_eq!(Solution::max_subarray_sum_circular2(vec![-1, -2]), -1);

        // assert_eq!(Solution::max_subarray_sum_circular2(vec![1]), 1);
        // assert_eq!(Solution::max_subarray_sum_circular2(vec![1, 2]), 3);
        // assert_eq!(Solution::max_subarray_sum_circular2(vec![1, 2, 3]), 6);

        // assert_eq!(Solution::max_subarray_sum_circular2(vec![5, -3, 5]), 10);
        // assert_eq!(Solution::max_subarray_sum_circular2(vec![3, -2, 2, -3]), 3);
    }

    #[test]
    fn it_works2() {
        assert_eq!(Solution::max_subarray_sum_circular(vec![-1]), -1);
        assert_eq!(Solution::max_subarray_sum_circular(vec![-1, -2]), -1);

        assert_eq!(Solution::max_subarray_sum_circular(vec![1]), 1);
        assert_eq!(Solution::max_subarray_sum_circular(vec![1, 2]), 3);
        assert_eq!(Solution::max_subarray_sum_circular(vec![1, 2, 3]), 6);

        assert_eq!(Solution::max_subarray_sum_circular(vec![5, -3, 5]), 10);
        assert_eq!(Solution::max_subarray_sum_circular(vec![3, -2, 2, -3]), 3);
    }
}
