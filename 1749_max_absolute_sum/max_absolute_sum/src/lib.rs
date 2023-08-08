struct Solution;

impl Solution {
    // 分别求最大值，最小值，然后取绝对值的最大值即可
    pub fn max_absolute_sum(nums: Vec<i32>) -> i32 {
        Self::max_sum(&nums).abs().max(Self::min_sum(&nums).abs())
    }

    fn max_sum(nums: &Vec<i32>) -> i32 {
        let mut max = 0;
        let mut pre = 0;
        for n in nums {
            // cur: 当前元素作为结尾的最大子数组
            let mut cur = n.clone();
            if pre > 0 {
                cur += pre;
            }
            pre = cur;
            max = max.max(cur);
        }

        max
    }

    fn min_sum(nums: &Vec<i32>) -> i32 {
        let mut min = 0;
        let mut pre = 0;
        for n in nums {
            // cur: 当前元素作为结尾的最小子数组
            let mut cur = n.clone();
            if pre < 0 {
                cur += pre;
            }
            pre = cur;
            min = min.min(cur);
        }

        min
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::max_absolute_sum(vec![1, -3, 2, 3, -4]), 5);
        assert_eq!(Solution::max_absolute_sum(vec![2, -5, 1, -4, 3, -2]), 8);
        assert_eq!(Solution::max_absolute_sum(vec![1]), 1);
        assert_eq!(Solution::max_absolute_sum(vec![0]), 0);
        assert_eq!(Solution::max_absolute_sum(vec![-1]), 1);
        assert_eq!(Solution::max_absolute_sum(vec![1, -1]), 1);
    }
}
