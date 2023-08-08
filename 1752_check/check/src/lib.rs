struct Solution;

impl Solution {
    // 找到最小值，然后从最小值开始判断是否非递减
    pub fn check(nums: Vec<i32>) -> bool {
        let min = nums.iter().enumerate().min_by_key(|x| x.1).unwrap();
        let mut pre = min.1.clone();
        for i in 0..nums.len() {
            if nums[(min.0 + i) % nums.len()] < pre {
                return false;
            }
            pre = nums[(min.0 + i) % nums.len()]
        }

        true
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::check(vec![3, 4, 5, 1, 2]), true);
        assert_eq!(Solution::check(vec![2, 1, 3, 4]), false);
        assert_eq!(Solution::check(vec![1, 2, 3]), true);
        assert_eq!(Solution::check(vec![1, 2]), true);
        assert_eq!(Solution::check(vec![2, 1]), true);
    }
}
