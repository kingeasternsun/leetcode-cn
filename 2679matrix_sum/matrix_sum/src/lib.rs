struct Solution;
impl Solution {
    pub fn matrix_sum(mut nums: Vec<Vec<i32>>) -> i32 {
        nums.iter_mut().for_each(|x| x.sort_unstable());
        let mut ans = 0;
        for c in 0..nums[0].len() {
            ans += nums.iter().map(|x| x[c]).max().unwrap();
        }
        ans
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let mut nums = vec![vec![7, 2, 1], vec![6, 4, 2], vec![6, 5, 3], vec![3, 2, 1]];
        assert_eq!(Solution::matrix_sum(nums), 15);
        assert_eq!(Solution::matrix_sum(vec![vec![1]]), 1);
    }
}
